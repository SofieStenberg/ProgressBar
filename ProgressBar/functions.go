// Package progressbar contains struct and functions for the progress bar
package progressbar

import (
	"math"
	"strconv"
	"time"

	"github.com/gookit/color"
)

/*
	This package contains the functions to display a progressbar.
	Start by using progressbar.CreateProgressBar(maxValue float64) to get a progressbar instance.
	The 'maxValue' is how many iterations the function is supposed to go through.
	This is a required parameter that must be passed for the bar to be able to calculate the progress.

	The bar initializes with default parameters, but some of these can be changed in order to
	customize the bar according to own preferences.

	With the call		instance.Description	 You can add a string with a description of the bar.

	With the call		instace.Length			 You change the length of the displayed bar in the terminal.
												 Keep in mind that this variable must be a string.

	With the call		instance.Graph			 You can change the char that makes the bar progress

	You can also custumize the colors on the output. This is done with the HEX value of the color you want to use.
	If you want to change color on the description output, simple do: instance.DescriptionColor = "HEX-value".
	Remember, the hex-value must be a string and start with #. All the variables that you can change the color for, ends with xxxColor
*/

// Used instead of divide by 60 as multiplication is faster than division.
const inverseSixty float64 = 0.01666666666666666666666666666667

// ProgressBar is the struct containing the parameters for the progress bar
type ProgressBar struct {
	percent       float64       // The percentage of the progress.
	Current       int           // Keep track of the value of where in the loop the process is.
	Total         int           // Total value. Indicates when the iterations is finished.
	totalInverse  float64       // The inverse of the total value, used for optimization purposes.
	Graph         string        // The cashed string used for displaying the actual bar in the terminal.
	Char          string        // The char used for building the progress bar string.
	Description   string        // Optional description for the progress.
	Length        int           // The length of the bar (How many chars long the bar should be).
	startTime     time.Time     // Marks the time when the progress bar starts.
	elapsedTime   time.Duration // Marks how long time it has been since 'start'.
	estimatedTime time.Duration // Used to get an estimated time of how long time is left of the progress
	isRunning     bool          //	Used to know when to start the timer

	// The below variables is for holding the HEX-value in order to get a colored output
	DescriptionColor string
	CurrentColor     string
	TotalColor       string
	GraphColor       string
	EstimatedColor   string
	ElapsedColor     string
	PercentColor     string
}

// Default values for the variables in the struct ProgressBar.
func Default(b *ProgressBar, n int) {
	b.percent = 0
	b.Current = 0
	b.Total = n
	b.totalInverse = 1.0 / float64(b.Total)
	b.Char = "█"
	b.Description = ""
	b.Graph = ""
	b.Length = 50
	b.isRunning = false
	b.estimatedTime = 0
	b.elapsedTime = 0
	b.startTime.IsZero()

	b.DescriptionColor = "#44cef6"
	b.GraphColor = "#0eb83a"
	b.PercentColor = "#c3272b"
	b.CurrentColor = "#549688"
	b.TotalColor = "#4b5cc4"
	b.EstimatedColor = "#ff7500"
	b.ElapsedColor = "#d9b611"

}

// Create is the constructor for the struct 'ProgressBar'.
// It initializes the bars parameters to its default values.
func Create(n int) *ProgressBar {
	var b *ProgressBar
	b = new(ProgressBar)
	Default(b, n)

	return b
}

// Reset resets the bar to the default values.
func (b *ProgressBar) Reset(n int) {
	Default(b, n)
}

// calculatePercent is used to calculate the percentage of the progress.
func calculatePercent(b *ProgressBar) float64 {
	val := float64(b.Current) * b.totalInverse
	per := math.Min(val*100, 100)
	return per
}

// Estimation is a function that calculates the estimated time left of the progress.
func (b *ProgressBar) estimation() time.Duration {
	elap := time.Since(b.startTime)                                // elapsed time
	iter := float64(b.Total) * (1 - float64(b.percent)/100)        // iterations left
	timePerIns := float64(elap.Nanoseconds()) / float64(b.Current) // timePerInstace
	timeLeft := iter * timePerIns                                  // timeLeft

	return time.Duration(int64(timeLeft))
}

