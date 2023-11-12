package maps

import "fmt"

func ShowMaps() {
	countries := make(map[string]string)
	fmt.Println("\nCountries ", countries)

	countries["mexico"] = "D.F."
	countries["argentina"] = "Buenos Aires"
	fmt.Println(countries)
	fmt.Println(countries["argentina"])

	points := map[string]int{
		"Barcelona": 39,
		"Real Madrid": 38,
		"Chivas": 37,
		"Boca Juniors": 30,
	}

	fmt.Println("points: ", points)

	for club, score := range points {
		fmt.Printf("Team: %s, Has an score of: %d\n", club, score)
	}

	delete(points, "Real Madrid")
	fmt.Println(points)

	score, exists := points["Juventus"]
	fmt.Printf("The captured score is %d, and the club exists = %t\n", score, exists)

	score, exists = points["Chivas"]
	fmt.Printf("The captured score is %d, and the club exists = %t\n", score, exists)
}