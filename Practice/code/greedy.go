package main

import "fmt"

type Score struct {
	score  int64
	result []Pizza
}

func solveGreedy(model Model) {
	var scores []Score
	pizza := model.Pizza
	for len(pizza) > 0 {
		score, result := solveGreedyInternal(model.MaxSlices, pizza)
		scores = append(scores, Score{score, result})
		pizza = pizza[:len(pizza)-1]
	}

	var max Score
	for _, score := range scores {
		if score.score > max.score {
			max = score
		}
	}

	result := Output{
		Pizza: max.result,
	}
	fmt.Println(result.String())
}

func solveGreedyInternal(max int64, pizza []Pizza) (score int64, result []Pizza) {
	for i := len(pizza) - 1; i >= 0; i-- {
		if max-pizza[i].Slices >= 0 {
			max -= pizza[i].Slices
			score += pizza[i].Slices
			result = append(result, pizza[i])
		}
	}
	return
}
