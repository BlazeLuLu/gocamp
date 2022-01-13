package main

import (
	"testing"
)

func TestCase1(t *testing.T) {
	var e elevator
	//var p person
	currentFloor := e.stay(3)
	//requestFloorSlice := p.request()
	if currentFloor != 3 {
		t.Fatalf("预期电梯停在3楼，但是停在了%d楼。", e.currentFloor)
	}
}

func TestCase2(t *testing.T) {
	var e elevator
	var p person
	currentFloor := e.stay(1)
	if currentFloor != 1 {
		t.Fatalf("电梯初始应停在1楼，但是停在了%d楼。", e.currentFloor)
	}
	requestFloorSlice := p.request(3)
	e.run(requestFloorSlice)

}

func TestCase3(t *testing.T) {
	var e elevator
	var p person
	currentFloor := e.stay(3)
	if currentFloor != 3 {
		t.Fatalf("电梯初始应停在1楼，但是停在了%d楼。", e.currentFloor)
	}
	requestFloorSlice := p.request(4, 2)
	e.run(requestFloorSlice)
}

func TestCase4(t *testing.T) {
	var e elevator
	var p person
	currentFloor := e.stay(3)
	if currentFloor != 3 {
		t.Fatalf("电梯初始应停在1楼，但是停在了%d楼。", e.currentFloor)
	}
	requestFloorSlice := p.request(4, 5, 2)
	e.run(requestFloorSlice)
}
