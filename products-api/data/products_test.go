package data

import (
	"testing"
)

func TestValidation(t *testing.T) {
	p := &Product{
		Name:  "jam",
		Price: 3.0,
		SKU: "ab-cd-ef",
	}

	err := p.Validator()

	if err != nil {
		t.Fatal(err)
	}
}
