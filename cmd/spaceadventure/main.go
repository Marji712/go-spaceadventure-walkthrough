package main

import "go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	var ps = spaceadventure.PlanetarySystem{"Solare System"}
	spaceadventure.Start(ps)
}
