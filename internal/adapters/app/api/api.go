package api

import "github.com/hex/internal/ports"

type Adapter struct {
	airth ports.AirthmaticPort
	db    ports.DBPort
}

func NewAdapter(arith ports.AirthmaticPort, db ports.DBPort) *Adapter {
	return &Adapter{airth: arith, db: db}
}

func (api Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := api.airth.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := api.airth.Substraction(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "substract")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := api.airth.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (api Adapter) GetDivison(a, b int32) (int32, error) {
	answer, err := api.airth.Divison(a, b)
	if err != nil {
		return 0, err
	}
	err = api.db.AddToHistory(answer, "divison")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
