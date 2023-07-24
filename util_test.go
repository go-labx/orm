package orm

import "testing"

func TestMapToString(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		expected string
	}{
		{
			name:     "Test Case 1",
			input:    map[string]string{"key1": "value1", "key2": "value2"},
			expected: "key1=value1&key2=value2",
		},
		{
			name:     "Test Case 2",
			input:    map[string]string{"key": "value"},
			expected: "key=value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MapToString(tc.input)
			if result != tc.expected {
				t.Errorf("expected %s, got %s", tc.expected, result)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"All lower case", "testcase", "testcase"},
		{"Camel case", "testCase", "test_case"},
		{"With numbers", "testCase123", "test_case123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToSnakeCase(tt.input)
			if result != tt.expected {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
