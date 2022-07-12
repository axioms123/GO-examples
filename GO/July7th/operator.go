package main

import (
	"fmt"
	"math"
)

type Operator func() {

}


func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}


}
