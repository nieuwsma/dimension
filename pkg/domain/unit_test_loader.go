package domain

import (
	"encoding/json"
	"os"
	"path/filepath"
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
			var user TestCase
			err = json.Unmarshal(content, &user)
			if err != nil {
				return nil, err
			}

			user.FileName = file.Name()

			// Append the user to the testCases slice
			testCases = append(testCases, user)
		}
	}

	return testCases, nil
}

// need to include a filename in the structure so i know where to trace it back to
func GetTestCases() (testCases []TestCase, err error) {
	//todo need to make this more robust!
	dir := "./pkg/domain/test_cases/" // Replace with your directory path
	print(os.Getwd())
	testCases, err = digestDirectory(dir)
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

func (t *TestCase) ToDimension() Dimension {

	var SpherePairs []SpherePair
	for k, v := range t.RawDimension {
		SpherePairs = append(SpherePairs, *NewSpherePair(SphereID(k), NewColorLong(v)))
	}
	dimension, _ := NewDimension(SpherePairs...)
	return *dimension
}
