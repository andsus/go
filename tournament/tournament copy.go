package tournament

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"sort"
// 	"strings"
// )

//scoreBoard keeps track of the score
// type scoreBoard map[string]*team

//team is the win/loss record of the team
// type team struct {
// 	name                            string
// 	played, win, loss, draw, points int
// }

// Tally counts up the wins and losses for a sports season
// func Tally(reader io.Reader, writer io.Writer) error {
// 	scanner := bufio.NewScanner(reader)
// 	board := make(scoreBoard)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if line == "" || strings.HasPrefix(line, "#") {
// 			continue
// 		}
// 		if err := board.addGame(line); err != nil {
// 			return err
// 		}
// 	}

// 	teams := board.getTeams()
// 	sort.Slice(teams, func(i, j int) bool {
// 		if teams[i].points == teams[j].points {
// 			return teams[i].name < teams[j].name
// 		}
// 		return teams[i].points > teams[j].points
// 	})
// 	header := "Team                           | MP |  W |  D |  L |  P"
// 	bodyFmtString := "%-31s| %2d | %2d | %2d | %2d | %2d\n"

// 	fmt.Fprintln(writer, header)
// 	for _, team := range teams {
// 		fmt.Fprintf(writer, bodyFmtString,
// 			team.name, team.played,
// 			team.win, team.draw,
// 			team.loss, team.points)
// 	}
// 	return nil
// }

// addGame validates and adds a game to the scoreboard.
// func (b scoreBoard) addGame(game string) error {
// 	fields := strings.Split(game, ";")
// 	if len(fields) != 3 {
// 		return fmt.Errorf("game not properly formatted: %s", game)
// 	}
// 	home, homeOk := b[fields[0]]
// 	away, awayOk := b[fields[1]]
// 	if !homeOk {
// 		home = &team{name: fields[0]}
// 		b[fields[0]] = home
// 	}
// 	if !awayOk {
// 		away = &team{name: fields[1]}
// 		b[fields[1]] = away
// 	}

// 	home.played++
// 	away.played++
// 	switch fields[2] {
// 	case "win":
// 		home.points += 3
// 		home.win++
// 		away.loss++
// 	case "loss":
// 		home.loss++
// 		away.points += 3
// 		away.win++
// 	case "draw":
// 		home.points++
// 		home.draw++
// 		away.points++
// 		away.draw++
// 	default:
// 		return fmt.Errorf("unknown win condition: %s", game)
// 	}
// 	return nil
// }

// getTeams converts the scoreboard into a list.
// func (b scoreBoard) getTeams() []team {
// 	var teams []team
// 	for _, team := range b {
// 		teams = append(teams, *team)
// 	}
// 	return teams
// }
