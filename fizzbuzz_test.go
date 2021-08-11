package fizzbuzzgo

import (
	"fmt"
	"testing"
)

var fbtests = []struct {
	in  int
	out string
}{
	{1, "1"},
	{2, "2"},
	{3, "fizz"},
	{4, "4"},
	{5, "buzz"},
	{6, "fizz"},
	{15, "fizzbuzz"},
	{17, "17"},
	{60, "fizzbuzz"},
	{97, "97"},
	{99, "fizz"},
	{100, "buzz"},
}

func TestFizzBuzz(t *testing.T) {
	for _, tt := range fbtests {
		t.Run(fmt.Sprintf("FizzBuzz(%v)", tt.in), func(t *testing.T) {
			got := FizzBuzz(tt.in)
			if got != tt.out {
				t.Errorf("got %q, want %q", got, tt.out)
			}
		})
	}
}
