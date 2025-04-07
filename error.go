package efmt

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
)

func New(text string, kv ...KeyValue) error {
	return &e{
		err:    errors.New(text),
		values: kv,
	}
}

func Wrap(err error, text string, kv ...KeyValue) error {
	if err == nil {
		return nil
	}

	return errorf("%s: %w", text, err).Add(kv)
}

func KV(key string, value any) KeyValue {
	return KeyValue{Key: key, Value: value}
}

func errorf(format string, a ...any) *e {
	return &e{
		err:    fmt.Errorf(format, a...),
		values: make([]KeyValue, 0, 1),
	}
}

type e struct {
	err    error
	values []KeyValue
}

func (e *e) Error() string {
	return e.err.Error()
}

func (e *e) Unwrap() error {
	return e.err
}

func (e *e) Values() []KeyValue {
	return e.values
}

func (e *e) Add(keyValues []KeyValue) *e {
	e.values = append(e.values, keyValues...)

	return e
}

type KeyValue = lo.Entry[string, any]
