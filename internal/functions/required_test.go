package functions

import (
	"reflect"
	"strings"
	"testing"
)

func TestRequired(t *testing.T) {
	tests := map[string]struct {
		msg       string
		value     interface{}
		expectErr bool
	}{
		// The following values are OK.
		"false":       {"", false, false},
		"true":        {"", false, false},
		"string":      {"", "foo bar baz", false},
		"slice":       {"", []string{"foo", "bar", "baz"}, false},
		"empty-slice": {"", []string{}, false},
		"nil-slice":   {"", []string(nil), false}, // Technically, []string(nil) != nil
		"integer":     {"", 12345, false},
		"zero":        {"", 0, false},
		"map":         {"", map[string]string{"A": "B", "C": "D"}, false},
		"empty-map":   {"", map[string]string{}, false},
		"nil-map":     {"", map[string]string(nil), false}, // Technically, []map[string]string(nil) != nil
		// The following values are NOT OK.
		"empty-string": {"strings cannot be empty", "", true},
		"nil-values":   {"values cannot be nil", nil, true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := required(test.msg, test.value)

			if !reflect.DeepEqual(actual, test.value) {
				t.Errorf("value should not change; expected %v, got %v", test.value, actual)
			}

			switch {
			case err == nil && test.expectErr:
				t.Errorf("expected error, got none")
			case err != nil && !test.expectErr:
				t.Errorf("unexpected error: %s", err.Error())
			case err != nil && test.expectErr && !strings.Contains(err.Error(), test.msg):
				t.Errorf("error should contain message")
			}
		})
	}
}
