package catalog_test

import (
	"encoding/json"
	"testing"

	"cafemovil-back/internal/catalog"
)

func TestExplicitPricesPreserveZeroAndLargeAmounts(t *testing.T) {
	for _, amount := range []float64{0, 9999999999.99} {
		encoded, err := json.Marshal(catalog.Product{Price: &amount, Currency: "COP"})
		if err != nil {
			t.Fatal(err)
		}
		var product map[string]any
		if err := json.Unmarshal(encoded, &product); err != nil {
			t.Fatal(err)
		}
		if product["price"] != amount {
			t.Fatalf("price changed: %s", encoded)
		}
	}
}

func TestMissingPriceIsNullNotZero(t *testing.T) {
	encoded, err := json.Marshal(catalog.Product{ID: "unpriced", Currency: "COP"})
	if err != nil {
		t.Fatal(err)
	}
	var product map[string]any
	if err := json.Unmarshal(encoded, &product); err != nil {
		t.Fatal(err)
	}
	price, present := product["price"]
	if !present || price != nil {
		t.Fatalf("missing price must be explicit null, got %s", encoded)
	}
}
