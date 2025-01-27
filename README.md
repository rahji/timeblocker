# TimeBlocker

TimeBlocker is a simple CLI program that generates a schedule of possible meeting times.

## Command-Line Options

```
-h, --help            Show this help
    --start=STRING    Start time of work day (HH:MM)
    --end=STRING      End time of work day (HH:MM)
    --duration=INT    Duration of each meeting slot, in minutes
    --break=5         Duration of the break between each meeting, in minutes
    --csv             Show output as CSV instead of a Markdown table
    --kitchen         Show output times as kitchen time aka 12-hour time
```

## Installation

`go install github.com/rahji/timeblocker@latest`

## Notes

Considering `timeblocker` outputs Markdown by default, you might consider piping its output to [glow](https://github.com/charmbracelet/glow)
