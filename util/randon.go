package util

import (
	"math/rand"
	"time"
)

type random struct{}

var Random random

func (r *random) GenerateNumber() int {
	crank := rand.New(rand.NewSource(time.Now().UnixNano()))
	return crank.Int()
}
