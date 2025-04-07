package efmt

import (
	"errors"

	"github.com/samber/lo"
)

func ValuePairs(err error) []KeyValue {
	var res []KeyValue

	for err := err; err != nil; {
		if e, ok := err.(*e); ok {
			res = append(res, e.Values()...)
		}

		err = errors.Unwrap(err)
	}

	return res
}

func Values(err error) map[string]any {
	// reversing to get the most upper (re-written) value of the same key
	return lo.FromEntries(lo.Reverse(ValuePairs(err)))
}

func Value[T any](err error, key string) (T, bool) {
	for _, kv := range ValuePairs(err) {
		if kv.Key == key {
			if val, ok := kv.Value.(T); ok {
				return val, true
			}
		}
	}

	return lo.Empty[T](), false
}
