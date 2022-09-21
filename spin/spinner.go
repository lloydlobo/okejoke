package spin

import (
	"time"

	"github.com/briandowns/spinner"
)

// RenderSpinner prints a loading progress indicator.
//
// spinner is a simple package to add a spinner / progress indicator to any terminal application.
//
// New provides a pointer to an instance of Spinner with the supplied options.
// s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"
// For more details about package spinner:
// https://pkg.go.dev/github.com/briandowns/spinner#section-readme
func RenderSpinner() {
	chars := spinner.CharSets[7] // CharSets contains the available character sets
	delay := 100 * time.Millisecond
	s := spinner.New(chars, delay) // Build the spinner with chars, delay (duration).
	s.Color("red")                 // Set the spinner color to red
	s.Start()                      // Start the spinner.
	time.Sleep(time.Second * 4)    // Run for some time to simulate work.
	s.Stop()                       // Stops the indicator.
}
