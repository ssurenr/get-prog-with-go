package main

import (
	"fmt"
	"strings"
)

type location struct {
	lat, long float64
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
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

func main() {
	type rover struct {
		name, site	string
		position location
	}

	roverlist := []rover{
		{
			"Spirit",
			"Columbia Memorial Site",
			newLocation(
				coordinate{14, 34, 6.2, 'S'},
				coordinate{175, 28, 21.5, 'E'},
			),
		},
		{
			"Opportunity",
			"Challenger Memorial Station",
			newLocation(
				coordinate{1,56,46.3,'S'},
				coordinate{354,28,24.2,'E'},
				),
		},
		{
			name: "Curiosity",
			site: "Bradbury Landing",
			position: newLocation(
				coordinate{
					degree: 4,
					minute: 35,
					second: 22.2,
					h:      'S',
				},
				coordinate{137,26,30.1,'E'},
				),
		},
		{
			name: "InSight",
			site: "Elysium Planitia",
			position: newLocation(
				coordinate{4,30,0.0,'N'},
				coordinate{135,54,0,'E'},
				),
		},
	}

	header := fmt.Sprintf("%-20s %-30s %10s %10s","Rover or Lander","Landing Site","Lattitude","Longitute")
	fmt.Println(header)
	fmt.Println(strings.Repeat("-",len(header)))
	for _, item := range roverlist {
		fmt.Printf("%-20s %-30s %10.4f %10.4f\n",item.name,item.site,item.position.lat,item.position.long)
	}
}
