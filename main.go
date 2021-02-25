package main

import (
	progressbar "progressbar/ProgressBar"
	"time"
)

func main() {

	// Declare and initialize the 'bar' variable.
	b := progressbar.Create(100)
	// Choose to set a description to the bar. Not a necessity
	b.Description = "Foor-loop 100"
	//b.Length = 25
	for i := 0; i <= b.Total; i++ {
		time.Sleep(5 * time.Nanosecond)

		// At the end of each iteration, we update the progress bar.
		b.Update(i)
	}

	b.Reset(100)
	// Choose to set a description to the bar. Not a necessity
	b.Description = "Pipeline 100"
	b.Char = "@"
	b.Length = 75
	b.DescriptionColor = "#845a33"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(5 * time.Nanosecond)
		b.UpdatePipeline(1)
	}

	b.Reset(500)
	b.Description = "Foor-loop 500"
	for i := 0; i <= b.Total; i++ {
		time.Sleep(1 * time.Nanosecond)

		b.Update(i)
	}

	b.Reset(500)
	b.Description = "Pipeline 500"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {

		time.Sleep(1 * time.Nanosecond)
		b.UpdatePipeline(1)
	}

	b.Reset(1000)
	b.Description = "Foor-loop 1'000"
	for i := 0; i <= b.Total; i++ {
		time.Sleep(1 * time.Nanosecond)
		b.Update(i)
	}

	b.Reset(1000)
	b.Description = "Pipeline 1'000"
	b.Char = "@"
	for i := 0; i <= b.Total; i++ {
		time.Sleep(1 * time.Nanosecond)
		b.UpdatePipeline(1)
	}
	/*
		b.Reset(10000)
		b.Description = "Foor-loop 10'000"
		for i := 0; i <= b.Total; i++ {
			b.Update(i)
		}

		b.Reset(10000)
		b.Description = "Pipeline 10'000"
		b.Char = "@"
		for i := 0; i <= b.Total; i++ {
			b.UpdatePipeline(1)
		}

		b.Reset(100000)
		b.Description = "Foor-loop 100'000"
		for i := 0; i <= b.Total; i++ {
			b.Update(i)
		}

		b.Reset(100000)
		b.Description = "Pipeline 100'000"
		b.Char = "@"
		for i := 0; i <= b.Total; i++ {
			b.UpdatePipeline(1)
		}
	*/
	////////////////////////////////////////////////////
	/*
		b.Reset(1000000)
		b.Description = "The 1'000'000 for-loop"
		b.Length = 75
		for i := 0; i <= b.Total; i++ {
			b.Update(i)
		}

		b.Reset(1000000)
		b.Description = "Pipeline 1'000'000"
		b.Char = "@"
		for i := 0; i <= b.Total; i++ {
			b.UpdatePipeline(1)
		}
	*/

	//b.Reset(500)

}
