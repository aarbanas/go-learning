package methods

import "fmt"

type Gesture struct {
	Name, Emoji string
}

func (g *Gesture) Update(name, emoji string) {
	g.Name = name
	g.Emoji = emoji
}

func MethodAction() {
	g := Gesture{
		Name:  "OK-Hand",
		Emoji: "👌",
	}

	fmt.Printf("The current gesture is: %s %s\n", g.Name, g.Emoji)

	g.Update("Thumbs-Up", "👍")

	fmt.Printf("The updated gesture is: %s %s\n", g.Name, g.Emoji)
}
