package core

import (
	"fmt"
	"math"
	"maze3d/internal/maze"
	"maze3d/internal/render"
	"strings"
	"time"
)

const (
	screenW = 120
	screenH = 22
)

type Screen struct {
	Grid [screenW][screenH]int
}

func (m *Model) View() string {
	var b strings.Builder
	switch m.Status {
	case Title:
		mstatus := "OFF"
		if m.ShowMiniMap {
			mstatus = "ON"
		}

		b.WriteString("\n")
		b.WriteString(" ███    ███  █████  ███████ ███████  ██████  ██████  \n")
		b.WriteString(" ████  ████ ██   ██    ███  ██            ██ ██   ██ \n")
		b.WriteString(" ██ ████ ██ ███████   ███   █████      ████  ██   ██ \n")
		b.WriteString(" ██  ██  ██ ██   ██  ███    ██            ██ ██   ██ \n")
		b.WriteString(" ██      ██ ██   ██ ███████ ███████  ██████  ██████  \n")
		b.WriteString("\n")
		b.WriteString("\n")
		b.WriteString(fmt.Sprintf("                   [M] MiniMap: %s\n", mstatus))
		b.WriteString("                   [Enter] Start\n")
		b.WriteString("                   [Esc] Quit\n")
	case Ready:
		remaining := 5 - int(time.Since(m.StartTime).Seconds())
		msg := []string{
			"",
			"  " + fmt.Sprintf("Starting in %d seconds...", remaining),
			"",
			"     [W] UP",
			"     [S] DOWN",
			"     [A] LEFT",
			"     [D] RIGHT",
			"     [G] Give up",
			"     [ESC] Quit",
			"",
			"     " + render.Bg(render.PlayerStyle).Render("P") + " : Player",
			"     " + render.Bg(render.GoalStyle).Render("G") + " : Goal",
		}

		b.WriteString(m.fullMap(msg))
	case Giveup:
		msg := []string{
			"",
			"  ====================",
			"  😭     GIVE UP    😭",
			"  ====================",
			"",
			"      [N] New Game",
			"      [T] Try Again",
			"      [ENTER] Title",
			"      [Esc] Quit",
		}
		b.WriteString(m.fullMap(msg))
	case Cleared:
		msg := []string{
			"",
			"  ====================",
			"  ✨     GOAL!!!    ✨",
			"  ====================",
			"",
			"      [N] New Game",
			"      [T] Try Again",
			"      [ENTER] Title",
			"      [Esc] Quit",
		}
		b.WriteString(m.fullMap(msg))
	case Playing:
		// buffer for drawing
		scr := Screen{}
		for x := range scr.Grid {
			for y := range scr.Grid[x] {
				scr.Grid[x][y] = render.BackgroundStyle
			}
		}

		for sx := 0; sx < screenW; sx++ {
			rayAngle := m.DIR + (float64(sx)/screenW-0.5)*m.FOV

			RX := m.PX
			RY := m.PY

			dist := 0.0

			for dist < 20 {

				RX += math.Cos(rayAngle) * 0.05
				RY += math.Sin(rayAngle) * 0.05

				if RX < 0 || RY < 0 || RX >= maze.W || RY >= maze.H {
					break
				}

				if m.Maze.Grid[int(RX)][int(RY)] == maze.WALL {
					break
				}

				dist += 0.05
			}

			// fisheye correction
			dist = dist * math.Cos(rayAngle-m.DIR)

			// wall
			height := int(15 / (dist + 0.1))
			shade := render.GrayScale(dist)

			start := screenH/2 - height
			end := screenH/2 + height

			for y := start; y <= end; y++ {
				if y >= 0 && y < screenH {
					scr.Grid[sx][y] = shade
				}
			}

			// floor
			for y := screenH/2 + height; y < screenH; y++ {
				// distance from screen center
				dyScreen := float64(y) - float64(screenH)/2.0

				if dyScreen == 0 {
					continue
				}

				// distance to floor
				rowDist := (float64(screenH) / (2.0 * dyScreen)) / math.Cos(rayAngle-m.DIR)

				// correct ray direction
				floorX := m.PX + math.Cos(rayAngle)*rowDist
				floorY := m.PY + math.Sin(rayAngle)*rowDist

				fx := int(floorX)
				fy := int(floorY)

				if fx >= 0 && fy >= 0 && fx < maze.W && fy < maze.H && m.Maze.Grid[fx][fy] == maze.GOAL {
					scr.Grid[sx][y] = render.GoalStyle
				} else {
					scr.Grid[sx][y] = render.FloorStyle
				}
			}
		}

		mSize := 2                       // size of mini map
		offsetX := screenW - mSize*2 - 3 // char position of mini map
		offsetY := mSize + 1             // char position of mini map

		if m.ShowMiniMap {
			m.miniMap(int(m.PX), int(m.PY), mSize, offsetX, offsetY, &scr)
		}

		pchar := func(dir float64) rune {
			for dir < 0 {
				dir += 2 * math.Pi
			}
			for dir >= 2*math.Pi {
				dir -= 2 * math.Pi
			}

			switch {
			case dir < math.Pi/4:
				return '>'
			case dir < 3*math.Pi/4:
				return 'v'
			case dir < 5*math.Pi/4:
				return '<'
			case dir < 7*math.Pi/4:
				return '^'
			default:
				return '>'
			}
		}

		//var b strings.Builder
		for y := range screenH {
			for x := range screenW {

				ch := " "

				if m.ShowMiniMap {
					// player on mini map
					if x == offsetX {
						if y == offsetY {
							ch = string(pchar(m.DIR))
						}
					}
				}

				b.WriteString(
					render.Bg(scr.Grid[x][y]).Render(ch),
				)

			}
			b.WriteRune('\n')
		}
	}
	return b.String()

}

func (m *Model) fullMap(msg []string) string {
	var b strings.Builder

	for y := 0; y < maze.H; y++ {
		for x := 0; x < maze.W; x++ {
			color := render.BackgroundStyle
			ch := " "

			if x == int(m.PX) && y == int(m.PY) {
				color = render.PlayerStyle
				ch = "P"
			} else if m.Maze.Grid[x][y] == maze.WALL {
				color = render.WallStyle
			} else if m.Maze.Grid[x][y] == maze.GOAL {
				color = render.GoalStyle
				ch = "G"
			}
			b.WriteString(render.Bg(color).Render(ch))
		}
		if msg != nil && y < len(msg) {
			b.WriteString(msg[y])
		}
		b.WriteRune('\n')
	}

	return b.String()
}

func (m *Model) miniMap(px, py, size, offsetX, offsetY int, scr *Screen) {

	for y := size * -1; y <= size; y++ {
		for x := size * -2; x <= size*2; x++ {

			wx := px + x
			wy := py + y

			var color int

			if wx >= 0 && wy >= 0 && wx < maze.W && wy < maze.H {

				color = render.MaskStyle

				if m.Maze.Grid[wx][wy] == maze.WALL {
					color = render.WallStyle
				}
				if m.Maze.Grid[wx][wy] == maze.GOAL {
					color = render.GoalStyle
				}
			} else {
				color = render.MaskStyle
			}

			sx := offsetX + x
			sy := offsetY + y
			if sx >= 0 && sx < screenW && sy >= 0 && sy < screenH {
				scr.Grid[sx][sy] = color
			}
		}
	}

}
