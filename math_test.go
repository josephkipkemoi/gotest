package main

import "testing"
import m "go-takabin/takabin/math"

type testpair struct {
	values []float64
	average float64
}

type testmin struct {
	values[] float64
	min float64
}

type testmax struct {
	values[] float64
	max float64
}

type testfib struct {
	values int
	fib int
}

var tests = []testpair {
	{[]float64{0,6}, 3},
	{[]float64{1,1,1,1}, 1},
	{[]float64{-1,1}, 0},
}

var mintests = []testmin {
	{[]float64{1,2}, 1},
	{[]float64{1,2,4,-100, 9}, -100},
	{[]float64{-1,1,9,10}, -1},
}

var maxtests = []testmax {
	{[]float64{0,-10,-9999999}, 0},
	{[]float64{1,2,4,-100, 9}, 9},
	{[]float64{-1,1,9,10}, 10},
}

var fibtests = []testfib {
	{int(10), 55},
	{int(13), 233},
	{int(15), 610},
	{int(0), 0},
	{int(1), 1},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := m.Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got",v,
			)
		}
	}
}

func TestMin(t *testing.T) {
	for _, pair := range mintests {
		v := m.Min(pair.values)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got",v,
			)
		}
	}
}

func TestMax(t *testing.T) {
	for _, pair := range maxtests {
		v := m.Max(pair.values)
		if v != pair.max {
			t.Error(
				"For", pair.values,
				"expected", pair.max,
				"got", v,
			)
		}
	}
}

func TestFibonacci(t *testing.T) {
	for _, pair := range fibtests {
		v := m.Fib(pair.values)
		if v != pair.fib {
			t.Error(
				"For", pair.values,
				"expected", pair.fib,
				"got", v,
			)
		}
	}
}