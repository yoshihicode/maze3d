package maze

const (
	W     = 41
	H     = 23
	GoalX = W - 3
	GoalY = H - 3
)

const (
	EMPTY = 0 // path
	WALL  = 1 // wall
	GOAL  = 2 // goal
)

type Maze struct {
	Grid [W][H]int
}

func NewMaze() *Maze {
	m := &Maze{}
	m.init()
	return m
}

func (m *Maze) init() {
	// initialize with walls
	for x := range m.Grid {
		for y := range m.Grid[x] {
			m.Grid[x][y] = WALL
		}
	}

	startX, startY := 1, 1

	m.Grid[startX][startY] = EMPTY

	m.generate(startX, startY)
	m.breakWalls(0.10)
	m.createRoom(GoalX, GoalY, 3)
}
