package tui

import "github.com/charmbracelet/lipgloss"

var (
	// Header Styles
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("255")).
			Background(lipgloss.Color("63")).
			Padding(0, 1)

	subTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginLeft(1)

	// Message Box Styles
	// Replaced with a "Chat Bubble" style
	bubbleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")). // White
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")). // Purple-ish
			Padding(0, 1).
			MarginBottom(1) // Spacing between bubbles

	// Content Styles
	idStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("87")). // Cyan
		Bold(true)

	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("244")). // Grey
			Italic(true)

	attrKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")). // Green
			Bold(true)

	attrValueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("10")) // Green

	dataBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(0, 1).
			MarginTop(0)

	dataContentStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("15")) // White

	// Footer
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginTop(1)
)
