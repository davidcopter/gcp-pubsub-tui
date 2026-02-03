package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"gcp-pubsub-tui/pkg/utils"

	"cloud.google.com/go/pubsub/v2"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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
	viewport viewport.Model
	// when true, the viewport should auto-scroll to bottom on next render
	shouldAutoScroll bool
}

func NewModel(subName string, msgChan <-chan *pubsub.Message) Model {
	return Model{
		subName:          subName,
		msgChan:          msgChan,
		messages:         []*pubsub.Message{},
		viewport:         viewport.New(0, 0),
		shouldAutoScroll: true,
	}
}

func (m Model) Init() tea.Cmd {
	return waitForMessage(m.msgChan)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()
		switch {
		case s == "q" || s == "ctrl+c" || msg.Type == tea.KeyCtrlC:
			return m, tea.Quit
		case s == "c" || s == "C":
			m.messages = []*pubsub.Message{}
			m.viewport.GotoTop()
			m.shouldAutoScroll = true
			return m, nil
		case s == "up" || msg.Type == tea.KeyUp || s == "k":
			m.viewport.LineUp(1)
			m.shouldAutoScroll = false
			return m, nil
		case s == "down" || msg.Type == tea.KeyDown || s == "j":
			m.viewport.LineDown(1)
			// If moving down, and we reach bottom, re-enable auto-scroll.
			if m.viewport.AtBottom() {
				m.shouldAutoScroll = true
			} else {
				m.shouldAutoScroll = false
			}
			return m, nil
		case msg.Type == tea.KeyPgUp:
			m.viewport.PageUp()
			m.shouldAutoScroll = false
			return m, nil
		case msg.Type == tea.KeyPgDown:
			m.viewport.PageDown()
			if m.viewport.AtBottom() {
				m.shouldAutoScroll = true
			} else {
				m.shouldAutoScroll = false
			}
			return m, nil
		}
	case PubSubMsg:
		m.messages = append(m.messages, msg)
		// Keep last 100 messages
		if len(m.messages) > 100 {
			m.messages = m.messages[len(m.messages)-100:]
		}
		// Auto-scroll to bottom only if viewport was already at bottom.
		// We can't call GotoBottom() here because SetContent (in View)
		// updates the viewport's internal lines. Instead, set a flag
		// and perform the GotoBottom after SetContent during rendering.
		if m.viewport.AtBottom() {
			m.shouldAutoScroll = true
		}
		return m, waitForMessage(m.msgChan)
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height - 4 // Reserve space for header + help
		m.viewport.Width = m.width - 4
		m.viewport.Height = m.height
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

	// Messages - build content OLDEST first (chronological order)
	var content strings.Builder

	for i := 0; i < len(m.messages); i++ {
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
		content.WriteString("\n\n")
	}

	if len(m.messages) == 0 {
		content.WriteString(subTitleStyle.Render("Waiting for messages..."))
	}

	m.viewport.SetContent(content.String())

	// If a recent message requested auto-scroll, perform it now that
	// the viewport content has been updated.
	if m.shouldAutoScroll {
		m.viewport.GotoBottom()
		m.shouldAutoScroll = false
	}

	help := helpStyle.Render("Press 'q' to quit • 'c' to clear • ↑/↓ to scroll")

	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"\n",
		m.viewport.View(),
		help,
	)
}
