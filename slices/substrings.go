package slices

import "fmt"

func CreateSubstrings() {
	words := []string{"saltwater", "backstage", "bedrock", "sandcastle", "snowman"}

	w1 := words[0][4:9]
	w2 := words[1][4:9]
	w3 := words[2][3:7]
	w4 := words[3][4:10]
	w5 := words[4][4:7]

	fmt.Println(w1, w2, w3, w4, w5)
}
