package test

import "testing"

func TestThat(t *testing.T) {
	That(t, 1, ShouldNotEqual, 1)
}
