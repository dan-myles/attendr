package log

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	Prefix:          "Watcher 👀",
	Level:           log.DebugLevel,
})

func SetLogOptions() {
	// DebugLevel is the style for debug level.
	log.DebugLevelStyle = lipgloss.NewStyle().
		SetString("DEBUG").
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.AdaptiveColor{
			Light: "63",
			Dark:  "63",
		})

	// InfoLevel is the style for info level.
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("INFO").
		Bold(true).
		MaxWidth(4).
		Foreground(lipgloss.AdaptiveColor{
			Light: "39",
			Dark:  "86",
		})

	// WarnLevel is the style for warn level.
	log.WarnLevelStyle = lipgloss.NewStyle().
		SetString("WARN").
		Bold(true).
		MaxWidth(4).
		Foreground(lipgloss.AdaptiveColor{
			Light: "208",
			Dark:  "192",
		})

	// ErrorLevel is the style for error level.
	log.ErrorLevelStyle = lipgloss.NewStyle().
		SetString("ERROR").
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.AdaptiveColor{
			Light: "203",
			Dark:  "204",
		})

	// FatalLevel is the style for error level.
	log.FatalLevelStyle = lipgloss.NewStyle().
		SetString("FATAL").
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.AdaptiveColor{
			Light: "133",
			Dark:  "134",
		})
}
