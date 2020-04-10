package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRangeStruct(t *testing.T) {
	testcase := []struct {
		testname string
		min      int64
		max      int64
		expect   error
	}{
		{
			testname: "下端と上端を持つ構造体を生成するテスト",
			min:      3,
			max:      8,
			expect:   nil,
		},
		{
			testname: "下端が設定されていない時には構造体を生成できないようにする",
			min:      8,
			max:      3,
			expect:   NewRangeInvalidRangeElementError,
		},
	}

	for _, tc := range testcase {
		t.Run(tc.testname, func(t *testing.T) {
			closedRange, err := NewRange(tc.min, tc.max)

			if tc.expect != nil {
				assert.EqualError(t, tc.expect, err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.min, closedRange.Min)
				assert.Equal(t, tc.max, closedRange.Max)
			}
		})
	}
}
