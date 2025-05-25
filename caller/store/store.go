package store

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

// DATA MODEL
type Operation struct {
	Operator string `json:"operator"`
	Val      int    `json:"val"`
}

type Calculation struct {
	Id           string      `json:"id"`
	InitialValue int         `json:"initialValue"`
	Operations   []Operation `json:"operations"`
}

// F I R E D B
type FireDB struct {
	*db.Client
}

var fireDB FireDB

func (db *FireDB) Connect() error {
	home, err := os.Getwd()
	if err != nil {
		return err
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(filepath.Join(home, "calculator-d17e0-firebase-adminsdk-fbsvc-f775f37e36.json"))
	config := &firebase.Config{DatabaseURL: "https://calculator-d17e0-default-rtdb.europe-west1.firebasedatabase.app/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return fmt.Errorf("error initializing firebase app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return fmt.Errorf("error initializing firebase database: %v", err)
	}
	db.Client = client
	return nil
}

func FirebaseDB() *FireDB {
	return &fireDB
}

type Store struct {
	*FireDB
}

func NewStore() *Store {
	d := FirebaseDB()
	return &Store{
		FireDB: d,
	}
}

func (s *Store) Create(c *Calculation) error {
	if err := s.NewRef("calculations/"+c.Id).Set(context.Background(), c); err != nil {
		return err
	}
	return nil
}

func (s *Store) Delete(c *Calculation) error {
	return s.NewRef("calculations/" + c.Id).Delete(context.Background())
}

func (s *Store) GetById(c string) (*Calculation, error) {
	calculation := &Calculation{}
	if err := s.NewRef("calculations/"+c).Get(context.Background(), calculation); err != nil {
		return nil, err
	}

	if calculation.Id == "" {
		return nil, errors.New("error: calculation with id could not be found")
	}

	return calculation, nil
}
