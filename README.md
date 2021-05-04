# test

A golang test library. It's essentially a light wrapper around all assertions from https://github.com/smartystreets/assertions.

## Usage

```golang
package test

import (
	"testing"

	"github.com/edaniels/test"
)

func TestExample(t *testing.T) {
	test.That(t, 1, ShouldEqual, 1)
	test.That(t, 1, ShouldNotEqual, "1")
}
```
