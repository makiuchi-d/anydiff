package anydiff_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/makiuchi-d/anydiff"
)

func TestDiff(t *testing.T) {

	a1 := []byte("abcdefg")
	b1 := []byte("abXceZg")
	e1 := "==+=-=-+="
	testDiff(t, a1, b1, anydiff.Cmp, e1)

	a2 := []string{
		"line 1",
		"line 2",
		"line 3",
		"line 4",
	}
	b2 := []string{
		"line 1",
		"line 2 mod",
		"line 3",
		"line 5",
	}
	e2 := "=-+=-+"
	testDiff(t, a2, b2, anydiff.Cmp, e2)

	a3 := []int{1, 2, 3, 4, 5}
	b3 := []float32{2, 3.5, 4, 6}
	c3 := func(a *int, b *float32) bool { return float32(*a) == *b }
	e3 := "-=-+=-+"
	testDiff(t, a3, b3, c3, e3)

	type Line struct {
		Num  int
		Text string
	}
	a4 := []Line{
		{1, "foo"},
		{2, "bar"},
		{3, "baz"},
		{4, "qux"},
	}
	b4 := []Line{
		{1, "foo"},
		{2, "BAR"}, // case insensitive: no changed
		{5, "qux"},
	}
	c4 := func(a, b *Line) bool {
		return a.Num == b.Num && strings.ToLower(a.Text) == strings.ToLower(b.Text)
	}
	e4 := "==--+"
	testDiff(t, a4, b4, c4, e4)
}

func testDiff[A, B any](t *testing.T, a []A, b []B, cmp func(*A, *B) bool, exp string) {
	ret := anydiff.Diff(a, b, cmp)
	if !reflect.DeepEqual([]byte(ret), []byte(exp)) {
		t.Fatalf("Diff:\na: %v\nb: %v\nedit:  %q\nwants: %q", a, b, ret, exp)
	}
}

func TestEditDistance(t *testing.T) {
	tests := []struct {
		seq string
		exp int
	}{
		{"==========", 0},
		{"=-+===--=+", 5},
		{"-----+++++", 10},
	}

	for _, test := range tests {
		e := anydiff.Edit(test.seq)
		if d := e.Distance(); d != test.exp {
			t.Errorf("distance of %q = %v wants %v", test.seq, d, test.exp)
		}
	}
}
