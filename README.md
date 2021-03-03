# ProgressBar

Install:

go get github.com/SofieStenberg/ProgressBar/ProgressBar

Usage:

import (progressbar "github.com/SofieStenberg/ProgressBar/ProgressBar")

b := progessbar.Create(tot int)
for i :=0; i <= tot; i++ {

    /*
	Do something
    */
	
    b.Update(i) 
}


This Golang package contains the functions to display a progressbar.
Start by using progressbar.CreateProgressBar(maxValue float64) to get a progressbar instance.
The 'maxValue' is how many iterations the function is supposed to go through.
This is a required parameter that must be passed for the bar to be able to calculate the progress.

If you want to use the solution with pipelines, you create the bar as previous, but update with instance.UpdatePipeline.

The bar initializes with default parameters, but some of these can be changed in order to
customize the bar according to own preferences.

With the call		instance.Description	 You can add a string with a description of the bar.

With the call		instace.Length			 You change the length of the displayed bar in the terminal.
            												 Keep in mind that this variable must be a string.

With the call		instance.Graph			 You can change the char that makes the bar progress

You can also custumize the colors on the output. This is done with the HEX value of the color you want to use.
If you want to change color on the description output, simple do: instance.DescriptionColor = "HEX-value".
Remember, the hex-value must be a string and start with #. All the variables that you can change the color for, ends with xxxColor
