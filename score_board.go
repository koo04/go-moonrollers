package moonrollers

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	populatedCellStyle = lipgloss.NewStyle().
				Border(lipgloss.NormalBorder(), true).
				BorderForeground(lipgloss.Color("#404040")).
				Width(5).
				Height(1).
				Align(lipgloss.Center)
	emptyCellStyle = populatedCellStyle.Foreground(lipgloss.Color("#404040"))
)

func (g *Game) RenderPoints(short, debug bool) string {
	style := lipgloss.NewStyle().Margin(1)
	if debug {
		style = style.Border(lipgloss.ASCIIBorder()).Margin(0)
	}

	if short {
		out := make([]string, 0)
		for _, p := range g.Players {
			points := lipgloss.NewStyle().
				Foreground(lipgloss.Color(p.Color)).
				Render(p.Name + ": " + strconv.Itoa(p.Points))
			out = append(out, points)
		}

		return lipgloss.JoinVertical(lipgloss.Left, style.Padding(1).Render(strings.Join(out, "\n")))
	}

	rows := make([]string, 0)
	for row := range 10 {
		var rowCells []string
		for col := range 5 {
			rowCells = append(rowCells, g.RenderPointCell(row, col))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, rowCells...))
	}

	morePointsRow := make([]string, 0)
	for playerIndex := range 5 {
		morePointsRow = append(morePointsRow, g.RenderAdditionalPointsCell(playerIndex))
	}
	rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, morePointsRow...))

	return style.Render(strings.Join(rows, "\n"))
}

func (g *Game) RenderAdditionalPointsCell(playerIndex int) string {
	if playerIndex >= len(g.Players) {
		return emptyCellStyle.Render("")
	}

	points := g.Players[playerIndex].Points

	if points < 50 {
		return emptyCellStyle.Render("")
	}

	output := make([]string, 0)

	for range points / 50 {
		output = append(output, lipgloss.NewStyle().Foreground(g.Players[playerIndex].Color).Render("■"))
	}

	return populatedCellStyle.Render(strings.Join(output, " "))
}

func (g *Game) RenderPointCell(row int, col int) string {
	points := (row)*5 + col + 1
	var pips []string
	for _, player := range g.Players {
		if player.Points%50 == points {
			pip := lipgloss.NewStyle().Foreground(player.Color).Render("■")
			pips = append(pips, pip)
		}
	}
	if len(pips) == 0 {
		return emptyCellStyle.Render(strconv.Itoa(points))
	}
	return populatedCellStyle.Render(strings.Join(pips, ""))
}
