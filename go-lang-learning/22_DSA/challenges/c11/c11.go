package main

func Cakes(recipe, available map[string]int) int {
	minCakes := -1

	for ing, req := range recipe {
		avail, exists := available[ing]
		if !exists || avail < req {
			return 0
		}
		possible := avail / req

		if minCakes == -1 || possible < minCakes {
			minCakes = possible
		}
	}
	if minCakes == -1 {
		return 0
	}
	return minCakes
}

func main() {
	print(Cakes(map[string]int{"flour": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
	print(Cakes(map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100}, map[string]int{"sugar": 500, "flour": 2000, "milk": 2000}))

}
