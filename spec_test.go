package cryptolabs_test

import (
	"fmt"
)

type spec struct {
	text   string
	key    string
	output string
}

func (s spec) String() string {
	return fmt.Sprintf("{%s %s %s}", s.text, s.key, s.output)
}
