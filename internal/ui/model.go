package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type mainChoice struct {
	hidden bool
	cursor int
	items  []string
}

type flakesSection struct {
	hidden bool
	cursor int
	items  []string
}

type flakeUpdateInputSection struct {
	hidden bool
	input  textinput.Model
}

type flakeInitSection struct {
	hidden bool
	cursor int
	items  []string
}

func newModel() model {
	ti1 := textinput.New()
	ti1.Placeholder = "what input to update?"
	ti1.Blur()
	return model{
		current: 0,
		mainSect: mainChoice{
			cursor: 0,
			items: []string{
				"Flakes->",
			},
		},
		flakesSect: flakesSection{
			cursor: 0,
			items: []string{
				"Init new flake->",
				"Update input->",
				"Update flake",
			},
		},
		flakeUpdateInputSect: flakeUpdateInputSection{
			input: ti1,
		},
		flakeInitSect: flakeInitSection{
			cursor: 0,
			items: []string{
				"Go",
			},
		},
	}
}

type model struct {
	current              int
	mainSect             mainChoice
	flakesSect           flakesSection
	flakeUpdateInputSect flakeUpdateInputSection
	flakeInitSect        flakeInitSection
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
