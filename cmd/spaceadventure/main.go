package main

import "go-spaceadventure-walkthrough/internal/spaceadventure"

func main() {
	spaceadventure.Start(
		spaceadventure.PlanetarySystem{Name:"Solar System", Planets: mockPlanets()},
	)
}

func mockPlanets() []spaceadventure.Planet {
	return []spaceadventure.Planet{
		spaceadventure.Planet{"Tatooine", "Desert planet"},
		spaceadventure.Planet{"Dagobah", "Swamp planet"},
		spaceadventure.Planet{"Hoth", "Icy planet"},
		spaceadventure.Planet{"Jaku", "Party planet"},
		spaceadventure.Planet{"Mustafar", "Fiery planet"},
		spaceadventure.Planet{"Coruscant", "Urban planet"},
		spaceadventure.Planet{"Bespin", "Cloudy planet"},
		spaceadventure.Planet{"Naboo", "Temple planet"},
		spaceadventure.Planet{"Kashyyyk", "Swampy planet"},
		spaceadventure.Planet{"Endor", "Ewok planet"},
		spaceadventure.Planet{"Alderaan", "Gone planet"},
	}
}