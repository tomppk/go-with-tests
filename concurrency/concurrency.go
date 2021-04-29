package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

//Instead of waiting for a website to respond before sending a request to the next website, we will tell our computer to make the next request while it is waiting.
// Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.

// It returns a map of each URL checked to a boolean value - true for a good response, false for a bad response.
// You also have to pass in a WebsiteChecker which takes a single URL and returns a boolean. This is used by the function to check all the websites.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// Because the only way to start a goroutine is to put go in front of a function call, we often use anonymous functions when we want to start a goroutine. An anonymous function literal looks just the same as a normal function declaration, but without a name (unsurprisingly).

	// 		Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing. Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.
	// The body of the anonymous function above is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function) each of which will add its result to the results map.

// Alongside the results map we now have a resultChannel, which we make in the same way. chan result is the type of the channel - a channel of result. The new type, result has been made to associate the return value of the WebsiteChecker with the url being checked - it's a struct of string and bool. As we don't need either value to be named, each of them is anonymous within the struct; this can be useful in when it's hard to know what to name a value.
// Now when we iterate over the urls, instead of writing to the map directly we're sending a result struct for each call to wc to the resultChannel with a send statement. This uses the <- operator, taking a channel on the left and a value on the right

// The next for loop iterates once for each of the urls. Inside we're using a receive expression, which assigns a value received from a channel to a variable. This also uses the <- operator, but with the two operands now reversed: the channel is now on the right and the variable that we're assigning to is on the left:

// By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time. Although each of the calls of wc, and each send to the result channel, is happening in parallel inside its own process, each of the results is being dealt with one at a time as we take values out of the result channel with the receive expression.
// We have parallelized the part of the code that we wanted to make faster, while making sure that the part that cannot happen in parallel still happens linearly. And we have communicated across the multiple processes involved by using channels.

	for _, url := range urls {
		go func(u string) {
			// Send statement: channel <- value
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression: variable we are assigning value <- channel returning value
		r := <- resultChannel
		// We then use the result received to update the map.
		results[r.string] = r.bool
	}

	return results
}
