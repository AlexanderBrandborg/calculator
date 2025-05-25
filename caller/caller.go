package main

// NOTE: Deployed version has a fit if you send it a GET with a body. While local version doesn't.

import (
	"alexander/caller/calculation"
	"alexander/caller/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type InitInput struct {
	Value int `json:"value"`
}

type OperationInput struct {
	Value int `json:"value"`
}

type IdOutput struct {
	Id string `json:"id"`
}

type ResultOutput struct {
	Id     string  `json:"id"`
	Result float64 `json:"result"`
}

// ROUTING
// Router struct makes the 'store' available in the handlers without the need for a global variable.
type Router struct {
	store *model.Store
}

func (router *Router) initHandler(w http.ResponseWriter, r *http.Request) {
	// Validate body
	initInput := InitInput{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&initInput); err != nil {
		http.Error(w, fmt.Errorf("error: body could not be unmarshalled. error='%w'", err).Error(), 400)
		return
	}

	// Add new calculation to database
	newUuid := uuid.New()
	newCalculation := model.Calculation{Id: newUuid.String(), InitialValue: initInput.Value, Operations: make([]model.Operation, 0)}

	if err := router.store.Create(&newCalculation); err != nil {
		http.Error(w, "error: failed to update calculation in database", 500)
		return
	}

	jsonBytes, _ := json.Marshal(IdOutput{Id: newUuid.String()})
	w.Write(jsonBytes)
}

func (router *Router) enterHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := router.verifyId(r.PathValue("id")); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	c, fetchError := router.store.GetById(id)
	if fetchError != nil {
		http.Error(w, fmt.Errorf("error: id does not match any known calculations. id='%s'", id).Error(), 404) // Pass ID in here
		return
	}

	result, err := calculation.Enter(c)

	if err != nil {
		http.Error(w, fmt.Errorf("error: could not reduce expression. error='%w'", err).Error(), 500)
	}

	jsonBytes, _ := json.Marshal(ResultOutput{Id: id, Result: result})
	w.Write(jsonBytes)
}

func (router *Router) verifyId(id string) error {
	if err := uuid.Validate(id); err != nil {
		return fmt.Errorf("error: id is not a legal uuid. id='%s'", id)
	}
	return nil
}

func (router *Router) extendCalculation(w http.ResponseWriter, r *http.Request, extendFunc func(*model.Calculation, int)) {
	// Fetch Calculation
	id := r.PathValue("id")
	if err := router.verifyId(id); err != nil {
		http.Error(w, err.Error(), 400)
	}

	c, fetchError := router.store.GetById(id)
	if fetchError != nil {
		http.Error(w, fmt.Errorf("error: id does not match any known calculations. id='%s'", id).Error(), 404) // Pass ID in here
		return
	}

	// Validate body
	opInput := OperationInput{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&opInput); err != nil {
		http.Error(w, fmt.Errorf("error: body could not be unmarshalled. error='%w'", err).Error(), 400) // Split this out into multiple error messages?
		return
	}

	// Update and respond
	extendFunc(c, opInput.Value)
	if err := router.store.Create(c); err != nil {
		http.Error(w, "error: failed to update calculation in database", 500)
		return
	}

	jsonBytes, _ := json.Marshal(IdOutput{Id: id})
	w.Write(jsonBytes)
}

func (router *Router) addHandler(w http.ResponseWriter, r *http.Request) {
	router.extendCalculation(w, r, calculation.Add)
}

func (router *Router) subtractHandler(w http.ResponseWriter, r *http.Request) {
	router.extendCalculation(w, r, calculation.Subtract)
}

func (router *Router) multiplyHandler(w http.ResponseWriter, r *http.Request) {
	router.extendCalculation(w, r, calculation.Multiply)
}

func (router *Router) divideHandler(w http.ResponseWriter, r *http.Request) {
	router.extendCalculation(w, r, calculation.Divide)
}

func (router *Router) statusHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	c, _ := router.store.GetById(id)

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	w.Write(jsonBytes)
}

func (router *Router) deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if err := router.verifyId(id); err != nil {
		http.Error(w, err.Error(), 400)
	}

	c, fetchError := router.store.GetById(id)
	if fetchError != nil {
		http.Error(w, fmt.Errorf("error: id does not match any known calculations. id='%s'", id).Error(), 404)
		return
	}

	err := router.store.Delete(c)
	if err != nil {
		http.Error(w, fmt.Errorf("error: failed to delete calculation. id='%s'", id).Error(), 500)
		return
	}
}

func main() {
	log.Print("calculator: starting server...")

	err := model.FirebaseDB().Connect()
	if err != nil {
		log.Fatal(err)
	}

	router := &Router{store: model.NewStore()}

	http.HandleFunc("POST /v1/init", router.initHandler)
	http.HandleFunc("GET /v1/{id}", router.statusHandler)
	http.HandleFunc("PATCH /v1/add/{id}", router.addHandler)
	http.HandleFunc("PATCH /v1/sub/{id}", router.subtractHandler)
	http.HandleFunc("PATCH /v1/mult/{id}", router.multiplyHandler)
	http.HandleFunc("PATCH /v1/div/{id}", router.divideHandler)
	http.HandleFunc("GET /v1/enter/{id}", router.enterHandler)
	http.HandleFunc("DELETE /v1/{id}", router.deleteHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Note: Log all http errors
	log.Printf("calculator: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
