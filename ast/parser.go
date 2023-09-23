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

func ValidateExpression(input string, valueMap map[int]bool) bool {
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

func validate(expr *Expr, validationMap map[int]bool) bool {
	var isValid *bool
	for _, orValue := range expr.Or {
		var answer *bool
		if orValue.And == nil {
			continue
		}
		for _, andValue := range orValue.And {
			// This means that AND is just carrying a value
			var ans bool
			if andValue.Num != nil {
				val := validationMap[*andValue.Num]
				if answer == nil {
					ans = val
				} else {
					ans = *answer && val
				}
			} else {
				val := validate(andValue.Expr, validationMap)
				if answer == nil {
					answer = &val
				} else {
					ans = *answer && val
				}

			}
			if answer == nil {
				answer = &ans
			} else {
				res := *answer && ans
				answer = &res
			}
		}
		if isValid == nil {
			isValid = answer
		} else {
			res := *isValid || *answer
			isValid = &res
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
	return *isValid
}
