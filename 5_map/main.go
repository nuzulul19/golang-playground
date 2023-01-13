package main

func main() {
	// var colors map[string]string // return an empty map (map[])
	// colors := make(map[string]string) // return an empty map (map[])
	// colors["white"] = "#ffffff" // adding element to map
	// delete(colors, "white") // delete element of map

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		println(k, v)
	}
}
