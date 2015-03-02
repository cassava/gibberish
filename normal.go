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
	n := int(s)
	if n <= 0 {
		return 0
	}
	return n
}
