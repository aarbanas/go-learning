package structs

import "fmt"

type Education struct {
	Faculty string
	City    string
}

type Entity struct {
	Name      string
	Age       int
	Height    int
	Education // now this field doesn't have a name, just the Education struct type
}

func main() {
	var homer Entity
	homer.Name = "Homer"
	homer.Age = 39
	homer.Height = 185

	// we can access the promoted fields of the nested struct Address directly:
	homer.City = "Rijeka"
	homer.Faculty = "RITEH"

	// we can print the promoted fields City and State directly:
	fmt.Println("City:", homer.City)       // City: Springfield
	fmt.Println("Faculty:", homer.Faculty) // State: Oregon
}
