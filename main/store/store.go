package store

type Operation struct {
	Operator string `json:"operator"`
	Val      int    `json:"val"`
}

type Calculation struct {
	Id           string      `json:"id"`
	InitialValue int         `json:"initialValue"`
	Operations   []Operation `json:"operations"`
}

type Store interface {
	Create(calculation *Calculation) error
	GetById(string) (*Calculation, error)
	Delete(calculation *Calculation) error
}
