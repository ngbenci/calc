package calc_test

import (
	"fmt"
	"testing"

	"github.com/ngbenci/calc"
)

func TestAdd(t *testing.T) {
	a, b := 1, 2
	res := calc.Sum(a, b)
	fmt.Println(res)

}
