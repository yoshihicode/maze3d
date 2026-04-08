package core

import (
	"math"
	"maze3d/internal/maze"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		s := msg.String()
		if s == "ctrl+c" || s == "esc" {
			return m, tea.Quit
		}

		if m.Status == Title {
			if s == "enter" {
				m.Status = Ready
				m.StartTime = time.Now()
			}
			if s == "esc" {
				return m, tea.Quit
			}
			if s == "m" {
				m.ShowMiniMap = !m.ShowMiniMap
			}
			return m, nil
		}

		if m.Status == Ready {
			return m, nil
		}

		if m.Status == Cleared || m.Status == Giveup {
			if s == "enter" {
				m.Status = Title
				*m = *NewModel()
				return m, tick()
			}
			if s == "n" {
				smm := m.ShowMiniMap
				*m = *NewModel()
				m.ShowMiniMap = smm
				m.Status = Ready
				m.StartTime = time.Now()
				return m, tick()
			}
			if s == "t" {
				m.Status = Ready
				m.PX = 1.5
				m.PY = 1.5
				m.DIR = 0
				m.StartTime = time.Now()
				return m, tick()
			}
			return m, nil
		}

		m.Keys[s] = true
	case tickMsg:
		if m.Keys["w"] {
			m.move(0.05)
		}
		if m.Keys["s"] {
			m.move(-0.05)
		}
		if m.Keys["a"] {
			m.DIR -= math.Pi / 40
		}
		if m.Keys["d"] {
			m.DIR += math.Pi / 40
		}
		if m.Keys["g"] && m.Status == Playing {
			m.Status = Giveup
			return m, tick()
		}

		// reset keys
		m.Keys = make(map[string]bool)

		if time.Since(m.StartTime) > 5*time.Second && m.Status == Ready {
			m.Status = Playing
		}

		return m, tick()
	}

	return m, nil
}

func (m *Model) move(step float64) {

	radius := 0.2 // radius of player
	margin := 0.01

	nx := m.PX + math.Cos(m.DIR)*step
	ny := m.PY + math.Sin(m.DIR)*step

	isWall := func(x, y float64) bool {
		ix, iy := int(x), int(y)
		if ix < 0 || ix >= maze.W || iy < 0 || iy >= maze.H {
			return true
		}
		return m.Maze.Grid[ix][iy] == maze.WALL
	}

	if !isWall(nx+radius+margin, m.PY+radius) &&
		!isWall(nx+radius+margin, m.PY-radius) &&
		!isWall(nx-radius-margin, m.PY+radius) &&
		!isWall(nx-radius-margin, m.PY-radius) {
		m.PX = nx
	}
	if !isWall(m.PX+radius, ny+radius+margin) &&
		!isWall(m.PX-radius, ny+radius+margin) &&
		!isWall(m.PX+radius, ny-radius-margin) &&
		!isWall(m.PX-radius, ny-radius-margin) {
		m.PY = ny
	}

	if m.Maze.Grid[int(m.PX)][int(m.PY)] == maze.GOAL {
		m.Status = Cleared
	}
}
