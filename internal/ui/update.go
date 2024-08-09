package ui

import (
	"fmt"
	"os"

	"github.com/ProggerX/lazynix/pkg/nix/flakes"
	"github.com/fatih/color"

	tea "github.com/charmbracelet/bubbletea"
)

func updateFlake() {
	fmt.Println()
	color.White("Updating flake...")
	out, err := flakes.Update()
	if err != nil {
		color.Red("Error! Output:")
		color.Red(string(out))
		os.Exit(1)
	}
	color.White("Output:")
	color.White(string(out))
	os.Exit(0)
}

func updateInput(name string) {
	fmt.Println()
	color.White("Updating flake...")
	out, err := flakes.UpdateInput(name)
	if err != nil {
		color.Red("Error! Output:")
		color.Red(string(out))
		os.Exit(1)
	}
	color.White("Output:")
	color.White(string(out))
	os.Exit(0)
}

func initFlake(templ string) {
	fmt.Println()
	color.White("Creating flake with template %s\n", templ)
	err := flakes.Init(templ)
	if err != nil {
		color.Red("Error!")
		os.Exit(1)
	}
	color.White("Success!")
	os.Exit(0)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.current == 3 {
		m.flakeUpdateInputSect.input.Focus()
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				updateInput(m.flakeUpdateInputSect.input.Value())
				return m, tea.Quit
			case "ctrl+c":
				return m, tea.Quit
			}
		}
		var cmd tea.Cmd
		m.flakeUpdateInputSect.input, cmd = m.flakeUpdateInputSect.input.Update(msg)
		return m, cmd
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "j", "k":
			switch m.current {
			case 0:
				if msg.String() == "j" {
					if m.mainSect.cursor < len(m.mainSect.items)-1 {
						m.mainSect.cursor++
					}
				} else {
					if m.mainSect.cursor > 0 {
						m.mainSect.cursor--
					}
				}
			case 1:
				if msg.String() == "j" {
					if m.flakesSect.cursor < len(m.flakesSect.items)-1 {
						m.flakesSect.cursor++
					}
				} else {
					if m.flakesSect.cursor > 0 {
						m.flakesSect.cursor--
					}
				}
			case 2:
				if msg.String() == "j" {
					if m.flakeInitSect.cursor < len(m.flakeInitSect.items)-1 {
						m.flakeInitSect.cursor++
					}
				} else {
					if m.flakeInitSect.cursor > 0 {
						m.flakeInitSect.cursor--
					}
				}
			}
		case "l":
			switch m.current {
			case 0:
				switch m.mainSect.cursor {
				case 0:
					m.current = 1
				}
			case 1:
				switch m.flakesSect.cursor {
				case 0:
					m.current = 2
				case 1:
					m.current = 3
					return m, nil
				case 2:
					updateFlake()
				}
			case 2:
				initFlake(m.flakeInitSect.items[m.flakeInitSect.cursor])
			}
		case "h":
			switch m.current {
			case 1:
				m.current = 0
			case 2, 3:
				m.current = 1
			}
		}
	}
	return m, nil
}
