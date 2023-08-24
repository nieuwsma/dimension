package domain

import (
	"fmt"
	"testing"
)

func TestValidateRuleChecker(t *testing.T) {

	testCases, err := GetTestCases()
	if err != nil {
		t.Errorf("Errors %v", err)
	}
	for _, testCase := range testCases {

		score, bonus, _ := ScoreTurn(testCase.Tasks, testCase.ToDimension())

		if score != testCase.Score && bonus != testCase.Bonus {
			printErr := fmt.Sprintf("file: %v, expected score %v expecte bonus %v, actual score %v, actual bonus %v", testCase.FileName, score, bonus, testCase.Score, testCase.Bonus)
			t.Errorf("failed test case %s", printErr)

		}

	}

	//for _, test := range tests {
	//	dimMap := make(map[string]Sphere)
	//	for _, pair := range test.spherePairs {
	//		dimMap[string(pair.ID)] = pair.Sphere
	//	}
	//	d := &Dimension{Dimension: dimMap}
	//	err := d.ValidateSpheres()
	//	if err != nil && err.Error() != test.expectedErr.Error() {
	//		t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
	//	}
	//}
}
