/* 	Run this testfile by typing in the terminal;
 	'go test'

	To run all the benchmarks in the test, typ in the terminal:
	'go test -bench=.
*/

package main

import (
	progressbar "progressbar/ProgressBar"
	"testing"
	"time"
)

func TestNorm100(t *testing.T) {
	b := progressbar.Create(100)
	b.Description = "Foor-loop 100"
	for i := 0; i <= b.Total; i++ {
		/*
			The sleep here is to illustrate some work to be done inside the loop
		*/
		time.Sleep(20 * time.Millisecond)
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe100(t *testing.T) {
	b := progressbar.Create(100)
	b.Description = "Pipeline 100"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(20 * time.Millisecond)
		b.UpdatePipeline(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestNorm1000(t *testing.T) {
	b := progressbar.Create(1000)
	b.Description = "Foor-loop 1'000"
	b.Length = 50
	b.GraphColor = "#fa7b62"
	b.DescriptionColor = "#801dae"
	b.Current = 999
	b.Reset(1000)
	for i := 0; i <= b.Total; i += 10 {
		time.Sleep(10 * time.Millisecond)
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe1000(t *testing.T) {
	b := progressbar.Create(1000)
	b.Description = "Pipeline 1'000"
	b.GraphColor = "#ffb95d"
	b.Length = 75
	b.Char = "@"
	for i := 0; i <= b.Total; i += 10 {

		time.Sleep(10 * time.Millisecond)
		b.UpdatePipeline(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestNorm10000(t *testing.T) {
	b := progressbar.Create(10000)
	b.Description = "Foor-loop 10'000"
	for i := 0; i <= b.Total; i++ {
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe10000(t *testing.T) {
	b := progressbar.Create(10000)
	b.Description = "Pipeline 10'000"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {
		b.UpdatePipeline(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestNorm100000(t *testing.T) {
	b := progressbar.Create(100000)
	b.Description = "Foor-loop 100'000"
	for i := 0; i <= b.Total; i++ {
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe100000(t *testing.T) {
	b := progressbar.Create(100000)
	b.Description = "Pipeline 100'000"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {
		b.UpdatePipeline(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestReset(t *testing.T) {
	b := progressbar.Create(50)
	b.Length = 75
	b.Char = "$"
	b.Current = 49
	b.GraphColor = "#c3272b"
	b.Graph = "$$$$"

	b.Reset(50)

	expCurr := 0
	expGraph := ""

	if b.GraphColor != "#c3272b" {
		t.Error("Expected char: {}, but was: {}", "#c3272b", b.GraphColor)
	}

	if b.Graph != expGraph {
		t.Error("Expected char: {}, but was: {}", expGraph, b.Graph)
	}

	if b.Current != expCurr {
		t.Error("Expected char: {}, but was: {}", expCurr, b.Current)
	}
}

func BenchmarkUpdate(b *testing.B) {
	p := progressbar.Create(b.N)
	for i := 0; i < b.N; i++ {
		p.Update(i)
	}
}

func BenchmarkUpdatePipeline(b *testing.B) {
	p := progressbar.Create(b.N)
	for i := 0; i < b.N; i++ {
		p.UpdatePipeline(i)
	}
}

func BenchmarkReset(b *testing.B) {
	p := progressbar.Create(50)
	p.Description = "Testing"
	p.Current = 25
	p.Length = 100
	for i := 0; i < b.N; i++ {
		p.Reset(50)
	}
}
