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

// Item is an entity with weight and cost that store
// in knapsack.
type Item struct {
	Cost   float64
	Weight float64
}

// Problem is a knapsack problem instance
// with gradian and objective function.
type Problem struct {
	Capacity float64 // knapsack capacity
	Items    []Item  // items that can be stored in knapsack

	Mu float64 // lagrangian relaxation
}

// Func evaluates the objective function at the given location. Func
// must not modify x.
func (p Problem) Func(x []float64) float64 {
	var sum float64

	for i, item := range p.Items {
		sum += item.Cost * x[i]
		sum -= p.Mu * (item.Weight * x[i])
		sum += p.Mu * (x[i]*x[i] - x[i])
		sum += p.Mu * (1 - x[i])
	}
	sum += p.Mu * p.Capacity

	return sum
}

// Grad evaluates the gradient at x and stores the result in-place in grad.
// Grad must not modify x.
func (p Problem) Grad(grad []float64, x []float64) {

	for i, item := range p.Items {
		grad[i] = 0
		grad[i] += item.Cost
		grad[i] -= p.Mu * item.Weight
		grad[i] += p.Mu * (2*x[i] - 1)
		grad[i] += p.Mu * (-1)
	}
}
