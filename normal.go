// Copyright (c) 2015, Ben Morgan. All rights reserved.
// Use of this source code is governed by an MIT license
// that can be found in the LICENSE file.

package gibberish

import (
	"math/rand"
	"time"
)

type Normal struct {
	Rand *rand.Rand // random number generator
	M    float64    // mean
	SD   float64    // standard deviation
}

func NewNormal(mean, standardDeviation float64) *Normal {
	if standardDeviation < 0 {
		return nil
	}

	return &Normal{
		Rand: rand.New(rand.NewSource(time.Now().UnixNano())),
		M:    mean,
		SD:   standardDeviation,
	}
}

func (n *Normal) Sample() int {
	s := n.Rand.NormFloat64()*n.SD + n.M
	m := int(s)
	if m <= 0 {
		return 0
	}
	return m
}
