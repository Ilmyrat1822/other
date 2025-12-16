package maps

import "fmt"

func main() {

	players := map[int]string{
		87: "Joao Neves",
		33: "Warren Zaire-Emery",
	}
	fmt.Println(players)

	players[2] = "Achraf Hakimi"
	players[10] = "Ousmane Dembele"
	players[11] = "Neymar"
	fmt.Println(players)
	delete(players, 11)
	fmt.Println(players)
}
