package main

import (
	"bytes"
	"testing"
)

func TestTransform(t *testing.T) {
	tests := map[string]struct {
		vars      map[string]string
		input     string
		output    string
		expectErr bool
	}{
		"simple": {
			vars:      map[string]string{"FOO": "BAR", "BAZ": "QLUX"},
			input:     "FOO={{ .FOO }}, BAZ={{ .BAZ }}",
			output:    "FOO=BAR, BAZ=QLUX",
			expectErr: false,
		},
		"missing-var": {
			vars:      map[string]string{"FOO": "BAR"},
			input:     `FOO={{ .FOO }}, BAZ={{ .BAZ }}`,
			output:    `FOO=BAR, BAZ=`,
			expectErr: false,
		},
		"required-and-present": {
			vars:      map[string]string{"FOO": "BAR", "BAZ": "QLUX"},
			input:     `FOO={{ .FOO }}, BAZ={{ required "BAZ must be set" .BAZ }}`,
			output:    `FOO=BAR, BAZ=QLUX`,
			expectErr: false,
		},
		"required-and-missing": {
			vars:      map[string]string{"FOO": "BAR"},
			input:     `FOO={{ .FOO }}, BAZ={{ required "BAZ must be set" .BAZ }}`,
			output:    `FOO=BAR, BAZ=`,
			expectErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var in, out bytes.Buffer

			in.WriteString(test.input)
			err := transform(&out, &in, test.vars)

			switch {
			case err == nil && !test.expectErr:
				// No error expected, no error returned 👌
			case err == nil && test.expectErr:
				t.Errorf("expected error, got none")
			case err != nil && !test.expectErr:
				t.Errorf("unexpected error: %s", err.Error())
			case err != nil && test.expectErr:
				// Expected error, got one 👌
			}

			actual := out.String()
			if actual != test.output {
				t.Errorf("wrong output: expected %q, got %q", test.output, actual)
			}
		})
	}
}
