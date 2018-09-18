package matchers

import (
	"context"
	"strings"
)

type strategyContextKey string

func (c strategyContextKey) String() string {
	return "go-concurrency-limits|strategy|" + string(c)
}

// StringPredicateContextKey is the StringPredicate context key
// use this in your context.Context
const StringPredicateContextKey = strategyContextKey("stringPredicate")

// StringPredicateMatcher implements the string predicate matcher.
func StringPredicateMatcher(matchString string, caseInsensitive bool) func(ctx context.Context) bool {
	return func(ctx context.Context) bool {
		val := ctx.Value(StringPredicateContextKey)
		if val != nil {
			strVal, ok := val.(string)
			if ok {
				if caseInsensitive {
					return strings.ToLower(strVal) == strings.ToLower(matchString)
				}
				return strVal == matchString
			}
		}
		return false
	}
}