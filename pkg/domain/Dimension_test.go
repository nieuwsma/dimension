package domain

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateSpheres(t *testing.T) {
	tests := []struct {
		spherePairs []SpherePair
		expectedErr error
		description string
	}{
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Green)},
				{ID: "b", Sphere: *NewSphere(Green)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "d", Sphere: *NewSphere(Green)},
			},
			expectedErr: fmt.Errorf("%s has 4 spheres, maximum is 3", Green.LongHand()),
			description: "Exceed color count",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Green)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Black)},
				{ID: "d", Sphere: *NewSphere(Orange)},
			},
			expectedErr: nil,
			description: "Valid sphere colors",
		},
		//... add more test cases as needed
	}

	for _, test := range tests {
		dimMap := make(map[string]Sphere)
		for _, pair := range test.spherePairs {
			dimMap[string(pair.ID)] = pair.Sphere
		}
		d := &Dimension{Dimension: dimMap}
		err := d.ValidateSpheres()
		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
		}
	}
}

func TestValidateGeometry(t *testing.T) {
	tests := []struct {
		spherePairs []SpherePair
		expectedErr error
		description string
	}{
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Green)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Black)},
				{ID: "d", Sphere: *NewSphere(Orange)},
				{ID: "e", Sphere: *NewSphere(Green)},
				{ID: "f", Sphere: *NewSphere(Blue)},
				{ID: "g", Sphere: *NewSphere(Black)},
				{ID: "h", Sphere: *NewSphere(Orange)},
				{ID: "i", Sphere: *NewSphere(Green)},
				{ID: "j", Sphere: *NewSphere(Blue)},
				{ID: "k", Sphere: *NewSphere(Black)},
				{ID: "l", Sphere: *NewSphere(Orange)},
				{ID: "m", Sphere: *NewSphere(Green)},
				{ID: "n", Sphere: *NewSphere(Blue)},
			},
			expectedErr: errors.New("too many spheres"),
			description: "Too many spheres",
		},
		// ... other test cases
	}

	for _, test := range tests {
		dimMap := make(map[string]Sphere)
		for _, pair := range test.spherePairs {
			dimMap[string(pair.ID)] = pair.Sphere
		}
		d := &Dimension{Dimension: dimMap}
		err := d.ValidateGeometry()
		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
		}
	}
}
