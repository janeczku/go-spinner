## Simple CLI Spinner for Go

This is a simple spinner / activity indicator for Go command line apps.
Useful when you want your users to know that there is activity and that the program isn't hung.    
#### Special perks
* The whole spinner line is completely cleared from the console once it is stopped.
* Spinner automagically hides itself if you have piped stdout to somewhere else than the console (You don't really want a spinner in your logger, do you?).    
 
[![asciicast](https://asciinema.org/a/20208.png)](https://asciinema.org/a/20208)

### Installation

To install `spinner.go`, simply run:
```
$ go get github.com/janeczku/go-spinner
```

Make sure your `PATH` includes to the `$GOPATH/bin` directory so your commands can be easily used:
```
export PATH=$PATH:$GOPATH/bin
```

### Example usage

``` go
package main

import (
	"time"
	"github.com/janeczku/go-spinner"
	
)

func main() {
	s := spinner.StartNew("This may take some while...")
	time.Sleep(3 * time.Second) // something more productive here
	s.Stop()
}
```

### API

``` go
s := spinner.StartNew(title string)
```
Quickstart. Creates a new spinner with default options and start it. 

``` go
s := spinner.NewSpinner(title string)
```
Creates a new spinner object.

``` go
s.SetSpeed(time.Millisecond)
```
Sets a custom speed for the spinner animation (default 150ms/frame).

``` go
s.SetCharset([]string)
```
If you don't like the spinning stick, give it an Array of strings like `{".", "o", "0", "@", "*"}`.

``` go
s.Start()
```
Start printing out the spinner animation

``` go
s.Stop()
```
Stops the spinner and clears the line.
