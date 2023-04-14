package orm

import "testing"

func TestMapToString(t *testing.T) {
	m := map[string]string{"key1": "value1", "key2": "value2"}
	got := mapToString(m)
	want := "key1=value1&key2=value2"
	if got != want {
		t.Errorf("mapToString(%v) = %q, want %q", m, got, want)
	}
}
