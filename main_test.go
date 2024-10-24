package main

import (
	"errors"
	"math"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		in   string
		want struct {
			out float64
			err error
		}
	}{
		{
			in: "3 + 5",
			want: struct {
				out float64
				err error
			}{
				out: 8,
				err: nil,
			},
		},
		{
			in: "5 / 0",
			want: struct {
				out float64
				err error
			}{
				out: 0,
				err: DivisionByZeroError,
			},
		},
		{
			in: "invalid",
			want: struct {
				out float64
				err error
			}{
				out: 0,
				err: InvalidExpressionError,
			},
		},
		{
			in: "1 + 2 * (3 + 4 / 2 - (1 + 2)) * 2 + 1",
			want: struct {
				out float64
				err error
			}{
				out: 10,
				err: nil,
			},
		},
		{
			in: "8/(4 + 3 *(2 +1))-5",
			want: struct {
				out float64
				err error
			}{
				out: -4.385,
				err: nil,
			},
		},
	}

	for _, test := range tests {
		got, err := Calc(test.in)
		if err != nil && !errors.Is(err, test.want.err) {
			t.Errorf("Calc(%s) returned invalid error: got %v, want %v", test.in, err, test.want.err)
		}
		if err == nil && math.Round(got*1000)/1000 != math.Round(test.want.out*1000)/1000 {
			t.Errorf("Calc(%s) = %v, want %v", test.in, got, test.want.out)
		}
	}
}
