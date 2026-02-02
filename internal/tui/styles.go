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
	messageBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")). // Purple-ish
			Padding(0, 1).
			MarginBottom(1).
			Width(80)

	// Content Styles
	idStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("14")). // Bright Cyan
		Bold(true)

	timeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("11")) // Yellow

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
