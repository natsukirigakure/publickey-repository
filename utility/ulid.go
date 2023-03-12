package utility

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

func NewUlidString() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	return ulid.MustNew(ms, entropy).String()
}