func stopTime(b *ProgressBar) {

	if b.Current == b.Total {
		b.elapsedTime = time.Since(b.startTime)
		color.HEX(b.ElapsedColor, false).Println("\nTime elapsed: ", b.elapsedTime, "\n")
		b.isRunning = false
	}
}

// Update is the function that updates the current state of the progress bar.
func (b *ProgressBar) Update(i int) {
	// Uppdate the 'Current'-parameter with the value of the iteration in the loop
	b.Current = i

	// If at the beginning of the process, the timer starts.
	if !b.isRunning {
		b.startTime = time.Now()
		b.isRunning = true
	}

	uppdateBar(b)
	drawBar(b)

	// If the process is att 100%, a.k.a finished, the timer stops.
	if b.Current == b.Total {
		stopTime(b)
	}
}

// uppdateBar is the function that draws the progressbar to the terminal.
func uppdateBar(b *ProgressBar) {
	// Calculation of the current progress in percent.
	b.percent = calculatePercent(b)
	//e := b.estimation()
	var percent float64

	// As long the 'Percent'-parameter is below 100%, we update the bar.
	if b.percent <= 100 {
		// Calculates how many chars to update the bar with based on the length of the bar.
		percent = (b.percent * 0.01) * float64(b.Length)

		// Calculate how many charachters the parameter RatingBar currently holds.
		currentProgress := []rune(b.Graph)

		// Calculates how many characters to update the bar with,
		// based on how many characters it had the last update.
		progressSinceLast := int(percent) - len(currentProgress)

		// Uppdates the progress in the string-holder.
		for i := 0; i < progressSinceLast; i++ {
			b.Graph += b.Char
		}
	}
}

func drawBar(b *ProgressBar) {
	// Get the estimated time left of the progress.
	e := b.estimation()
	// Converts the variable Legth + 21 to a string so it can be used in the below printf to get the right length of the bar.
	l := strconv.Itoa(b.Length + 21)

	// Sets the color to the right parameters.
	desc := color.HEX(b.DescriptionColor, false).Sprint(b.Description)
	gra := color.HEX(b.GraphColor, false).Sprint(b.Graph)
	per := color.HEX(b.PercentColor, false).Sprint(int(b.percent))
	cur := color.HEX(b.CurrentColor, false).Sprint(b.Current)
	tot := color.HEX(b.TotalColor, false).Sprint(b.Total)

	if e.Seconds() > 60.00 {
		// If over a minute we need to calculate the exact minirues and seconds in order
		// to get the rith output before we sett the color to this parameter.
		min := e.Seconds() * inverseSixty
		sec := int(e.Seconds()) % 60
		estTime := color.HEX(b.EstimatedColor, false).Sprintf("estimated time: %.0fmin %ds ", min, sec)
		color.Printf("\r %s |%-"+l+"s|%s%% %s/%s %s ", desc, gra, per, cur, tot, estTime)
	} else {
		if e.Seconds() >= 0 {
			// If the time left is under a minute, the color is set to the parameter without further ado
			estTime := color.HEX(b.EstimatedColor, false).Sprintf("estimated time: %0.1fs ", e.Seconds())
			color.Printf("\r %s |%-"+l+"s|%s%% %s/%s %s ", desc, gra, per, cur, tot, estTime)
		}
	}
}

/*
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	////////////////////////////Below is a solution for a progressbar that uses pipeline//////////////////////////////////////////

	////////////////////The pipeline-solution is used exactly as the progressbar above by the user///////////////////////////////

	//////////////The only difference is that the user calls 'instance.UpdatePipeline' instead of 'instance.Update'/////////////

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
*/

// UpdatePipeline is the function called when the user wants to use the progressbar
// that is implemented with pipeline. It sends the updated progress through a channel
func (b *ProgressBar) UpdatePipeline(i int) {
	out := make(chan int)
	go func() {
		out <- i
		close(out)
	}()
	b.receive(out)
}

// receive is the function that updates the current state of the progress based on the value sent from the channel.
// It then calls the function to draw the grahp.
func (b *ProgressBar) receive(ch <-chan int) {
	n := <-ch
	b.Current = n

	go func() {
		if !b.isRunning {
			b.startTime = time.Now()
			b.isRunning = true
		}
	}()

	uppdateBar(b)
	drawBar(b)

	// If the process is att 100%, a.k.a finished, the timer stops.
	if b.Current == b.Total {
		stopTime(b)
	}
}
