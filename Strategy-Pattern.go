package main

import (
	"fmt"
	"strconv"
)

type SafetyPlacer interface {
	placeSafties(string)
}

type Climber struct {
	rocksClimbed int
	sp           SafetyPlacer
}

type IceSafetyPlacer struct {
	temp int
	// db
	// api
	// data
}

// IceSafetyPlacer Struct Implementing the placeSafties, to full fill the the interface requirement
func (sp *IceSafetyPlacer) placeSafties(message string) {
	if sp.temp > 28 {
		fmt.Println("Temperature is to high to place safeties, Please climb down")
		fmt.Println("Waiting for temp to cool down,")
		fmt.Println("Okay, safe to climb now")
		sp.temp = 11
	} else {
		fmt.Println("Placing Safety after Ice Rocks Climbed", message)
	}
}

type StoneSafetyPlacer struct {
	density float64
	// on do some minimal thing
}

// StoneSafetyPlacer Struct Implementing the placeSafties, to full fill the the interface requirement
func (sp *StoneSafetyPlacer) placeSafties(message string) {
	if sp.density > 43 {
		fmt.Println("You need special tools which you do not have")
	} else {
		fmt.Println("Placing Safety after Stone Rocks Climbed", message)
	}
}

func (climber *Climber) climbRock() int {
	climber.rocksClimbed += 1
	if climber.rocksClimbed > 10 {
		climber.sp.placeSafties(strconv.Itoa(climber.rocksClimbed))
	}
	return climber.rocksClimbed
}

// it will take a struct that implements SafePlacer Interface
func newRockClimber(sp SafetyPlacer) *Climber {
	rc := &Climber{rocksClimbed: 0, sp: sp}
	return rc
}
