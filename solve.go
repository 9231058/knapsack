/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-08-2018
 * |
 * | File Name:     solve.go
 * +===============================================
 */

package knapsack

import (
	"github.com/1995parham/knapsack/problem"

	"gonum.org/v1/gonum/optimize"
)

// Solve solves given knapsack problem with gradient descent
// method
func Solve(p problem.Problem) (*optimize.Result, error) {
	op := optimize.Problem{
		Func: p.Func,
		Grad: p.Grad,
	}

	var x []float64
	x = make([]float64, len(p.Items))

	settings := optimize.DefaultSettingsLocal()
	settings.Recorder = nil
	settings.FunctionConverge = nil

	result, err := optimize.Minimize(op, x, settings, &optimize.GradientDescent{})
	if err != nil {
		return nil, err
	}
	if err = result.Status.Err(); err != nil {
		return nil, err
	}

	// correct objective function value
	result.F = -result.F

	return result, nil
}
