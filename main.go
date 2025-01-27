package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
)

type Flags struct {
	Start    string `kong:"required,help='Start time of work day (HH:MM)'"`
	End      string `kong:"required,help='End time of work day (HH:MM)'"`
	Duration int    `kong:"required,help='Duration of each meeting slot, in minutes'"`
	Break    int    `kong:"default=5,help='Duration of the break between each meeting, in minutes'"`
	Csv      bool   `kong:"default=false,help='Show output as CSV instead of a Markdown table'"`
	Kitchen  bool   `kong:"default=true,help='Show output times as kitchen time aka 12-hour time'"`
}

func main() {
	var flags Flags
	ctx := kong.Parse(&flags)
	_ = ctx

	// Parse day start time
	daystart, err := time.Parse("15:04", flags.Start)
	if err != nil {
		fmt.Printf("Invalid start time format: %v\n", err)
		os.Exit(1)
	}

	// Parse day end time
	dayend, err := time.Parse("15:04", flags.End)
	if err != nil {
		fmt.Printf("Invalid end time format: %v\n", err)
		os.Exit(1)
	}

	// Validate input
	if dayend.Before(daystart) {
		fmt.Println("End time must be after start time")
		os.Exit(1)
	}

	if flags.Break <= 0 {
		fmt.Println("Duration must be positive")
		os.Exit(1)
	}

	var table strings.Builder

	currentstart := daystart
	for !currentstart.After(dayend) {
		currentend := currentstart.Add(time.Duration(flags.Duration) * time.Minute)

		// Don't include slots that extend beyond the end time
		if currentend.After(dayend) {
			break
		}

		format := "| %-10s | %-10s |\n"
		if flags.Csv {
			format = "%s,%s\n"
		}

		s := currentstart.Format("15:04")
		e := currentend.Format("15:04")
		if flags.Kitchen {
			s = currentstart.Format(time.Kitchen)
			e = currentend.Format(time.Kitchen)
		}

		table.WriteString(fmt.Sprintf(format, s, e))

		currentstart = currentend.Add(time.Duration(flags.Break) * time.Minute)
	}

	if !flags.Csv {
		fmt.Print("\n# Meeting Schedule\n\n")
		fmt.Print("| Start Time | End Time   |\n")
		fmt.Print("|------------|------------|\n")
	}

	fmt.Print(table.String())
}
