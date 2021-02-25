package main

import (
	progressbar "progressbar/ProgressBar"
	"testing"
	"time"
)

func TestNorm100(t *testing.T) {
	b := progressbar.Create(100)
	b.Description = "Foor-loop 40"
	for i := 0; i <= b.Total; i++ {
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
		b.UpdatePipeline(1)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestNorm500(t *testing.T) {
	b := progressbar.Create(500)
	b.Description = "Foor-loop 500"
	for i := 0; i <= b.Total; i++ {
		time.Sleep(5 * time.Millisecond)
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe500(t *testing.T) {
	b := progressbar.Create(500)
	b.Description = "Pipeline 500"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(5 * time.Millisecond)
		b.UpdatePipeline(1)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestNorm1000(t *testing.T) {
	b := progressbar.Create(1000)
	b.Description = "Foor-loop 1'000"
	for i := 0; i <= b.Total; i++ {
		time.Sleep(1 * time.Millisecond)
		b.Update(i)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}

func TestPipe1000(t *testing.T) {
	b := progressbar.Create(1000)
	b.Description = "Pipeline 1'000"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(1 * time.Millisecond)
		b.UpdatePipeline(1)
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
		b.UpdatePipeline(1)
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
		b.UpdatePipeline(1)
	}

	if b.Current != b.Total {
		t.Error("Last current expected to be {}, but were {}", b.Total, b.Current)
	}
}
