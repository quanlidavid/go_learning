package client

import (
	"go_learning/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSerie(10))
	t.Log(series.Square(5))
}
