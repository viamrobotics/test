package test

import "testing"

func TestThat(t *testing.T) {
	That(t, 1, ShouldEqual, 1)
}

func TestNonFatalThat(t *testing.T) {
	NonFatalThat(t, 1, ShouldEqual, 1)
}
