package tl

import (
	"errors"
	"fmt"
	"strings"
)

// Type of a Definition or a Parameter.
type Type struct {
	Namespace  []string `json:"namespace,omitempty"`   // namespace components of the type
	Name       string   `json:"name,omitempty"`        // the name of the type
	Bare       bool     `json:"bare,omitempty"`        // whether this type is bare or boxed
	GenericRef bool     `json:"generic_ref,omitempty"` // whether the type name refers to a generic definition
	GenericArg *Type    `json:"generic_arg,omitempty"` // generic arguments of the type
}

func (p Type) String() string {
	var b strings.Builder
	if p.GenericRef {
		b.WriteRune('!')
	}
	for _, ns := range p.Namespace {
		b.WriteString(ns)
		b.WriteRune('.')
	}
	b.WriteString(p.Name)
	if p.GenericArg != nil {
		b.WriteRune('<')
		b.WriteString(p.GenericArg.String())
		b.WriteRune('>')
	}
	return b.String()
}

func (p *Type) Parse(s string) error {
	if strings.HasPrefix(s, ".") {
		return errors.New("type can't start with dot")
	}
	if strings.HasPrefix(s, "!") {
		p.GenericRef = true
		s = s[1:]
	}

	// Parse `type<generic_arg>`.
	if pos := strings.Index(s, "<"); pos >= 0 {
		if !strings.HasSuffix(s, ">") {
			return errors.New("invalid generic")
		}
		p.GenericArg = &Type{}
		if err := p.GenericArg.Parse(s[pos+1 : len(s)-1]); err != nil {
			return fmt.Errorf("failed to parse generic: %w", err)
		}
		s = s[:pos]
	}

	// Parse `ns1.ns2.name`.
	ns := strings.Split(s, ".")
	if len(ns) == 1 {
		p.Name = ns[0]
	} else {
		p.Name = ns[len(ns)-1]
		p.Namespace = ns[:len(ns)-1]
	}
	if p.Name == "" {
		return errors.New("blank name")
	}
	if !isValidName(p.Name) {
		return fmt.Errorf("invalid name %q", p.Name)
	}
	for _, ns := range p.Namespace {
		if !isValidName(ns) {
			return fmt.Errorf("invalid namespace part %q", ns)
		}
	}

	// Bare types starts from lowercase.
	if len(p.Name) > 0 {
		p.Bare = p.Name[0:1] == strings.ToLower(p.Name[0:1])
	}
	return nil
}

// Compile-time interface implementation assertion.
var (
	_ fmt.Stringer = Type{}
)
