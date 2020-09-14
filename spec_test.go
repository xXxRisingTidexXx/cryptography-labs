package cryptolabs_test

import (
	"fmt"
)

type spec struct {
	text   string
	key    string
	output string
}

func (spec *spec) String() string {
	return fmt.Sprintf("{%s %s %s}", spec.text, spec.key, spec.output)
}
