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
	"fmt"
	"log"

	"github.com/1995parham/knapsack/problem"

	"gonum.org/v1/gonum/optimize"
)

// Sovle given knapsack problem with gradient descent
// method
func Solve(p problem.Problem) {
	p := optimize.Problem{
		Func: p.Func,
		Grad: p.Grad,
	}

	var x []int
	x = make([]int, len(p.Items))

	settings := optimize.DefaultSettingsLocal()
	settings.Recorder = nil
	settings.GradientThreshold = 1e-12
	settings.FunctionConverge = nil

	result, err := optimize.Maximize(p, x, settings, &optimize.GradientDescent{})
	if err != nil {
		log.Fatal(err)
	}
	if err = result.Status.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("result.Status: %v\n", result.Status)
	fmt.Printf("result.X: %v\n", result.X)
	fmt.Printf("result.F: %v\n", result.F)
	fmt.Printf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations)
}
