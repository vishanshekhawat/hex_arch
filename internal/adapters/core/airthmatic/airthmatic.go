package airthmatic

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Addition(a, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Substraction(a, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multiplication(a, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Divison(a, b int32) (int32, error) {
	return a / b, nil
}
