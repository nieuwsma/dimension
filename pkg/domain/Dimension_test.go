package domain

import (
	"errors"
	"fmt"
	"strings"
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
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Green)},
			},
			expectedErr: nil,
			description: "Only center sphere pair",
		},
		{
			spherePairs: []SpherePair{
				{ID: "b", Sphere: *NewSphere(Green)},
				{ID: "c", Sphere: *NewSphere(Blue)},
				{ID: "d", Sphere: *NewSphere(Black)},
			},
			expectedErr: nil,
			description: "Outer ring without center",
		},
		{
			spherePairs: []SpherePair{
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "i", Sphere: *NewSphere(Blue)},
			},
			expectedErr: errors.New("missing center sphere"),
			description: "Tropical ring without center",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "i", Sphere: *NewSphere(Blue)},
			},
			expectedErr: errors.New("is missing a required neighbor"),
			description: "Tropical ring with center, missing required neighbors",
		},
		{
			spherePairs: []SpherePair{
				{ID: "n", Sphere: *NewSphere(Green)},
			},
			expectedErr: errors.New("missing center sphere"),
			description: "Top sphere without required equatorial and center spheres",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Green)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Black)},
				{ID: "d", Sphere: *NewSphere(Orange)},
				{ID: "e", Sphere: *NewSphere(Green)},
				{ID: "h", Sphere: *NewSphere(Blue)},
				{ID: "j", Sphere: *NewSphere(Black)},
				{ID: "l", Sphere: *NewSphere(Orange)},
				{ID: "n", Sphere: *NewSphere(Black)},
			},
			expectedErr: errors.New("is missing a required neighbor"),
			description: "Missing neighbor",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "d", Sphere: *NewSphere(Black)},
				{ID: "e", Sphere: *NewSphere(Black)},
				{ID: "f", Sphere: *NewSphere(White)},
				{ID: "g", Sphere: *NewSphere(White)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "i", Sphere: *NewSphere(Orange)},
				{ID: "j", Sphere: *NewSphere(Blue)},
			},
			expectedErr: errors.New("invalid tropical ring configuration"),
			description: "Invalid tropical ring configuration with center",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "d", Sphere: *NewSphere(Black)},
				{ID: "e", Sphere: *NewSphere(Black)},
				{ID: "f", Sphere: *NewSphere(White)},
				{ID: "g", Sphere: *NewSphere(White)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "k", Sphere: *NewSphere(Orange)},
			},
			expectedErr: nil,
			description: "Full valid configuration",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "d", Sphere: *NewSphere(Black)},
				{ID: "e", Sphere: *NewSphere(Black)},
				{ID: "f", Sphere: *NewSphere(White)},
				{ID: "g", Sphere: *NewSphere(White)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "j", Sphere: *NewSphere(Orange)},
				{ID: "l", Sphere: *NewSphere(Orange)},
			},
			expectedErr: nil,
			description: "Full valid configuration",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "e", Sphere: *NewSphere(Black)},
				{ID: "f", Sphere: *NewSphere(White)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "k", Sphere: *NewSphere(Orange)},
			},
			expectedErr: nil,
			description: "Full valid configuration",
		},
		{
			spherePairs: []SpherePair{
				{ID: "a", Sphere: *NewSphere(Black)},
				{ID: "b", Sphere: *NewSphere(Blue)},
				{ID: "c", Sphere: *NewSphere(Green)},
				{ID: "d", Sphere: *NewSphere(Black)},
				{ID: "e", Sphere: *NewSphere(Black)},
				{ID: "f", Sphere: *NewSphere(White)},
				{ID: "g", Sphere: *NewSphere(White)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "j", Sphere: *NewSphere(Orange)},
				{ID: "l", Sphere: *NewSphere(Orange)},
				{ID: "n", Sphere: *NewSphere(Orange)},
			},
			expectedErr: nil,
			description: "Full valid configuration",
		},
		{
			spherePairs: []SpherePair{
				{ID: "n", Sphere: *NewSphere(Green)},
				{ID: "a", Sphere: *NewSphere(Green)},
			},
			expectedErr: errors.New("missing a equitorial ring sphere neighbor b"),
			description: "Top sphere without required equatorial and center spheres",
		},
		{
			spherePairs: []SpherePair{
				{ID: "n", Sphere: *NewSphere(Green)},
				{ID: "h", Sphere: *NewSphere(Green)},
			},
			expectedErr: errors.New("missing center sphere"),
			description: "Top sphere without required equatorial and center spheres",
		},
		{
			spherePairs: []SpherePair{
				{ID: "n", Sphere: *NewSphere(Green)},
				{ID: "h", Sphere: *NewSphere(Green)},
				{ID: "a", Sphere: *NewSphere(Green)},
			},
			expectedErr: errors.New("missing a required neighbor"),
			description: "Top sphere without required equatorial and center spheres",
		},
		// ... other test cases
	}

	for _, test := range tests {

		d, _ := NewDimension(test.spherePairs...)
		err := d.ValidateGeometry()
		// Compare errors
		if err == nil && test.expectedErr == nil {
			continue
			//strings.Contains(err.Error(), "is missing a required neighbor")
		} else if err != nil && test.expectedErr != nil && strings.Contains(err.Error(), test.expectedErr.Error()) {
			continue
		} else {
			t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
		}
	}
}
