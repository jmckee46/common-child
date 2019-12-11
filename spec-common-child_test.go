package main

import "testing"

func TestCommonChild1(t *testing.T) {
	s1 := "AGGTAB"
	s2 := "GXTXAYB"
	lcs := commonChild(s1, s2)

	if lcs != 4 {
		t.Errorf("got %d instead of 124", lcs)
	}
}

func TestCommonChild2(t *testing.T) {
	s1 := "HARRY"
	s2 := "SALLY"
	lcs := commonChild(s1, s2)

	if lcs != 2 {
		t.Errorf("got %d instead of 2", lcs)
	}
}

func TestCommonChild3(t *testing.T) {
	s1 := "AA"
	s2 := "BB"
	lcs := commonChild(s1, s2)

	if lcs != 0 {
		t.Errorf("got %d instead of 0", lcs)
	}
}

func TestCommonChild4(t *testing.T) {
	s1 := "SHINCHAN"
	s2 := "NOHARAAA"
	lcs := commonChild(s1, s2)

	if lcs != 3 {
		t.Errorf("got %d instead of 3", lcs)
	}
}

func TestCommonChild5(t *testing.T) {
	s1 := "ABCDEF"
	s2 := "FBDAMN"
	lcs := commonChild(s1, s2)

	if lcs != 2 {
		t.Errorf("got %d instead of 2", lcs)
	}
}
