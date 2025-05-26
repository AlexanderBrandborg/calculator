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

type FireDB struct {
	*db.Client
}

var fireDB FireDB

func FirebaseDB() *FireDB {
	return &fireDB
}

type FireDBStore struct {
	*FireDB
}

func (db *FireDB) Connect() (*FireDBStore, error) {
	home, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	opt := option.WithCredentialsFile(filepath.Join(home, "calculator-d17e0-firebase-adminsdk-fbsvc-f775f37e36.json"))
	config := &firebase.Config{DatabaseURL: "https://calculator-d17e0-default-rtdb.europe-west1.firebasedatabase.app/"}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app: %v", err)
	}
	client, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase database: %v", err)
	}
	db.Client = client

	return &FireDBStore{FireDB: db}, nil
}

func (s *FireDBStore) Create(c *Calculation) error {
	if err := s.NewRef("calculations/"+c.Id).Set(context.Background(), c); err != nil {
		return err
	}
	return nil
}

func (s *FireDBStore) Delete(c *Calculation) error {
	return s.NewRef("calculations/" + c.Id).Delete(context.Background())
}

func (s *FireDBStore) GetById(cId string) (*Calculation, error) {
	calculation := &Calculation{}
	if err := s.NewRef("calculations/"+cId).Get(context.Background(), calculation); err != nil {
		return nil, err
	}

	if calculation.Id == "" {
		return nil, errors.New("error: calculation with id could not be found")
	}

	return calculation, nil
}
