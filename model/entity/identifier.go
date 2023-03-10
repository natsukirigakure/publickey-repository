package entity

import (
	"crypto/rand"
	"github.com/oklog/ulid/v2"
	"io"
	"time"
)

type Identifier struct {
	identifier string `gorm:"size:26"`
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

var defaultGenerator IdentifierGenerator

type IdentifierGenerator interface {
	Generate() Identifier
}

func (g *ULIDGenerator) Generate() Identifier {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), g.entropy)
	return Identifier{id.String()}
}

func init() {
	defaultGenerator = newULIDGenerator(rand.Reader)
}

func newULIDGenerator(r io.Reader) *ULIDGenerator {
	return &ULIDGenerator{
		entropy: ulid.Monotonic(r, 0),
	}
}

func GenerateIdentifier() Identifier {
	return defaultGenerator.Generate()
}

// NewIdentifier は string 型から Identifier 構造体を生成する。
func NewIdentifier(id string) Identifier {
	return Identifier{id}
}

func (i Identifier) Value() string {
	return i.identifier
}

func (i Identifier) Equal(other Identifier) bool {
	return i.identifier == other.identifier
}
