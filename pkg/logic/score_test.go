package logic

import (
	"fmt"
	"testing"
)

func TestValidateRuleChecker(t *testing.T) {

	testCases, err := getTestCases()
	if err != nil {
		t.Errorf("Errors %v", err)
	}
	for _, testCase := range testCases {

		//tasks := Tasks(testCase.Tasks)
		score, bonus, taskViolations, errs := ScoreTurn(testCase.Tasks, testCase.ToDimension())

		if score != testCase.Score || bonus != testCase.Bonus {
			printErr := fmt.Sprintf("file: %v, expected score %v expected bonus %v, actual score %v, actual bonus %v. failed tasks %v, err %v", testCase.FileName, testCase.Score, testCase.Bonus, score, bonus, taskViolations, errs)
			t.Errorf("failed test case %s", printErr)

		}

	}
}
