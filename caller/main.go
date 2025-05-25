package main

// NOTE: Deployed version has a fit if you send it a GET with a body. While local version doesn't.

import (
	"alexander/main/calculation"
	"alexander/main/store"
	"context"
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

const calculationKey string = "calculation"

// ROUTING
type Router struct {
	store store.Store
}

func (router *Router) idLookupHandlerFunc(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if err := uuid.Validate(id); err != nil {
			http.Error(w, fmt.Errorf("error: id is not a legal uuid. id='%s'", id).Error(), 400)
			return
		}

		c, fetchError := router.store.GetById(id)
		if fetchError != nil {
			http.Error(w, fmt.Errorf("error: id does not match any known calculations. id='%s'", id).Error(), 404)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, calculationKey, c)
		r = r.WithContext(ctx)
		handler(w, r)
	}
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
	newCalculation := store.Calculation{Id: newUuid.String(), InitialValue: initInput.Value, Operations: make([]store.Operation, 0)}

	if err := router.store.Create(&newCalculation); err != nil {
		http.Error(w, "error: failed to update calculation in database", 500)
		return
	}

	jsonBytes, _ := json.Marshal(IdOutput{Id: newUuid.String()})
	w.Write(jsonBytes)
}

func (router *Router) enterHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(calculationKey).(*store.Calculation)

	result, err := calculation.Enter(c)

	if err != nil {
		http.Error(w, fmt.Errorf("error: could not reduce expression. error='%w'", err).Error(), 500)
		return
	}

	jsonBytes, _ := json.Marshal(ResultOutput{Id: c.Id, Result: result})
	w.Write(jsonBytes)
}

func (router *Router) extendCalculation(w http.ResponseWriter, r *http.Request, extendFunc func(*store.Calculation, int) error) {
	c := r.Context().Value(calculationKey).(*store.Calculation)

	// Validate body
	opInput := OperationInput{}
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&opInput); err != nil {
		http.Error(w, fmt.Errorf("error: body could not be unmarshalled. error='%w'", err).Error(), 400)
		return
	}

	// Update and respond
	if err := extendFunc(c, opInput.Value); err != nil {
		http.Error(w, err.Error(), 400)
	}

	if err := router.store.Create(c); err != nil {
		http.Error(w, "error: failed to update calculation", 500)
		return
	}

	jsonBytes, _ := json.Marshal(IdOutput{Id: c.Id})
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
	c := r.Context().Value(calculationKey).(*store.Calculation)

	jsonBytes, err := json.Marshal(c)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	w.Write(jsonBytes)
}

func (router *Router) deleteHandler(w http.ResponseWriter, r *http.Request) {
	c := r.Context().Value(calculationKey).(*store.Calculation)
	err := router.store.Delete(c)
	if err != nil {
		http.Error(w, fmt.Errorf("error: failed to delete calculation. id='%s'", c.Id).Error(), 500)
		return
	}
}

func main() {
	log.Print("calculator: starting server...")

	store, err := store.FirebaseDB().Connect()
	if err != nil {
		log.Fatal(err)
	}

	router := &Router{store: store}

	http.HandleFunc("POST /v1/init", router.initHandler)
	http.HandleFunc("GET /v1/{id}", router.idLookupHandlerFunc(router.statusHandler))
	http.HandleFunc("PATCH /v1/add/{id}", router.idLookupHandlerFunc(router.addHandler))
	http.HandleFunc("PATCH /v1/sub/{id}", router.idLookupHandlerFunc(router.subtractHandler))
	http.HandleFunc("PATCH /v1/mult/{id}", router.idLookupHandlerFunc(router.multiplyHandler))
	http.HandleFunc("PATCH /v1/div/{id}", router.idLookupHandlerFunc(router.divideHandler))
	http.HandleFunc("GET /v1/enter/{id}", router.idLookupHandlerFunc(router.enterHandler))
	http.HandleFunc("DELETE /v1/{id}", router.idLookupHandlerFunc(router.deleteHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("calculator: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
