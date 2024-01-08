package slices

import "fmt"

func VariadicFunctionAction() {
	var prefix, name1, name2 string
	fmt.Scan(&prefix, &name1, &name2)

	for _, line := range greeting(prefix, name1, name2) {
		fmt.Println(line)
	}
}

func greeting(prefix string, name ...string) []string {
	var myslice []string
	for _, value := range name {
		var text = (prefix + " " + value)
		myslice = append(myslice, text)
	}

	return myslice
}
