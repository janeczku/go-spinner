## Simple CLI Spinner for Go

This is a simple spinner / activity indicator for Go command line apps.
Useful when you want your users to know that there is activity and that the program isn't hung.    
#### Special perks
* Spinner is cleared from the console when stopped and the line reset.
* Spinner automagically hides itself if you have piped stdout to somewhere else than the console (You don't really want a spinner in your logger, do you?).    
 
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

**`s := spinner.StartNew(title string)`**

Quickstart. Creates a new spinner with default options and start it. 

**`s := spinner.NewSpinner(title string)`**

Creates a new spinner object.

**`s.SetSpeed(time.Millisecond)`**

Sets a custom speed for the spinner animation (default 150ms/frame).

**`s.SetCharset([]string)`**

If you don't like the spinning stick, give it an Array of strings like `{".", "o", "0", "@", "*"}`.

**`s.Start()`**

Start printing out the spinner animation

**`s.Stop()`**

Stops the spinner and clears the line.
