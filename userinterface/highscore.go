package userinterface

import (
	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/smonecke/go-mines/highscore"
)

const (
	HighscoreExitCodeMenu ExitCode = iota
	HighscoreExitCodeQuit
)

func (ui *UserInterface) ShowHighscores(hl highscore.List) ExitCode {
	renderHighscoreList(hl)

	for {
		ev := <-ui.eventQueue
		if ev.Type == termbox.EventResize {
			if ev.Width < 99 || ev.Height < 35 {
				exitcode := ui.ShowResize()
				if exitcode == ResizeExitCodeQuit {
					return MenuExitCodeExit
				}
				renderHighscoreList(hl)
			} else {
				termbox.Sync()
				renderHighscoreList(hl)
			}
		} else if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEnter {
				return HighscoreExitCodeMenu
			} else if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return HighscoreExitCodeQuit
			}
		}
	}
}

func longestEntryInAllSublists(hl highscore.List) int {
	longestEntry := 0

	for _, sublist := range []highscore.Entries{hl.Easy, hl.Normal, hl.Hard} {
		for i := 0; i < len(sublist); i++ {
			if len(fmt.Sprintf("%s - %d", sublist[i].Date, sublist[i].DurationInSeconds)) > longestEntry {
				longestEntry = len(fmt.Sprintf("%s - %d", sublist[i].Date, sublist[i].DurationInSeconds))
			}
		}
	}
	return longestEntry
}

func printSublist(sublist highscore.Entries, lengthOfLongestEntry int, terminalWidth int, line *int) {
	for i := 0; i < len(sublist); i++ {
		tbprint(terminalWidth/2-lengthOfLongestEntry/2, *line, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("%s - %d", sublist[i].Date, sublist[i].DurationInSeconds))
		*line++
	}
	if len(sublist) == 0 {
		tbprintCentered(*line, termbox.ColorWhite, termbox.ColorBlack, "no entries")
		*line++
	}
}

func countSublistsEntries(hl highscore.List) int {
	countOfEntries := 0

	for _, sublist := range []highscore.Entries{hl.Easy, hl.Normal, hl.Hard} {
		if len(sublist) > 1 {
			countOfEntries += len(sublist)
		} else {
			countOfEntries++
		}
	}
	return countOfEntries
}

func renderHighscoreList(hl highscore.List) {
	lengthOfLongestEntry := longestEntryInAllSublists(hl)
	terminalWidth, terminalHeight := termbox.Size()
	countOfLines := 7 + countSublistsEntries(hl)
	line := (terminalHeight-1)/2 - countOfLines/2

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	tbprintCentered(line, termbox.ColorBlack, termbox.ColorWhite, "HIGHSCORES")
	line += 2

	tbprintCentered(line, termbox.ColorWhite, termbox.ColorBlack, "EASY")
	line++
	printSublist(hl.Easy, lengthOfLongestEntry, terminalWidth, &line)
	line++

	tbprintCentered(line, termbox.ColorWhite, termbox.ColorBlack, "NORMAL")
	line++
	printSublist(hl.Normal, lengthOfLongestEntry, terminalWidth, &line)
	line++

	tbprintCentered(line, termbox.ColorWhite, termbox.ColorBlack, "HARD")
	line++
	printSublist(hl.Hard, lengthOfLongestEntry, terminalWidth, &line)

	tbprintCentered(terminalHeight-1, termbox.ColorWhite, termbox.ColorBlack, "[ENTER] - Main Menu  |  [q] - Quit")
	termbox.Flush()
}
