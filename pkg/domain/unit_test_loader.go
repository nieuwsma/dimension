package domain

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// ... (TestCase struct definition goes here)

func digestDirectory(dir string) ([]TestCase, error) {
	var testCases []TestCase

	// List all files in the directory
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Iterate through each file
	for _, file := range files {
		// Check if the file has a .json extension
		if filepath.Ext(file.Name()) == ".json" {
			// Read the file content
			content, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}

			// Unmarshal the content into a TestCase struct
			var testCase TestCase
			err = json.Unmarshal(content, &testCase)
			if err != nil {
				return nil, err
			}

			testCase.FileName = file.Name()

			// Append the testCase to the testCases slice
			testCases = append(testCases, testCase)
		}
	}

	return testCases, nil
}

// need to include a filename in the structure so i know where to trace it back to
func getTestCases() (testCases []TestCase, err error) {
	//todo need to make this more robust!
	dir := "./pkg/domain/test_cases/" // Replace with your directory path
	print(os.Getwd())
	testCases, err = digestDirectory(dir)
	return
}

func getGamesTestCases() (games map[int]Game) {
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

	// Create the games map
	games = make(map[int]Game)

	for _, info := range fileInfoList {
		gameNum, err := strconv.Atoi(info.Game)
		if err != nil {
			fmt.Println("Error converting game number to int:", err)
			continue
		}

		// If game doesn't exist in map, initialize it
		if _, exists := games[gameNum]; !exists {

			games[gameNum] = *NewGame(6, 60*time.Second, 12345)
		}

		game := games[gameNum]

		// Populate player data
		playerName := PlayerName(info.Player)
		if _, exists := game.Players[playerName]; !exists {
			game.Players[playerName] = NewPlayer(playerName)
		}

		// Populate round data
		roundNum, err := strconv.Atoi(info.Round)
		if err != nil {
			fmt.Println("Error converting round number to int:", err)
			continue
		}

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

		// Here, you can populate other fields based on the actual contents of your JSON files
		// For example, to populate the turn data, you'll need to read the JSON content and extract the required details
		var r Round
		r.Tasks = testCase.Tasks
		r.Resolved = true
		p := game.Players[playerName]
		t := Turn{
			Dimension:      testCase.ToDimension(),
			Score:          testCase.Score,
			Bonus:          testCase.Bonus,
			TaskViolations: nil,
			FileName:       info.File,
		}
		p.Turns[roundNum] = t
		p.ScoreRecord.Points += t.Score
		if t.Bonus {
			p.ScoreRecord.BonusTokens++

		}
		game.Players[playerName] = p

		game.Rounds[roundNum] = r
		game.Alive = false

		games[gameNum] = game
	}

	return
}

type TestCase struct {
	FileName     string
	Name         string            `json:"name"`
	Bonus        bool              `json:"bonus"`
	Score        int               `json:"score"`
	Tasks        Tasks             `json:"tasks"`
	RawDimension map[string]string `json:"dimension"`
}

// Define a struct to hold file information
type FileInfo struct {
	Player string
	Game   string
	Round  string
	File   string
}

func (t *TestCase) ToDimension() Dimension {

	var SpherePairs []SpherePair
	for k, v := range t.RawDimension {
		SpherePairs = append(SpherePairs, *NewSpherePair(SphereID(k), NewColorLong(v)))
	}
	dimension, _ := NewDimension(SpherePairs...)
	return *dimension
}
