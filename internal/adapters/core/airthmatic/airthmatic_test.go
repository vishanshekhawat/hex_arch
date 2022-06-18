package airthmatic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected:%v , got:%v", nil, err)
	}
	require.Equal(t, answer, int32(2))
}

func TestSubstraction(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Substraction(1, 1)
	if err != nil {
		t.Fatalf("expected:%v , got:%v", nil, err)
	}
	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected:%v , got:%v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}

func TestDivison(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Divison(1, 1)
	if err != nil {
		t.Fatalf("expected:%v , got:%v", nil, err)
	}
	require.Equal(t, answer, int32(1))
}
