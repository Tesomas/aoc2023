package main

type Game struct {
	Id     int
	Rounds []Round
}

type Round struct {
	Red   int
	Blue  int
	Green int
}
