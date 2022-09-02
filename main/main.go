package main

import (
	"log"
	"os"

	"exp/blackpoly/game"
	"exp/blackpoly/telegram"
	"fmt"

	"encoding/json"
	// "fmt"
	// "log"
	// "fmt"
	
)

func main() {

	telegram.TeleRun()
	// la := func(){
	// 	for i := 0; i < 1000; i++{
	// 	fmt.Println(i)
	// 	}
	// }
	// go la()

	// dsp := echotron.NewDispatcher(token, newBot)
	// log.Println(dsp.Poll())
	// l1 := "░"
	// // l2 := "▒"
	// // l3 := "▓"
	// l4 := "█"

	r, err := os.ReadFile("card.json")
	if err != nil {
		log.Println(err)
	}
	Map := make(map[string][]interface{})
	json.Unmarshal(r, &Map)
	for _, i := range Map["property"] {
		// fmt.Println(i.color)
		fmt.Printf("\n\n")
		j, err := json.Marshal(i)
		if err != nil {
			log.Println(err)
		}
		Cd := game.CardBase{}

		json.Unmarshal(j, &Cd)

		// Td := game.TitleDeedCard{}
		// Ud := game.UtilityCard{}
		// Trd := game.TrainCard{}
		// Taxd := game.TaxCard{}
		// Cod := game.CommunityCard{}
		// Chd := game.ChanceCard{}
		// Jd := game.Jailcard{}
		// Fd := game.FreePark{}
		// Gd := game.GoToJail{}

	}
}
