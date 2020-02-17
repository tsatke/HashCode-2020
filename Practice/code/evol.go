package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/*

T0 T1 T2 ... Tn (Tn ∈ {0,1})
P0 P1 P2 ... Pn (Pn ∈ N)
f(_) = SUM of 0 .. n on k( Tk * Pk )

M >= f(_)

Mutators
* Flip
* OR

*/

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	survivors []Vec
)

const (
	iterations          = 10000
	sampleSize          = 100
	deathFactor         = .7
	mutationProbability = .2
)

// T1 T... Tn {0,1}
func runEvol(model Model) {
	v := make(Vec, len(model.Pizza))

	for i := 0; i < sampleSize; i++ {
		// survivors = append(survivors, RandomVec(len(model.Pizza)))
		survivors = append(survivors, v)
	}

	// actually compute something
	for i := 0; i < iterations; i++ {
		doIteration(model)
	}

	survivors = computeSurvivors(survivors, fitness, 1, model)

	output := Output{}
	for i, s := range survivors[0] {
		if s {
			output.Pizza = append(output.Pizza, model.Pizza[i])
		}
	}
	fmt.Printf("%v\n", output.String())
}

func doIteration(model Model) {
	survivors = computeSurvivors(survivors, fitness, sampleSize*deathFactor, model)
	var children []Vec
	limit := int(sampleSize * (1 - deathFactor))
	for i := 0; i < limit; i += 2 {
		v1, v2 := crossover(survivors[rand.Intn(len(survivors))], survivors[rand.Intn(len(survivors))])
		children = append(children, v1, v2)
	}

	for i := range survivors {
		if rand.Float32() <= mutationProbability {
			survivors[i] = randomMutation(survivors[i], survivors[rand.Intn(len(survivors))], model)
		}
	}

	survivors = append(survivors, children...)
}

func randomMutation(v1, v2 Vec, model Model) Vec {
	switch rand.Intn(4) {
	case 0:
		return or(v1, v2) // OR
	case 1:
		return and(v1, v2)
	case 2:
		return flip(v1, model)
	case 3:
		return nxor(v1, v2)
	}
	return v1
}

type Vec []bool

func (v Vec) ToInt64Slice() (res []int64) {
	for _, b := range v {
		res = append(res, b2i64(b))
	}
	return
}

func b2i64(b bool) int64 {
	if b {
		return 1
	}
	return 0
}

func computeSurvivors(sample []Vec, fitness func(Vec, Model) int64, newSampleSize int, model Model) []Vec {
	sort.Slice(sample, func(i, j int) bool {
		return fitness(sample[i], model) < fitness(sample[j], model)
	})
	return sample[len(sample)-newSampleSize:]
}

func fitness(mutant Vec, model Model) int64 {
	var sum int64

	for i, v := range mutant.ToInt64Slice() {
		sum = v * model.Pizza[i].Slices
	}

	if sum <= model.MaxSlices {
		return sum
	}
	return 0
}

func crossover(m1, m2 Vec) (Vec, Vec) {
	pivotStart := rand.Intn(len(m1))
	pivotEnd := rand.Intn(len(m1))
	if pivotStart > pivotEnd {
		pivotStart, pivotEnd = pivotEnd, pivotStart
	}

	for i := pivotStart; i < pivotEnd; i++ {
		m1[i], m2[i] = m2[i], m1[i]
	}

	return m1, m2
}

func nxor(m1, m2 Vec) Vec {
	return binaryMutate(m1, m2, func(b1, b2 bool) bool { return b1 == b2 })
}
func or(m1, m2 Vec) Vec {
	return binaryMutate(m1, m2, func(b1, b2 bool) bool { return b1 || b2 })
}

func and(m1, m2 Vec) Vec {
	return binaryMutate(m1, m2, func(b1, b2 bool) bool { return b1 && b2 })
}

func binaryMutate(m1, m2 Vec, mutator func(bool, bool) bool) (v Vec) {
	for i := range m1 {
		v = append(v, mutator(m1[i], m2[i]))
	}
	return
}

func unaryMutate(m1 Vec, mutator func(bool) bool) (v Vec) {
	for i := range m1 {
		v = append(v, mutator(m1[i]))
	}
	return
}

func flip(mutant Vec, model Model) Vec {
	pos := rand.Intn(len(mutant))

	if rand.Int63n(model.MaxSlices) > model.Pizza[mutant.ToInt64Slice()[pos]].Slices {
		mutant[pos] = !mutant[pos]
	}

	return mutant
}

func shuffle(mutant Vec) Vec {
	return unaryMutate(mutant, func(b bool) bool {
		return rand.Intn(2) == 0
	})
}
