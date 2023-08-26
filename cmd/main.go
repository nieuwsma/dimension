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
// need to randomly draw a hand from the deck & track it
// need to create players
// need to track rounds (has a hand & players' dimensions)
// need to track game (has a draw deck, discard deck, players w/ scores)
