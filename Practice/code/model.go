package main

/*
Num Slices Maximum, Num Types
17 4
2 5 6 8

0 1 2 3
Type with Number of Slices
*/

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	MaxSlices  int64
	PizzaTypes int64

	NumSlices []int64
}

func FileFrom(name string) (file File) {
	lines := getLinesFromInput(name)
	fmt.Fscanf(bytes.NewReader([]byte(lines[0])), "%d %d", &file.MaxSlices, &file.PizzaTypes)
	for _, n := range strings.Split(lines[1], " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		file.NumSlices = append(file.NumSlices, int64(num))
	}
	return
}

type Model struct {
	MaxSlices int64
	Pizza     []Pizza
}

type Pizza struct {
	Id     int64
	Slices int64
}

func ModelFromFile(file File) (model Model) {
	model.MaxSlices = file.MaxSlices
	for i := int64(0); i < file.PizzaTypes; i++ {
		model.Pizza = append(model.Pizza, Pizza{
			Id:     i,
			Slices: file.NumSlices[i],
		})
	}
	return
}

type Output struct {
	Pizza []Pizza
}

func (o Output) String() string {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(len(o.Pizza)))
	buf.WriteRune('\n')
	for _, p := range o.Pizza {
		buf.WriteString(fmt.Sprintf("%v ", p.Id))
	}
	return buf.String()
}
func (o Output) Score() (score int64) {
	for _, pizza := range o.Pizza {
		score += pizza.Slices
	}
	return
}
