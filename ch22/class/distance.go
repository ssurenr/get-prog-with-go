package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"math"
)

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func rad(degree float64) float64 {
	return degree * math.Pi / 180
}

type world struct {
	radius float64
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

type coordinate struct {
	degree, minute, second float64
	h                      rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.degree + c.minute/60 + c.second/3600)
}

type rover struct {
	name, site string
	position   location
}

var (
	earth = world{radius: 6371.0}
	mars  = world{radius: 3389.5}
)

func main() {

	roverInfo := []rover{
		{
			name: "Spirit",
			site: "Columbia Memorial Site",
			position: newLocation(
				coordinate{14, 34, 6.2, 'S'},
				coordinate{175, 28, 21.5, 'E'},
			),
		},
		{
			name: "Opportunity",
			site: "Challenger Memorial Station",
			position: newLocation(
				coordinate{1, 56, 46.3, 'S'},
				coordinate{354, 28, 24.2, 'E'},
			),
		},
		{
			name: "Curiosity",
			site: "Bradbury Landing",
			position: newLocation(
				coordinate{4, 35, 22.2, 'S'},
				coordinate{137, 26, 30.1, 'E'},
			),
		},
		{
			name: "InSight",
			site: "Elysium Planitia",
			position: newLocation(
				coordinate{4, 30, 0.0, 'N'},
				coordinate{135, 54, 0, 'E'},
			),
		},
	}
	combinations := combin.Combinations(len(roverInfo), 2)

	for _, pair := range combinations {
		site1 := roverInfo[pair[0]]
		site2 := roverInfo[pair[1]]

		displayDistance(mars, site1.name, site2.name, site1.position, site2.position)
	}

	fmt.Println("")

	londonLocation := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})
	parisLocation := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})
	displayDistance(earth, "London", "Paris", londonLocation, parisLocation)

	myCityLocation := newLocation(coordinate{51, 87, 0, 'N'}, coordinate{0, 15, 0, 'E'})
	myCapitalLocation := londonLocation
	displayDistance(earth, "My City", "My Capital", myCityLocation, myCapitalLocation)

	mountSharpLocation := newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMonsLocation := newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	displayDistance(mars, "Mount Sharp", "Olympus Mons", mountSharpLocation, olympusMonsLocation)

}

func displayDistance(w world, s1name string, s2name string, site1 location, site2 location) {
	content := fmt.Sprintf("Distance between %s and %s:", s1name, s2name)
	fmt.Printf("%-50s %.4f\n", content, w.distance(site1, site2))
}
