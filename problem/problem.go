/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 28-08-2018
 * |
 * | File Name:     problem.go
 * +===============================================
 */

package problem

import "math"

// Item is an entity with weight and cost that store
// in knapsack.
type Item struct {
	Cost   float64
	Weight float64
}

// Problem is a knapsack problem instance
// with gradian and objective function.
// we formulate it as a minimization problem
// with logarithmic barrier and penalty methods.
type Problem struct {
	Capacity float64 // knapsack capacity
	Items    []Item  // items that can be stored in knapsack

	Mu float64 // lagrangian relaxation
}

// Func evaluates the objective function at the given location. Func
// must not modify x.
func (p Problem) Func(x []float64) float64 {
	var sum float64
	var penalty float64
	var barrier float64

	for i, item := range p.Items {
		sum -= item.Cost * x[i]
		barrier -= item.Weight * x[i]
		penalty += (x[i]*x[i] - x[i]) * (x[i]*x[i] - x[i])
	}
	barrier += p.Capacity

	return sum + (1/p.Mu)*penalty - (p.Mu)*math.Log(barrier)
}

// Grad evaluates the gradient at x and stores the result in-place in grad.
// Grad must not modify x.
func (p Problem) Grad(grad []float64, x []float64) {
	var weights float64

	for i, item := range p.Items {
		weights -= item.Weight * x[i]
	}
	weights += p.Capacity

	for i, item := range p.Items {
		grad[i] = 0
		grad[i] -= item.Cost
		grad[i] += p.Mu * item.Weight * (1 / weights)
		grad[i] += (1 / p.Mu) * 2 * (x[i]*x[i] - x[i]) * (2*x[i] - 1)
	}
}
