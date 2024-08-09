package ui

import tea "github.com/charmbracelet/bubbletea"

func Run() {
	prog := tea.NewProgram(newModel())
	prog.Run()
}
