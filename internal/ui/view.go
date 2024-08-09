package ui

import "github.com/fatih/color"

func (m model) View() string {
	s := color.WhiteString("---------------------------------\n")
	if m.current == 0 {
		s += color.WhiteString("Choose nix type:\n")
		for i := range m.mainSect.items {
			if m.mainSect.cursor == i {
				s += color.BlueString("- %s\n", m.mainSect.items[i])
			} else {
				s += color.WhiteString("- %s\n", m.mainSect.items[i])
			}
		}
	} else if m.current == 1 {
		s += color.WhiteString("What to do with flake?\n")
		for i := range m.flakesSect.items {
			if m.flakesSect.cursor == i {
				s += color.BlueString("- %s\n", m.flakesSect.items[i])
			} else {
				s += color.WhiteString("- %s\n", m.flakesSect.items[i])
			}
		}
	} else if m.current == 2 {
		s += color.WhiteString("What type of flake to init?\n")
		for i := range m.flakeInitSect.items {
			if m.flakeInitSect.cursor == i {
				s += color.BlueString("- %s\n", m.flakeInitSect.items[i])
			} else {
				s += color.WhiteString("- %s\n", m.flakeInitSect.items[i])
			}
		}
	} else if m.current == 3 {
		s += color.WhiteString("What input to update?\n")
		s += m.flakeUpdateInputSect.input.View() + "\n"
	}
	s += color.WhiteString("---------------------------------")
	return s
}
