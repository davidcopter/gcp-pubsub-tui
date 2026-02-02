package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"cloud.google.com/go/pubsub"
	"gcp-pubsub-tui/pkg/utils"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PubSubMsg wraps the cloud.google.com/go/pubsub Message for Bubble Tea
type PubSubMsg *pubsub.Message

type Model struct {
	subName  string
	messages []*pubsub.Message
	msgChan  <-chan *pubsub.Message
	err      error
	width    int
	height   int
}

func NewModel(subName string, msgChan <-chan *pubsub.Message) Model {
	return Model{
		subName:  subName,
		msgChan:  msgChan,
		messages: []*pubsub.Message{},
	}
}

func (m Model) Init() tea.Cmd {
	return waitForMessage(m.msgChan)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "c", "C":
			m.messages = []*pubsub.Message{}
			return m, nil
		}
	case PubSubMsg:
		m.messages = append(m.messages, msg)
		// Keep last 50 messages to prevent memory issues and huge renders
		if len(m.messages) > 50 {
			m.messages = m.messages[len(m.messages)-50:]
		}
		return m, waitForMessage(m.msgChan)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Adjust styles based on width if needed
		messageBoxStyle.Width(msg.Width - 4) // simple adjustment
	}
	return m, nil
}

// waitForMessage creates a command that waits for the next message from the channel
func waitForMessage(ch <-chan *pubsub.Message) tea.Cmd {
	return func() tea.Msg {
		msg, ok := <-ch
		if !ok {
			return nil // Channel closed
		}
		return PubSubMsg(msg)
	}
}

func (m Model) View() string {
	// Header
	header := headerStyle.Render(fmt.Sprintf("Pub/Sub Subscription: %s", m.subName))

	// Messages
	var content strings.Builder

	// Render messages newest first
	// We only render as many as reasonable.
	// In a real TUI we'd use a Viewport, but for this simpler version we stack them.

	count := 0
	for i := len(m.messages) - 1; i >= 0; i-- {
		msg := m.messages[i]

		// ID & Time
		id := idStyle.Render(fmt.Sprintf("ID: %s", msg.ID))
		ts := timeStyle.Render(fmt.Sprintf("Time: %s", msg.PublishTime.Format(time.RFC3339)))

		// Attributes
		var attrs string
		if len(msg.Attributes) > 0 {
			// Sort keys for consistent display
			keys := make([]string, 0, len(msg.Attributes))
			for k := range msg.Attributes {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			var attrBuilder strings.Builder
			for _, k := range keys {
				attrBuilder.WriteString(fmt.Sprintf("%s: %s  ", attrKeyStyle.Render(k), attrValueStyle.Render(msg.Attributes[k])))
			}
			attrs = attrBuilder.String()
		}

		// Data
		dataStr := utils.PrettyPrintJSON(msg.Data)
		dataBox := dataBoxStyle.Render(dataContentStyle.Render(dataStr))

		// Combine elements
		// Header line: ID | Time
		headerLine := lipgloss.JoinHorizontal(lipgloss.Left, id, "  ", ts)

		// Body
		var body string
		if attrs != "" {
			body = lipgloss.JoinVertical(lipgloss.Left, headerLine, attrs, dataBox)
		} else {
			body = lipgloss.JoinVertical(lipgloss.Left, headerLine, dataBox)
		}

		content.WriteString(messageBoxStyle.Render(body))
		content.WriteString("\n")

		count++
		// Heuristic limit to avoid rendering too much off-screen text
		if count > 10 {
			break
		}
	}

	if len(m.messages) == 0 {
		content.WriteString(subTitleStyle.Render("Waiting for messages..."))
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"\n",
		content.String(),
		helpStyle.Render("Press 'q' to quit • 'c' to clear messages"),
	)
}
