package main

import "fmt"

// Implemented simple search. There is room for improvement here.
func IsIn(item string, list []string) bool {
	for i := 0; i < len(list); i++ {
		if item == list[i] {
			return true
		}
	}

	return false
}

func And(a, b []string) []string {
	newSet := []string{}

	for i := 0; i < len(a); i++ {
		if IsIn(a[i], b) {
			newSet = append(newSet, a[i])
		}
	}

	return newSet
}

func Diff(a, b []string) []string {
	newSet := []string{}

	for i := 0; i < len(a); i++ {
		if !IsIn(a[i], b) {
			newSet = append(newSet, a[i])
		}
	}

	return newSet
}

func main() {

	statesToCover := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	stations := map[string][]string{
		"kone":   {"id", "nv", "ut"},
		"ktwo":   {"wa", "id", "mt"},
		"kthree": {"or", "nv", "ca"},
		"kfour":  {"nv", "ut"},
		"kfive":  {"ca", "az"},
	}

	finalStations := []string{}

	for len(statesToCover) > 0 {
		var bestStation *string       // melhor estação
		coveredStations := []string{} // estados_cobertos

		for station, states := range stations {
			covered := And(statesToCover, states)

			if len(covered) > len(coveredStations) {
				bestStation = &station
				coveredStations = covered
			}
		}

		statesToCover = Diff(statesToCover, coveredStations)
		finalStations = append(finalStations, *bestStation)
	}

	fmt.Printf("Final stations: %s\n", finalStations)

	// Important to mention:
	// Go's maps doesnt mantain an order to the keys, so sometimes, it can print 3 before 1.
	//
	// In our case here, not a big deal.
}
