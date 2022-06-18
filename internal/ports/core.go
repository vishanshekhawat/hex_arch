package ports

type AirthmaticPort interface {
	Addition(a int32, b int32) (int32, error)
	Substraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Divison(a int32, b int32) (int32, error)
}
