package main

import "go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	var ps = spaceadventure.PlanetarySystem{Name:"Solar System", Planets: []spaceadventure.Planet{
		spaceadventure.Planet{"Tatooine", "Desert planet"}},}
	spaceadventure.Start(ps)
}
