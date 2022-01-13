package main

import (
	"fmt"
)

type elevator struct {
	stayFloor    int
	currentFloor int
	targetFloor  int
}

type person struct {
	requestFloor []int
}

func (p *person) request(floor ...int) []int {
	for _, f := range floor {
		p.requestFloor = append(p.requestFloor, f)
	}
	return p.requestFloor
}

func (e *elevator) stay(floor int) int {
	e.currentFloor = floor
	return e.currentFloor
}

func (e *elevator) move(targetFloor int) {
	fmt.Printf("电梯移动到了%d楼\n", targetFloor)
}

func (e *elevator) open() {
	fmt.Println("电梯开门")
}

func (e *elevator) close() {
	fmt.Println("电梯关门")
}

func (e *elevator) run(requestFloorSlice []int) {
	for _, f := range requestFloorSlice {
		e.move(f)
		e.open()
		e.close()
	}
}
