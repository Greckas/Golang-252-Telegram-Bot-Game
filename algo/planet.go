package algo

import (
	"math"

	"github.com/bohdanlisovskyi/Golang-252-Telegram-Bot-Game/database/postgres"
)

//PlaceNewPlanet - return coordinates of new planet
func PlaceNewPlanet(amountOfPlanets int) (int, int) {
	if amountOfPlanets == 0 {
		return 0, 0
	}
	flatOfSpiral := 1.1
	cornerCoefficient := 16.0
	fAmountOfPlanets := float64(amountOfPlanets)
	//Archimedean spiral equation with some  modifications
	xCoordinate := flatOfSpiral * (fAmountOfPlanets + 1) * math.Cos(cornerCoefficient/math.Abs(math.Log10(fAmountOfPlanets+1))*math.Sqrt(fAmountOfPlanets+1))
	yCoordinate := flatOfSpiral * (fAmountOfPlanets + 1) * math.Sin(cornerCoefficient/math.Abs(math.Log10(fAmountOfPlanets+1))*math.Sqrt(fAmountOfPlanets+1))
	return int(xCoordinate), int(yCoordinate)
}

//FindNeighbors - return first "amount" neighbors of planet
func FindNeighbors(mainPlanet *postgres.Planet, listOfPlanets []postgres.Planet, amount int) []postgres.Planet {
	var listOfNeighbors = []postgres.Planet{}
	minRadius := 1
	maxRadius := 4
	increase := maxRadius - minRadius
	for len(listOfNeighbors) < amount && len(listOfNeighbors) < len(listOfPlanets)-1 {
		for _, planet := range listOfPlanets {
			dist := ManhDist(planet.XCoordinate, planet.YCoordinate, mainPlanet.XCoordinate, mainPlanet.YCoordinate)
			if dist >= minRadius && dist < maxRadius {
				listOfNeighbors = append(listOfNeighbors, planet)
				if len(listOfNeighbors) == amount {
					return listOfNeighbors
				}
			}
		}
		minRadius += increase
		maxRadius += increase
	}
	return listOfNeighbors
}