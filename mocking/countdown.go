package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Give os.Stdout ie. terminal as writer argument to Countdown where to print
// Give ConfigurableSleeper time.Second and time.Sleep
func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}


// Let's create a real sleeper which implements the interface we need
type DefaultSleeper struct {
}

func (d DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Let's define our dependency as an interface. This lets us then use a real Sleeper in main and a spy sleeper in our tests. By using an interface our Countdown function is oblivious to this and adds some flexibility for the caller.
type Sleeper interface {
	Sleep()
}

// Spies are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc. In our case, we're keeping track of how many times Sleep() is called so we can check it in our test.
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}


// We are using duration to configure the time slept and sleep as a way to pass in a sleep function. The signature of sleep is the same as for time.Sleep allowing us to use time.Sleep in our real implementation and the following spy in our tests
type ConfigurableSleeper struct {
	duration time.Duration
	sleep			func(time.Duration)
}

// Sleeps for duration given when initializing ConfigurableSleeper{5 * time.Second, time.Sleep}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Our CountdownOperationsSpy implements both io.Writer and Sleeper, recording every call into one slice. In this test we're only concerned about the order of operations, so just recording them as list of named operations is sufficient.
type CountdownOperationsSpy struct {
	Calls []string
}

// Append string "sleep" into Calls slice
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Append string "write" into Calls slice
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

const finalWord = "Go!"
const countdownStart = 3

// In test we will send to bytes.Buffer so our tests can capture what data is being generated.
// We use fmt.Fprint(out, "3") to print to bytes.Buffer instead of os.stdout
// which would be terminal. But we cannot test or compare strings printed to
// terminal. Both os.Stdout and bytes.Buffer implement io.Writer so by testing
// using bytes.Buffer we get accurate results same as we would be testing os.Stdout
//We're using fmt.Fprint which takes an io.Writer (like *bytes.Buffer) and sends a string to it.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
