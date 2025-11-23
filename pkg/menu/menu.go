package menu

import (
	"fmt"
	"log"
	"syscall"

	"github.com/gunzgo2mars/achi-cli/internal/core/model"
	"golang.org/x/term"
)

var options = []model.MenuList{
	{
		MID:      0,
		MenuName: "Create http service layout",
	},
	{
		MID:      1,
		MenuName: "Create cli layout",
	},
}
var selectedIndex = 0

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func draw() {
	clearScreen()

	log.Printf("===========================================\r\n")
	log.Printf("⚔ Achilles - go project layout generator ⚔\r\n")
	log.Printf("===========================================\r\n")
	for index, value := range options {
		if index == selectedIndex {
			// ANSI escape codes for reverse video (highlighting)
			log.Printf("\033[7m  ▸ %s \033[0m\r\n", value.MenuName)
		} else {
			log.Printf("  %s \r\n", value.MenuName)
		}
	}
}

func readKey(fd int) (byte, error) {
	var buf [1]byte
	n, err := syscall.Read(fd, buf[:])
	if err != nil {
		return 0, err
	}
	if n == 0 {
		return 0, nil
	}
	return buf[0], nil
}

func HandleSelection(oldState *term.State, fd int) (*model.MenuList, error) {
	for {
		draw()

		key, err := readKey(fd)
		if err != nil {
			return nil, err
		}

		// Handle ANSI escape sequence for arrow keys (usually starts with 0x1b, 0x5b)
		if key == 0x1b {
			key, _ = readKey(fd) // Read [
			if key == 0x5b {
				key, _ = readKey(fd)
				switch key {
				case 0x41: // Up Arrow
					if selectedIndex > 0 {
						selectedIndex--
					}
				case 0x42: // Down Arrow
					if selectedIndex < len(options)-1 {
						selectedIndex++
					}
				}
			}
		} else if key == 0x0a || key == 0x0d { // Enter key (Line Feed or Carriage Return)
			// Restore terminal state before exiting
			term.Restore(fd, oldState)
			clearScreen()
			return &options[selectedIndex], nil
		}
	}
}
