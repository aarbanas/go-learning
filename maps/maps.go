package maps

import "fmt"

func ExistsInMap(key string) {
	// DO NOT delete the code block below!
	fruits := map[string]string{"pear": "🍐", "apple": "🍎", "banana": "🍌"}

	// Write below the code to check if the input key exists in the map:
	// Try to use `ok` as the check variable name to follow the Go code style conventions!
	if val, ok := fruits[key]; ok {
		fmt.Println("The fruit", key, val, "is in the map")
		return
	}
	fmt.Println("The fruit", key, "is not in the map")
}
