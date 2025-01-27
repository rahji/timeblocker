package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	flag "github.com/spf13/pflag"
)

func main() {
	startTime := flag.String("start", "09:00", "Start time of work day (HH:MM)")
	endTime := flag.String("end", "17:00", "End time of work day (HH:MM)")
	duration := flag.Int("duration", 30, "Duration of each meeting slot, in minutes")
	buffer := flag.Int("buffer", 5, "Duration of the break between each meeting, in minutes")
	csv := flag.Bool("csv", false, "Print schedule in CSV format")
	flag.Parse()

	// Parse start time
	start, err := time.Parse("15:04", *startTime)
	if err != nil {
		fmt.Printf("Invalid start time format: %v\n", err)
		os.Exit(1)
	}

	// Parse end time
	end, err := time.Parse("15:04", *endTime)
	if err != nil {
		fmt.Printf("Invalid end time format: %v\n", err)
		os.Exit(1)
	}

	// Validate input
	if end.Before(start) {
		fmt.Println("End time must be after start time")
		os.Exit(1)
	}

	if *duration <= 0 {
		fmt.Println("Duration must be positive")
		os.Exit(1)
	}

	var table strings.Builder

	current := start
	for !current.After(end) {
		slotEnd := current.Add(time.Duration(*duration) * time.Minute)

		// Don't include slots that extend beyond the end time
		if slotEnd.After(end) {
			break
		}

		format := "| %-10s | %-10s |\n"
		if *csv {
			format = "%s,%s\n"
		}

		table.WriteString(fmt.Sprintf(format,
			current.Format("15:04"),
			slotEnd.Format("15:04")),
		)

		current = slotEnd.Add(time.Duration(*buffer) * time.Minute)
	}

	if !*csv {
		fmt.Print("\n# Meeting Schedule\n\n")
		fmt.Print("| Start Time | End Time   |\n")
		fmt.Print("|------------|------------|\n")
	}

	fmt.Print(table.String())
}
