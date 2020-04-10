package main

import (
	"errors"
	"fmt"
	"strconv"
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

func (r *Range) GetLowerEndPoint() string {
	return strconv.FormatInt(r.Min, 10)
}

func (r *Range) GetUpperEndPoint() string {
	return strconv.FormatInt(r.Max, 10)
}

func (r *Range) GetRange() string {
	return fmt.Sprintf("[%v,%v]", r.Min, r.Max)
}

func main() {

}
