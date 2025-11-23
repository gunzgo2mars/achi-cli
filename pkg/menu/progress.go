package menu

import (
	"fmt"
	"time"
)

type SpinnerProgress struct {
	done chan bool
}

func InitProgress(message string) *SpinnerProgress {

	s := &SpinnerProgress{
		done: make(chan bool),
	}

	go func() {
		spinerChars := []rune{'|', '/', '-', '\\'}
		i := 0
		for {
			select {
			case <-s.done:
				return
			default:
				fmt.Printf("\r%s %c", message, spinerChars[i])
				i = (i + 1) % len(spinerChars)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	return s
}

func (s *SpinnerProgress) Done() {
	s.done <- true
}
