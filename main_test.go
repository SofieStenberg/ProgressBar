// Run this testfile by typing;
// 'go test'
// in the terminal

package main

import (
	"testing"
	"time"

	progressbar "github.com/SofieStenberg/ProgressBar/ProgressBar"
)

func TestNorm100(t *testing.T) {
	b := progressbar.Create(100)
	b.Description = "Foor-loop 40"
	for i := 0; i <= b.Total; i++ {
		/*
			The sleep here is to illustrate some work to be done inside the loop
		*/
		time.Sleep(5 * time.Millisecond)
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe100(t *testing.T) {
	b := progressbar.Create(100)
	b.Description = "Pipeline 40"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(5 * time.Millisecond)
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
	for i := 0; i <= b.Total; i += 10 {
		time.Sleep(25 * time.Nanosecond)
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

		time.Sleep(25 * time.Nanosecond)
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

/*
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
*/
func TestReset(t *testing.T) {
	b := progressbar.Create(50)
	b.Length = 75
	b.Char = "$"
	b.Current = 49
	b.GraphColor = "#c3272b"

	b.Current = 49

	expLength := 50
	expChar := "█"
	expCurr := 0
	expColor := "#0eb83a"

	b.Reset(50)

	if b.Length != expLength {
		t.Error("Expected length: {}, but was: {}", expLength, b.Length)
	}
	if b.Char != expChar {
		t.Error("Expected char: {}, but was: {}", expChar, b.Char)
	}

	if b.Current != expCurr {
		t.Error("Expected char: {}, but was: {}", expCurr, b.Current)
	}

	if b.GraphColor != expColor {
		t.Error("Expected char: {}, but was: {}", expColor, b.GraphColor)
	}

}
