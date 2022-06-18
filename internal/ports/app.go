package ports

type APIPort interface {
	GetAddition(a, b int32) (int32, error)
	GetSubstraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivison(a, b int32) (int32, error)
}
