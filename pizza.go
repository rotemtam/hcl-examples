package hcl_examples

import (
	"fmt"
	"math"

	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/function"
)

type (
	Order struct {
		Contact *Contact `hcl:"contact,block"`
		Address *Address `hcl:"address,block"`
		Pizzas  []*Pizza `hcl:"pizza,block"`
	}
	Contact struct {
		Name  string `hcl:"name"`
		Phone string `hcl:"phone"`
	}
	Address struct {
		Street  string `hcl:"street"`
		City    string `hcl:"city"`
		Country string `hcl:"country"`
	}
	Pizza struct {
		Size     string   `hcl:"size"`
		Count    int      `hcl:"count,optional"`
		Toppings []string `hcl:"toppings,optional"`
	}
)

func ctx() *hcl.EvalContext {
	vars := make(map[string]cty.Value)
	for _, size := range []string{"S", "M", "L", "XL"} {
		vars[size] = cty.StringVal(size)
	}
	for _, topping := range []string{"olives", "onion", "feta_cheese", "garlic", "tomato"} {
		vars[topping] = cty.StringVal(topping)
	}
	// Define a the "for_diners" function.
	spec := &function.Spec{
		// Return a number.
		Type: function.StaticReturnType(cty.Number),
		// Accept a single input parameter, "diners", that is not-null number.
		Params: []function.Parameter{
			{Name: "diners", Type: cty.Number, AllowNull: false},
		},
		// The function implementation.
		Impl: func(args []cty.Value, _ cty.Type) (cty.Value, error) {
			d := args[0].AsBigFloat()
			if !d.IsInt() {
				return cty.NilVal, fmt.Errorf("expected int got %q", d)
			}
			di, _ := d.Int64()
			neededSlices := di * 3
			return cty.NumberFloatVal(math.Ceil(float64(neededSlices) / 8)), nil
		},
	}
	return &hcl.EvalContext{
		Variables: vars,
		Functions: map[string]function.Function{
			"for_diners": function.New(spec),
		},
	}
}
