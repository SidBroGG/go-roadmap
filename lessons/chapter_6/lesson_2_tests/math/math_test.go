// _test.go - обязательно для теста

package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("Add(1, 2) = %v, expected %v", result, expected)
	}
}
