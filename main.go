package main

import (
	"fmt"
	"os"

	"maze3d/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	model := core.NewModel()

	p := tea.NewProgram(
		model,
		//tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
