package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

//scoreBoard keeps track of the score
type scoreBoard map[string]team

//team is the win/loss record of the team
type team struct {
	name                            string
	played, win, loss, draw, points int
}

// Tally counts up the wins and losses for a sports season
func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	board := make(scoreBoard)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Split(line, ";")
		if len(fields) != 3 {
			return fmt.Errorf("game not properly formatted: %s", line)
		}

		a, b := board[fields[0]], board[fields[1]]
		a.name, b.name = fields[0], fields[1]

		a.played++
		b.played++
		switch fields[2] {
		case "win":
			a.points += 3
			a.win++
			b.loss++
		case "loss":
			a.loss++
			b.points += 3
			b.win++
		case "draw":
			a.points++
			a.draw++
			b.points++
			b.draw++
		default:
			return fmt.Errorf("unknown win condition: %s", line)
		}
		board[fields[0]] = a
		board[fields[1]] = b
	}
	teams := make([]team, 0, len(board))
	for _, t := range board {
		teams = append(teams, t)
	}
	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points == teams[j].points {
			return teams[i].name < teams[j].name
		}
		return teams[i].points > teams[j].points
	})
	header := "Team                           | MP |  W |  D |  L |  P"
	bodyFmtString := "%-31s| %2d | %2d | %2d | %2d | %2d\n"

	fmt.Fprintln(writer, header)
	for _, team := range teams {
		fmt.Fprintf(writer, bodyFmtString,
			team.name, team.played,
			team.win, team.draw,
			team.loss, team.points)
	}
	return nil
}
