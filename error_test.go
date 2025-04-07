package efmt_test

import (
	"errors"
	"testing"

	"github.com/vasilesk/efmt"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		kv      []efmt.KeyValue
		wantErr string
	}{
		{
			name:    "simple error",
			text:    "simple error",
			kv:      nil,
			wantErr: "simple error",
		},
		{
			name:    "error with key-value",
			text:    "error with context",
			kv:      []efmt.KeyValue{efmt.KV("key", "value")},
			wantErr: "error with context",
		},
		{
			name:    "error with multiple key-values",
			text:    "complex error",
			kv:      []efmt.KeyValue{efmt.KV("key1", "value1"), efmt.KV("key2", 42)},
			wantErr: "complex error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := efmt.New(tt.text, tt.kv...)
			if err.Error() != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err.Error(), tt.wantErr)

				return
			}

			// Check values
			if e, ok := err.(interface{ Values() []efmt.KeyValue }); ok {
				if len(e.Values()) != len(tt.kv) {
					t.Errorf("New() values count = %v, want %v", len(e.Values()), len(tt.kv))
				}
			} else {
				t.Errorf("New() error doesn't implement Values() method")
			}
		})
	}
}

func TestWrap(t *testing.T) {
	baseErr := errors.New("base error")

	tests := []struct {
		name    string
		err     error
		text    string
		kv      []efmt.KeyValue
		wantErr string
		wantNil bool
	}{
		{
			name:    "wrap nil error",
			err:     nil,
			text:    "wrapper",
			kv:      nil,
			wantNil: true,
		},
		{
			name:    "wrap simple error",
			err:     baseErr,
			text:    "wrapper",
			kv:      nil,
			wantErr: "wrapper: base error",
		},
		{
			name:    "wrap with key-value",
			err:     baseErr,
			text:    "wrapper",
			kv:      []efmt.KeyValue{efmt.KV("key", "value")},
			wantErr: "wrapper: base error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := efmt.Wrap(tt.err, tt.text, tt.kv...)

			if tt.wantNil {
				if err != nil {
					t.Errorf("Wrap() error = %v, want nil", err)
				}

				return
			}

			if err.Error() != tt.wantErr {
				t.Errorf("Wrap() error = %v, wantErr %v", err.Error(), tt.wantErr)

				return
			}

			// Check unwrapping
			unwrapped := errors.Unwrap(err)
			if unwrapped != tt.err {
				t.Errorf("Unwrap() = %v, want %v", unwrapped, tt.err)
			}

			// Check values
			if e, ok := err.(interface{ Values() []efmt.KeyValue }); ok {
				if len(e.Values()) != len(tt.kv) {
					t.Errorf("Wrap() values count = %v, want %v", len(e.Values()), len(tt.kv))
				}
			} else {
				t.Errorf("Wrap() error doesn't implement Values() method")
			}
		})
	}
}

func TestKV(t *testing.T) {
	tests := []struct {
		name  string
		key   string
		value any
	}{
		{
			name:  "string value",
			key:   "key",
			value: "value",
		},
		{
			name:  "int value",
			key:   "count",
			value: 42,
		},
		{
			name:  "bool value",
			key:   "valid",
			value: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kv := efmt.KV(tt.key, tt.value)
			if kv.Key != tt.key {
				t.Errorf("KV().Key = %v, want %v", kv.Key, tt.key)
			}

			if kv.Value != tt.value {
				t.Errorf("KV().Value = %v, want %v", kv.Value, tt.value)
			}
		})
	}
}
