package hcl_examples

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/stretchr/testify/require"
)

func TestOrder(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("testdata/order.hcl", nil, &o); err != nil {
		t.Fatalf("failed: %s", err)
	}
	require.EqualValues(t, Order{
		Contact: &Contact{
			Name:  "Sherlock Holmes",
			Phone: "+44 20 7224 3688",
		},
		Address: &Address{
			Street:  "221B Baker St",
			City:    "London",
			Country: "England",
		},
	}, o)
}

func TestPizza(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("testdata/pizza.hcl", ctx(), &o); err != nil {
		t.Fatalf("failed: %s", err)
	}
	require.EqualValues(t, Order{
		Pizzas: []*Pizza{
			{
				Size:  "XL",
				Count: 1,
				Toppings: []string{
					"olives",
					"feta_cheese",
					"onion",
				},
			},
		},
	}, o)
}

func TestDiners(t *testing.T) {
	var o Order
	if err := hclsimple.DecodeFile("testdata/diners.hcl", ctx(), &o); err != nil {
		t.Fatalf("failed: %s", err)
	}
	require.EqualValues(t, 2, o.Pizzas[0].Count)
}
