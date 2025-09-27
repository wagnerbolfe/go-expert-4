package entity

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ID = ulid.ULID

func NewID() ID {
	return ID(ulid.MustNew(ulid.Timestamp(time.Now()), rand.Reader))
}

func ParseID(s string) (ID, error) {
	id, err := ulid.Parse(s)
	return ID(id), err
}
