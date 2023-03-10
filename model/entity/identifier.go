package entity

import "github.com/oklog/ulid/v2"

type Identifier struct {
	identifier string
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func (g *ULIDGenerator) Generate() Identifier {
	return nil
}
