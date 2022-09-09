package test

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

const (
	success                = ""
	needExactValues        = "This assertion requires exactly %d comparison values (you provided %d)."
	shouldHaveResembled    = "Expected: '%s'\nActual:   '%s'\n(Should resemble)!"
	shouldNotHaveResembled = "Expected        '%#v'\nto NOT resemble '%#v'\n(but it did)!"
)

func need(needed int, expected []interface{}) string {
	if len(expected) != needed {
		return fmt.Sprintf(needExactValues, needed, len(expected))
	}
	return success
}

// ShouldResembleProto receives exactly two parameters and does a proto equal check.
func ShouldResembleProto(actual interface{}, expected ...interface{}) string {
	if message := need(1, expected); message != success {
		return message
	}

	if cmp.Equal(actual, expected[0], protocmp.Transform()) {
		return ""
	}

	return fmt.Sprintf(shouldHaveResembled, actual, expected[0]) +
		cmp.Diff(actual, expected[0], protocmp.Transform())
}

// ShouldNotResembleProto receives exactly two parameters and does an inverse proto equal check.
func ShouldNotResembleProto(actual interface{}, expected ...interface{}) string {
	if message := need(1, expected); message != success {
		return message
	} else if ShouldResembleProto(actual, expected[0]) == success {
		return fmt.Sprintf(shouldNotHaveResembled, actual, expected[0])
	}
	return success
}
