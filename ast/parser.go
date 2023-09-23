package ast

import (
	"github.com/alecthomas/participle/v2"
)

type Value struct {
	Num  *int  `  "[" @Int "]"`
	Expr *Expr `| "(" @@ ")"`
}

type Expr struct {
	Or []*AndExpr `@@ ("OR" @@)*`
}

type AndExpr struct {
	And []*Value `@@ ("AND" @@)*`
}

func ValidateExpression(input string, valueMap map[int]bool) (bool, error) {
	parser, err := participle.Build[Expr]()

	if err != nil {
		panic(err)
	}

	expr, err := parser.ParseString("", input)

	if err != nil {
		panic(err)
	}

	return validate(expr, valueMap)

	// fmt.Printf("%#v\n", expr)
	// fmt.Printf("\n\n-- %#v --\n\n", )
}

func validate(expr *Expr, validationMap map[int]bool) (bool, error) {
	var finalAns *bool
	for _, orValue := range expr.Or {
		var val1 *bool
		if orValue.And == nil {
			continue
		}
		for _, andValue := range orValue.And {
			// This means that AND is just carrying a value
			var collectedValue bool
			if andValue.Num != nil {
				val := validationMap[*andValue.Num]
				collectedValue = val
			} else {
				val, err := validate(andValue.Expr, validationMap)
				if err != nil {
					return false, err
				}
				collectedValue = val
			}

			if val1 == nil {
				val1 = &collectedValue
			} else {
				res := *val1 && collectedValue
				val1 = &res
			}
		}

		if finalAns == nil {
			finalAns = val1
		} else {
			res := *finalAns || *val1
			finalAns = &res
		}

		// val, exists := validationMap[*orValue.And]
		// if !exists {
		// 	return false
		// }
		// if isValid == nil {
		// 	isValid = &val
		// } else {
		// 	*isValid = *isValid || val
		// }
	}
	return *finalAns, nil
}
