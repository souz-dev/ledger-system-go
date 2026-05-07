package domain

import "errors"

type Direction string

const (
	Debit  Direction = "debit"
	Credit Direction = "credit"
)

var ErrInvalidDirection = errors.New("direction must be debit or credit")

func (d Direction) IsValid() bool {
	return d == Debit || d == Credit
}
