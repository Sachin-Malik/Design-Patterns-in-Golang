package main

import (
	"fmt"
)

// types

type iBuilder interface {
	setDoors(doorType string)
	setWall(wallCount int)
	setWindows(windowType string)
	setRoof(roofType string)
	setFlooring(flooringType string)
}

type House struct {
	doors    string
	walls    int
	windows  string
	roofType string
	flooring string
}

type NormalBuilder struct {
	doors    string
	walls    int
	windows  string
	roofType string
	flooring string
}

type IglooBuilder struct {
	doors    string
	walls    int
	windows  string
	roofType string
	flooring string
}

// functions

func (nb *NormalBuilder) setDoors(doorType string) {
	nb.doors = doorType
}

func (nb *NormalBuilder) setWall(wallCount int) {
	nb.walls = wallCount
}

func (nb *NormalBuilder) setWindows(windowType string) {
	nb.windows = windowType
}

func (nb *NormalBuilder) setRoof(roofType string) {
	nb.roofType = roofType
}

func (nb *NormalBuilder) setFlooring(flooringType string) {
	nb.flooring = flooringType
}

func (ib *IglooBuilder) setDoors(doorType string) {
	ib.doors = doorType
}

func (ib *IglooBuilder) setWall(wallCount int) {
	ib.walls = wallCount
}

func (ib *IglooBuilder) setWindows(windowType string) {
	ib.windows = windowType
}

func (ib *IglooBuilder) setRoof(roofType string) {
	ib.roofType = roofType
}

func (ib *IglooBuilder) setFlooring(flooringType string) {
	ib.flooring = flooringType
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doors:    b.doors,
		walls:    b.walls,
		windows:  b.windows,
		roofType: b.roofType,
		flooring: b.flooring,
	}
}

func (b *IglooBuilder) getHouse() House {
	return House{
		doors:    b.doors,
		walls:    b.walls,
		windows:  b.windows,
		roofType: b.roofType,
		flooring: b.flooring,
	}
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

// Builder

func getBuilder(houseType string) iBuilder {
	if houseType == "normal" {
		return newNormalBuilder()
	}
	if houseType == "igloo" {
		return newIglooBuilder()
	}
	return newNormalBuilder()
}

// Director

type Director struct {
	builder iBuilder
}

func newDirector(b iBuilder) Director {
	return Director{
		builder: b,
	}
}

func TestBuilderPatter() {
	// Initial configuration with walls only
	normalBuilder := newNormalBuilder()
	director := newDirector(normalBuilder)

	// Setting walls initially
	director.builder.setWall(1000)

	// Realizing that we need doors
	director.builder.setDoors("Wood")

	// Adding other components sequentially
	director.builder.setWindows("Glass")
	director.builder.setRoof("Shingles")
	director.builder.setFlooring("Hardwood")

	// Final house configuration
	normalHouse := normalBuilder.getHouse()
	fmt.Printf("House Configuration: %+v\n", normalHouse)
}
