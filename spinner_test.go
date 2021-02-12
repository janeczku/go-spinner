package spinner_test

import (
	"bytes"
	spinner "github.com/elisarver/go-spinner"
	"testing"
	"time"
)

func TestSynch(t *testing.T) {
			var out bytes.Buffer
			sp := spinner.NewSpinner("Message1")
			sp.Output = &out

			sp2 := spinner.NewSpinner("Message2")
			sp2.Output = &out

			sp.Start()
			time.Sleep(2 * time.Second)
			sp.Stop()

			sp2.Start()
			time.Sleep(2 * time.Second)
			sp2.Stop()

			got := out.String()
			want := "| Message1/ Message1- Message1\\ Message1| Message1/ Message1- Message1\\ Message1| Message1/ Message1- Message1\\ Message1| Message1/ Message1| Message2/ Message2- Message2\\ Message2| Message2/ Message2- Message2\\ Message2| Message2/ Message2- Message2\\ Message2| Message2/ Message2"
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
}


