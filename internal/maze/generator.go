package maze

import (
	"math/rand"
	"time"
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func (m *Maze) shuffle(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
}

// perfect maze generation (digging paths=0 from walls=1)
func (m *Maze) generate(x, y int) {
	dirs := []int{0, 1, 2, 3}
	m.shuffle(dirs)

	for _, d := range dirs {
		nx := x + dx[d]*2
		ny := y + dy[d]*2

		if nx > 0 && nx < W-1 && ny > 0 && ny < H-1 && m.Grid[nx][ny] == WALL {

			m.Grid[x+dx[d]][y+dy[d]] = EMPTY
			m.Grid[nx][ny] = EMPTY

			m.generate(nx, ny)
		}
	}
}

// make imperfect maze
func (m *Maze) breakWalls(percent float64) {

	total := W * H
	breakCount := int(float64(total) * percent)

	for range breakCount {

		x := rnd.Intn(W-2) + 1
		y := rnd.Intn(H-2) + 1

		if m.Grid[x][y] == WALL {

			open := 0

			for d := range 4 {
				if m.Grid[x+dx[d]][y+dy[d]] == EMPTY {
					open++
				}
			}

			if open >= 2 {
				m.Grid[x][y] = EMPTY
			}
		}
	}
}

// goal room
func (m *Maze) createRoom(cx, cy, size int) {

	half := size / 2

	for y := cy - half; y <= cy+half; y++ {
		for x := cx - half; x <= cx+half; x++ {

			if x > 0 && x < W-1 && y > 0 && y < H-1 {
				m.Grid[x][y] = EMPTY
			}
		}
	}
	m.Grid[cx][cy] = GOAL
}
