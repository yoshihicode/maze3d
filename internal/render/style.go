package render

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Bg(color int) lipgloss.Style {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(fmt.Sprintf("%d", color)))
}

func GrayScale(dist float64) int {

	minGray := 240
	maxGray := 250

	gray := maxGray - int(dist*3)

	if gray < minGray {
		gray = minGray
	}
	if gray > maxGray {
		gray = maxGray
	}

	return gray
}

var FloorStyle = 238
var GoalStyle = 2
var BackgroundStyle = 0
var MaskStyle = 236
var WallStyle = 255
var PlayerStyle = 1
