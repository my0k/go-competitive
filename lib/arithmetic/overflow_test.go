package arithmetic

import (
	"fmt"
	"math"
	"testing"
)

func TestProductOverflow(t *testing.T) {
	testcases := []struct {
		a, b, ub int
		leq      bool
	}{
		{
			a: 100, b: 100, ub: 10000,
			leq: true,
		},
		{
			a: 100, b: 101, ub: 10000,
			leq: false,
		},
		{
			a: 1000000000, b: 1000000000, ub: 1000000000000000000,
			leq: true,
		},
		{
			a: 1000000000, b: 1000000001, ub: 1000000000000000000,
			leq: false,
		},
		{
			a: 3, b: 33, ub: 100,
			leq: true,
		},
		{
			a: 3, b: 34, ub: 100,
			leq: false,
		},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d test case", i)
		t.Run(subTitle, func(t *testing.T) {
			actual := IsProductLeq(tc.a, tc.b, tc.ub)
			if actual != tc.leq {
				t.Errorf("got %v, want %v", actual, tc.leq)
			}
		})
	}
}

func TestSumOverflow(t *testing.T) {
	testcases := []struct {
		a, b, ub int
		leq      bool
	}{
		{
			a: 100, b: 100, ub: 200,
			leq: true,
		},
		{
			a: 100, b: 101, ub: 200,
			leq: false,
		},
		{
			a: 1000000000, b: 1000000000, ub: 2000000000,
			leq: true,
		},
		{
			a: 1000000000, b: 1000000001, ub: 2000000000,
			leq: false,
		},
		{
			a: 3, b: 33, ub: 36,
			leq: true,
		},
		{
			a: 3, b: 34, ub: 36,
			leq: false,
		},
	}

	for i, tc := range testcases {
		subTitle := fmt.Sprintf("%d test case", i)
		t.Run(subTitle, func(t *testing.T) {
			actual := IsSumLeq(tc.a, tc.b, tc.ub)
			if actual != tc.leq {
				t.Errorf("got %v, want %v", actual, tc.leq)
			}
		})
	}
}

func TestProductOverflowInt64(t *testing.T) {
	testcases := []struct {
		a, b int
		ok   bool
	}{
		{
			a: 1, b: math.MaxInt64,
			ok: false,
		},
		{
			a: 2, b: math.MaxInt64,
			ok: true,
		},
		{
			a: 1 << uint(31), b: 1 << uint(31),
			ok: false,
		},
		{
			a: 1 << uint(32), b: 1 << uint(32),
			ok: true,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := IsProductOverflow(tc.a, tc.b)
			if actual != tc.ok {
				t.Errorf("got %v, want %v", actual, tc.ok)
			}
		})
	}
}

func TestSumOverflowInt64(t *testing.T) {
	testcases := []struct {
		a, b int
		ok   bool
	}{
		{
			a: 1, b: math.MaxInt64,
			ok: true,
		},
		{
			a: 1, b: math.MaxInt64 - 1,
			ok: false,
		},
	}

	for i, tc := range testcases {
		testName := fmt.Sprintf("test %d", i)
		t.Run(testName, func(t *testing.T) {
			actual := IsSumOverflow(tc.a, tc.b)
			if actual != tc.ok {
				t.Errorf("got %v, want %v", actual, tc.ok)
			}
		})
	}
}
