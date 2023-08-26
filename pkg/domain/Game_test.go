package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"testing"
)

// Define a struct to hold file information
type FileInfo struct {
	Player string
	Game   string
	Round  string
	File   string
}

func TestValidateGame(t *testing.T) {
	// Set directory path
	dirPath := "./pkg/domain/test_cases/" // Replace with your directory path
	// Read files from directory
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		os.Exit(1)
	}

	var fileInfoList []FileInfo

	// Iterate through each file
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			splitName := strings.Split(file.Name(), "_")
			if len(splitName) < 4 {
				continue
			}

			player := strings.Split(splitName[1], "-")[1]
			game := strings.Split(splitName[2], "-")[1]
			round := strings.Split(strings.Split(splitName[3], "-")[1], ".")[0]

			fileInfoList = append(fileInfoList, FileInfo{
				Player: player,
				Game:   game,
				Round:  round,
				File:   file.Name(),
			})
		}
	}

	//// Here you can access fileInfoList to get details and open files
	//for _, info := range fileInfoList {
	//	fmt.Println("Player:", info.Player, "| Game:", info.Game, "| Round:", info.Round, "| File:", info.File)
	//}

	// To open each file and process:
	//for _, info := range fileInfoList {
	//	data, err := os.ReadFile(path.Join(dirPath, info.File))
	//	if err != nil {
	//		fmt.Println("Error reading file:", err)
	//		continue
	//	}
	//
	//	// Process the data (Here I'm just printing it out, you can modify it)
	//	fmt.Println("Contents of", info.File, ":", string(data))
	//}

	// Assuming other code is as before

	// Your struct definitions ...

	// Create the games map
	games := make(map[int]Game)

	for _, info := range fileInfoList {
		gameNum, err := strconv.Atoi(info.Game)
		if err != nil {
			fmt.Println("Error converting game number to int:", err)
			continue
		}

		// If game doesn't exist in map, initialize it
		if _, exists := games[gameNum]; !exists {
			games[gameNum] = Game{
				Players: make(map[PlayerName]Player),
				Rounds:  make(map[int]Round), // Assuming a fixed size of 6 rounds per game
				Alive:   true,                // You can modify this as per your logic
			}
		}

		game := games[gameNum]

		// Populate player data
		playerName := PlayerName(info.Player)
		if _, exists := game.Players[playerName]; !exists {
			game.Players[playerName] = Player{
				PlayerName:  playerName,
				Turns:       make(map[int]Turn),
				ScoreRecord: ScoreRecord{}, // You can populate this based on your file data
			}
		}

		//player := game.Players[playerName]

		// Populate round data
		roundNum, err := strconv.Atoi(info.Round)
		if err != nil {
			fmt.Println("Error converting round number to int:", err)
			continue
		}

		// If the round data is empty or not resolved, populate it
		game.Rounds[roundNum-1] = Round{
			Resolved: false, // Modify this as per your logic or file data
		}

		// Here, you can populate other fields based on the actual contents of your JSON files
		// For example, to populate the turn data, you'll need to read the JSON content and extract the required details

		games[gameNum] = game

		data, err := os.ReadFile(path.Join(dirPath, info.File))
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}
		// Unmarshal the content into a TestCase struct
		var testCase TestCase
		err = json.Unmarshal(data, &testCase)
		if err != nil {
			//return nil, err
		}

		testCase.FileName = info.File
		fmt.Println(data)
	}

	// Print the games map for verification
	fmt.Println(games)

}
