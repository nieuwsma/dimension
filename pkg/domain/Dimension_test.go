package domain

import (
	"errors"
	"fmt"
	"testing"
)

func TestValidateSpheres(t *testing.T) {
	tests := []struct {
		spheres     map[string]*Sphere
		expectedErr error
		description string
	}{
		{
			spheres: map[string]*Sphere{
				"a": NewSphere(Green),
				"b": NewSphere(Green),
				"c": NewSphere(Green),
				"d": NewSphere(Green),
			},
			expectedErr: fmt.Errorf("%s has 4 spheres, maximum is 3", Green.LongHand()),
			description: "Exceed color count",
		},
		{
			spheres: map[string]*Sphere{
				"a": NewSphere(Green),
				"b": NewSphere(Blue),
				"c": NewSphere(Black),
				"d": NewSphere(Orange),
			},
			expectedErr: nil,
			description: "Valid sphere colors",
		},
		//... add more test cases as needed
	}

	for _, test := range tests {
		d := &Dimension{Dimension: test.spheres}
		err := d.ValidateSpheres()
		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
		}
	}
}

func TestValidateGeometry(t *testing.T) {
	tests := []struct {
		spheres     map[string]*Sphere
		expectedErr error
		description string
	}{
		{
			spheres: map[string]*Sphere{
				"a": NewSphere(Green),
				"b": NewSphere(Blue),
				"c": NewSphere(Black),
				"d": NewSphere(Orange),
				"e": NewSphere(Green),
				"f": NewSphere(Blue),
				"g": NewSphere(Black),
				"h": NewSphere(Orange),
				"i": NewSphere(Green),
				"j": NewSphere(Blue),
				"k": NewSphere(Black),
				"l": NewSphere(Orange),
				"m": NewSphere(Green),
				"n": NewSphere(Blue),
			},
			expectedErr: errors.New("too many spheres"),
			description: "Too many spheres",
		},
		// ... other test cases
	}

	for _, test := range tests {
		d := &Dimension{Dimension: test.spheres}
		err := d.ValidateGeometry()
		if err != nil && err.Error() != test.expectedErr.Error() {
			t.Errorf("for %s, expected error %v, got %v", test.description, test.expectedErr, err)
		}
	}
}
