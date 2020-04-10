package main

import (
	"strconv"
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
			testname: "上端よりも下端が大きい閉区間を作成不可能にする",
			min:      8,
			max:      3,
			expect:   NewRangeInvalidRangeElementError,
		},
		{
			testname: "同じ数字を上端と下端に設定し構造体を生成することができる",
			min:      3,
			max:      3,
			expect:   nil,
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

func TestReturnDataPoints(t *testing.T) {
	testrange, err := NewRange(3, 8)
	samerange, _ := NewRange(0, 0)

	assert.NoError(t, err)

	testcase := []struct {
		testname  string
		execfunc  []func() string
		expectval []string
		// expecterr: error
	}{
		{
			testname: "下端を文字列表現で表すことができる",
			execfunc: []func() string{
				func() string {
					return testrange.GetLowerEndPoint()
				},
			},
			expectval: []string{strconv.FormatInt(testrange.Min, 10)},
			// expecterr: nil,
		},
		{
			testname: "上限を文字列表現で返すことができる",
			execfunc: []func() string{
				func() string {
					return testrange.GetUpperEndPoint()
				},
			},
			expectval: []string{strconv.FormatInt(testrange.Max, 10)},
		},
		{
			testname: "下端と上端を文字列表現で返すことができる",
			execfunc: []func() string{
				func() string {
					return testrange.GetRange()
				},
			},
			expectval: []string{"[3,8]"},
		},
		{
			testname: "判定元が [0,0] のときに 上端と下端が 0 になっているのを確認",
			execfunc: []func() string{
				func() string {
					return samerange.GetLowerEndPoint()
				},
				func() string {
					return samerange.GetUpperEndPoint()
				},
			},
			expectval: []string{"0", "0"},
		},
	}

	for _, tc := range testcase {
		t.Run(tc.testname, func(t *testing.T) {
			for i, f := range tc.execfunc {
				result := f()

				if err != nil {
					t.FailNow()
				} else {
					assert.Equal(t, tc.expectval[i], result)
				}
			}
		})
	}
}
