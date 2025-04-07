package efmt

import (
	"errors"

	"github.com/samber/lo"
)

// ValuePairs extracts all key-value pairs from an error and its wrapped errors.
// It traverses the error chain and collects structured data from each error.
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

// Values extracts all key-value pairs from an error chain and converts them to a map.
// When multiple errors contain the same key, the value from the outermost error is used.
func Values(err error) map[string]any {
	// reversing to get the most upper (re-written) value of the same key
	return lo.FromEntries(lo.Reverse(ValuePairs(err)))
}

// Value retrieves a typed value for a specific key from the error chain.
// It returns the value and a boolean indicating whether the key was found
// and if the value could be converted to the requested type.
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
