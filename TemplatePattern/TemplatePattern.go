package TemplatePattern

import "fmt"

// The template pattern is a behavioral design pattern that defines the skeleton of an algorithm in a method, deferring some steps to subclasses.
// It lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

type Game interface {
	Initialize()
	StartPlay()
	EndPlay()
	Play()
}

type Cricket struct{}

func (c Cricket) Initialize() {
	fmt.Println("Cricket Game Initialized! Start playing.")
}

func (c Cricket) StartPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c Cricket) EndPlay() {
	fmt.Println("Cricket Game Finished!")
}

func (c Cricket) Play() {
	c.Initialize()
	c.StartPlay()
	c.EndPlay()
}

type Football struct{}

func (f Football) Initialize() {
	fmt.Println("Football Game Initialized! Start playing.")
}

func (f Football) StartPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f Football) EndPlay() {
	fmt.Println("Football Game Finished!")
}

func (f Football) Play() {
	f.Initialize()
	f.StartPlay()
	f.EndPlay()
}
