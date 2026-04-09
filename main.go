package main

import (
	"flag"
	"fmt"
	"os"

	"maze3d/internal/core"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of maze3d:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  maze3d [options]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  -h, --help\tprint this message\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Controls:\n")
		fmt.Fprintf(flag.CommandLine.Output(), " Title\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [M] MiniMap: ON/OFF\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [Enter] Start\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [Esc] Quit\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), " Playing\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [G] Give up\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [Esc] Quit\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [W] Move Forward\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [S] Move Backward\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [A] Rotate Left\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [D] Rotate Right\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), " Give up\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [N] New Game\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [T] Try Again\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [ENTER] Title\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [Esc] Quit\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), " Cleared\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [N] New Game\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [T] Try Again\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [ENTER] Title\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  [Esc] Quit\n\n")
		flag.PrintDefaults()
	}
	flag.Parse()

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
