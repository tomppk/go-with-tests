package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

// Proper term for mocks are "test doubles"

// We know we want our Countdown function to write data somewhere and io.Writer is the de-facto way of capturing that as an interface in Go.

// In main we will send to os.Stdout so our users see the countdown printed to the terminal.
// In test we will send to bytes.Buffer so our tests can capture what data is being generated.
func TestCountdown(t *testing.T) {

		t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationsSpy{})
		// The backtick syntax is another way of creating a string but lets you put things like newlines which is perfect for our test.
		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		// Update the tests to inject a dependency on our Spy and assert that the sleep has been called 4 times.
		// if spySleeper.Calls != 4 {
		// 	t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		// }

	})

	// We can now add a sub-test into our test suite which verifies our sleeps and prints operate in the order we hope
	// We now have two tests spying on the Sleeper so we can now refactor our test so one is testing what is being printed and the other one is ensuring we're sleeping in between the prints. Finally we can delete our first spy as it's not used anymore.
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf(("wanted calls %v got %v"), want, spySleepPrinter.Calls)
		}
	})
}

// Use mock type SpyTime to test actual type ConfigurableSleeper
// ConfigurableSleeper gets sleepTime, and spyTime.Sleep() method
// spytime.Sleep() assigns sleepTime 5s to spyTime.durationSlept property
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
