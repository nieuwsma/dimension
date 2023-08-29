package main

import (
	"dimension/pkg/domain"
	"time"
)

func main() {

	//dim, err := domain.NewDimension(*domain.NewSpherePair(domain.A, domain.Black), *domain.NewSpherePair(domain.B, domain.Black))
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(dim)
	//
	//u, err := json.MarshalIndent(domain.GetGeometry(), "", "  ")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(u))

	game := domain.NewGame(6, 60*time.Second, 1234234)
	game.AddPlayer("Andrew")
	game.AddPlayer("Jessica")
	tasks, _ := game.NextRound()
	game.Deck.Shuffle()

	c, r, _ := game.GetCurrentRound()
	round := game.Rounds[c-1]
	round.Tasks = nil
	game.Rounds[c-1] = round
	if len(r.Tasks) > len(tasks) {
		print("mismatch")
	}
	_ = game.EndRound(false)
	//print(err)
}

//todo
//need to create an expose an API;
//probably need to create a VERY VERY VERY dumb AI; like only obeys the quantity; places it whereever (that way I can generate some more test data)
//need to think about creating a CLI to interact with it? can it remember game session identifier? etc?
//I think player and game needs tokens to protect them from unlawful submissions
