package elo_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/owenoclee/elo-go"
)

func TestCalculate(t *testing.T) {
	leftIn, rightIn := 1200.0, 1000.0
	{
		// Default k-factor and deviation (left win)
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.LeftWin)
		assertImpreciseEqual(t, 1207.7, leftOut)
		assertImpreciseEqual(t, 992.3, rightOut)
	}
	{
		// Default k-factor and deviation (right win)
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.RightWin)
		assertImpreciseEqual(t, 1175.7, leftOut)
		assertImpreciseEqual(t, 1024.3, rightOut)
	}
	{
		// Default k-factor and deviation (draw)
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.Draw)
		assertImpreciseEqual(t, 1191.7, leftOut)
		assertImpreciseEqual(t, 1008.3, rightOut)
	}
	{
		// Non-standard k-factor
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.LeftWin, elo.KFactor(10.0))
		assertImpreciseEqual(t, 1202.4, leftOut)
		assertImpreciseEqual(t, 997.6, rightOut)
	}
	{
		// Non-standard deviation
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.LeftWin, elo.Deviation(150.0))
		assertImpreciseEqual(t, 1201.4, leftOut)
		assertImpreciseEqual(t, 998.6, rightOut)
	}
	{
		// Non-standard k-factor and deviation
		leftOut, rightOut := elo.Calculate(leftIn, rightIn, elo.LeftWin, elo.KFactor(20.0), elo.Deviation(225.0))
		assertImpreciseEqual(t, 1202.3, leftOut)
		assertImpreciseEqual(t, 997.7, rightOut)
	}
}

func ExampleCalculate() {
	// Basic usage
	left, right := elo.Calculate(1200.0, 1000.0, elo.LeftWin)
	fmt.Printf("%.1f %.1f\n", left, right)

	// Using a custom K-factor and Deviation
	left, right = elo.Calculate(1200.0, 1000.0, elo.RightWin, elo.KFactor(10.0), elo.Deviation(200.0))
	fmt.Printf("%.1f %.1f\n", left, right)

	// Output: 1207.7 992.3
	// 1190.9 1009.1
}

func assertImpreciseEqual(t *testing.T, expected, actual float64) {
	if math.Abs(expected-actual) > 0.1 {
		t.Fatalf("expected %.1f, got %.1f", expected, actual)
	}
}
