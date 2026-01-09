package tui

import "github.com/charmbracelet/lipgloss"

var (
	primaryCol   = lipgloss.Color("#FF0055")
	secondaryCol = lipgloss.Color("#00A8FF")
	accentCol    = lipgloss.Color("#FFD700")
	errorCol     = lipgloss.Color("#FF0000")
	successCol   = lipgloss.Color("#39FF14")
	bgCol        = lipgloss.Color("#0F0F17")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Underline(true).
			Foreground(primaryCol).
			Background(bgCol).
			Align(lipgloss.Center)
	SubStyle = lipgloss.NewStyle().
			Foreground(accentCol).
			Background(bgCol).
			Align(lipgloss.Center)
	EnvStyle = lipgloss.NewStyle().
			Foreground(secondaryCol).
			Background(bgCol).
			Align(lipgloss.Center)
	BorderBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryCol).
			Background(bgCol).
			Padding(1, 4).
			Width(52).
			Align(lipgloss.Center)
	LabelStyle = lipgloss.NewStyle().
			Foreground(secondaryCol).
			Bold(true)
	TaskStyle = lipgloss.NewStyle().
			Foreground(accentCol).
			Bold(true)
	SuccessStyle = lipgloss.NewStyle().
			Foreground(successCol).
			Bold(true)
	InfoStyle = lipgloss.NewStyle().
			Foreground(accentCol)
	ErrorStyle = lipgloss.NewStyle().
			Foreground(errorCol).
			Bold(true)
)
