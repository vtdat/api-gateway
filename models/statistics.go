package models

import (
	"math"
	"sort"
)

type Quantile struct {
	Pool       []int   `json:"pool" binding:"required"`
	Percentile float64 `json:"percentile" binding:"required"`
}

func (q *Quantile) CalcQuantile() int {
	sort.Ints(q.Pool)
	totalObs := len(q.Pool)
	index := int(math.Floor(q.Percentile/100 * float64(totalObs + 1)))
	if index < 1 {
		return q.Pool[0]
	}
	if index > totalObs {
		return q.Pool[totalObs - 1]
	}
	return q.Pool[index-1]
}