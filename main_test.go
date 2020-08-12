package main

import (
	"testing"
)

func TestGetCountOfSignChanges(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"-", 0},
		{"----", 0},
		{"+", 0},
		{"++++", 0},
		{"-+", 1},
		{"+-", 1},
		{"+-+", 2},
		{"-+-", 2},
		{"--+-", 2},
		{"-----++--", 2},
		{"-+-+-+-+-+", 9},
		{"+-+-+-+-+-", 9},
		{"+--+++-+--++", 6},
		{"+--+++-+--+-", 7},
		{"+abc", 0},
		{"", 0},
		{"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", 0},
	}
	for _, c := range cases {
		got := getCountOfSignChanges(c.in)
		if got != c.want {
			t.Errorf("GetCountOfSignChanges(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCalculateMinimumFlipsNeeded(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"-", 1},
		{"----", 1},
		{"+", 0},
		{"++++", 0},
		{"-+", 1},
		{"+-", 2},
		{"+-+", 2},
		{"-+-", 3},
		{"--+-", 3},
		{"-----++--", 3},
		{"-+-+-+-+-+", 9},
		{"+-+-+-+-+-", 10},
		{"+--+++-+--++", 6},
		{"+--+++-+--+-", 8},
		{"+abc", 0},
		{"", 0},
		{"+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++", 0},
	}
	for _, c := range cases {
		got := calculateMinimumFlipsNeeded(c.in)
		if got != c.want {
			t.Errorf("CalculateMinimumFlipsNeeded(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}
