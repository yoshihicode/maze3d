package core

import (
	"math"
	"maze3d/internal/maze"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type State int

const (
	Title State = iota
	Ready
	Playing
	Cleared
	Giveup
)

type tickMsg time.Time

type Model struct {
	// player state
	PX  float64
	PY  float64
	DIR float64

	// field of view
	FOV float64

	// maze
	Maze maze.Maze

	// key state
	Keys map[string]bool

	// game state
	Status State

	// start time
	StartTime time.Time

	// show minimap
	ShowMiniMap bool
}

func NewModel() *Model {
	m := Model{
		PX:          1.5,
		PY:          1.5,
		DIR:         0,
		FOV:         math.Pi / 3,
		Maze:        *maze.NewMaze(),
		Keys:        make(map[string]bool),
		StartTime:   time.Now(),
		ShowMiniMap: true,
	}

	return &m
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second/30, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
