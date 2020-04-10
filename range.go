package main

import (
	"errors"
)

type Range struct {
	Min int64
	Max int64
}

var (
	NewRangeInvalidRangeElementError = errors.New("Invalid range : min must be less than max")
)

func NewRange(min, max int64) (*Range, error) {

	if min > max {
		return &Range{}, NewRangeInvalidRangeElementError
	}

	return &Range{
		Min: min,
		Max: max,
	}, nil
}

func main() {

}
