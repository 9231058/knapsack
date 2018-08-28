/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-08-2018
 * |
 * | File Name:     solve_test.go
 * +===============================================
 */

package knapsack

import (
	"testing"

	"github.com/1995parham/knapsack/problem"
)

func TestSolve(t *testing.T) {
	p := problem.Problem{
		Capacity: 100,
		Items: []problem.Item{
			problem.Item{
				Cost:   100,
				Weight: 90,
			},
			problem.Item{
				Cost:   10,
				Weight: 10,
			},
		},
		Mu: 0.001,
	}

	result, err := Solve(p)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("result.Status: %v\n", result.Status)
	t.Logf("result.X: %v\n", result.X)
	t.Logf("result.F: %v\n", result.F)
	t.Logf("result.Stats.FuncEvaluations: %d\n", result.Stats.FuncEvaluations)
}
