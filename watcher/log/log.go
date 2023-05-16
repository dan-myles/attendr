package log

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

var Logger = log.NewWithOptions(os.Stderr, log.Options{
	ReportCaller:    true,
	ReportTimestamp: true,
	Prefix:          "Watcher 👀",
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

	err := godotenv.Load()
	if err != nil {
		Logger.Fatal("Error loading .env file")
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "debug" {
		Logger.SetLevel(log.DebugLevel)
	}

	if logLevel == "warn" {
		Logger.SetLevel(log.WarnLevel)
	}

	if logLevel == "error" {
		Logger.SetLevel(log.ErrorLevel)
	}

	if logLevel == "info" {
		Logger.SetLevel(log.InfoLevel)
	}

	if logLevel == "fatal" {
		Logger.SetLevel(log.FatalLevel)
	}

	if logLevel == "all" {
		Logger.SetLevel(log.DebugLevel)
	}
}
