package utils

import (
	"math/rand"
	"time"
)

func GenerateUUID() string {
	return Generate(8) + "-" + Generate(4) + "-" + Generate(4) + "-" + Generate(4) + "-" + Generate(12)
}

func Generate(l int) (str string) {
	return newGenerator().Generate(l)
}

func newGenerator(flags ...int) *generator {
	src := []int{}
	for i := 0; i < 123; i++ {
		src = append(src, i)
	}
	charMap := src[97:123]
	return &generator{charMap, int64(len(charMap)), 0}
}

type generator struct {
	charMap []int
	length  int64
	val     int64
}

func (g *generator) Generate(l int) (str string) {
	for i := 0; i < l; i++ {
		str += g.randSymbol()
	}
	return
}

func (g *generator) randSymbol() string {
	rand.Seed(time.Now().UnixNano() + g.nextValue())
	return string(g.charMap[rand.Int63n(g.length)])
}

func (g *generator) nextValue() int64 {
	(*g).val++
	if g.val > 9999 {
		(*g).val = 0
	}
	return g.val
}
