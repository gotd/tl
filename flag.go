package tl

import (
	"fmt"
	"strconv"
	"strings"
)

// Flag describes conditional parameter.
type Flag struct {
	// Name of the parameter.
	Name string `json:"name"`
	// Index represent bit index.
	Index int `json:"index"`
}

func (f Flag) String() string {
	return fmt.Sprintf("%s.%d", f.Name, f.Index)
}

func (f *Flag) Parse(s string) error {
	pos := strings.Index(s, ".")
	if pos < 1 {
		return fmt.Errorf("bad flag %q", s)
	}
	f.Name = s[:pos]
	if !isValidName(f.Name) {
		return fmt.Errorf("name %q is invalid", f.Name)
	}
	idx, err := strconv.Atoi(s[pos+1:])
	if err != nil {
		return fmt.Errorf("bad index: %w", err)
	}
	f.Index = idx
	return nil
}
