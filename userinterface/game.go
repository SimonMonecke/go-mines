package userinterface

import (
	"strings"
	"time"

	"fmt"

	"github.com/nsf/termbox-go"
	"github.com/smonecke/go-mines/game"
	"github.com/smonecke/go-mines/highscore"
)

const (
	GameExitCodeWin ExitCode = iota
	GameExitCodeLoose
	GameExitCodeQuit
	GameExitCodeNew
)

const (
	StatusWon int = iota
	StatusLost
	StatusRunning
)

func (ui *UserInterface) ShowGame(g *game.Game) (ExitCode, error) {
	start := time.Now()
	ticker := time.NewTicker(time.Second)
	var paused time.Duration
	defer ticker.Stop()
	render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))

	for {
		select {
		case <-ticker.C:
			render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
		case ev := <-ui.eventQueue:
			if ev.Type == termbox.EventResize {
				if ev.Width < 99 || ev.Height < 35 {
					startPause := time.Now()
					exitcode := ui.ShowResize()
					if exitcode == ResizeExitCodeQuit {
						return GameExitCodeQuit, nil
					}
					paused = time.Since(startPause)
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				} else {
					termbox.Sync()
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				}
			} else if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyArrowDown {
					g.SelectSquareBottom()
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				} else if ev.Key == termbox.KeyArrowUp {
					g.SelectSquareTop()
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				} else if ev.Key == termbox.KeyArrowRight {
					g.SelectSquareRight()
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				} else if ev.Key == termbox.KeyArrowLeft {
					g.SelectSquareLeft()
					render(g, StatusRunning, int((time.Since(start) - paused).Seconds()))
				} else if ev.Key == termbox.KeyEnter {
					g.Uncover()
					elapsed := int((time.Since(start) - paused).Seconds())
					if g.IsWon() {
						render(g, StatusWon, elapsed)
						for {
							ev := <-ui.eventQueue
							if ev.Type == termbox.EventResize {
								if ev.Width < 99 || ev.Height < 35 {
									exitcode := ui.ShowResize()
									if exitcode == ResizeExitCodeQuit {
										return GameExitCodeQuit, nil
									}
									render(g, StatusWon, elapsed)
								} else {
									termbox.Sync()
									render(g, StatusWon, elapsed)
								}
							} else if ev.Type == termbox.EventKey {
								if ev.Key == termbox.KeyEnter {
									t := time.Now()
									if elapsed > 999 {
										elapsed = 999
									}

									err := highscore.AddEntry(g.Mode(), t.Format("2006-01-02"), elapsed)
									if err != nil {
										return GameExitCodeQuit, err
									}
									return GameExitCodeWin, nil
								} else if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
									return GameExitCodeQuit, nil
								}
							}
						}
					} else if g.IsLost() {
						render(g, StatusLost, elapsed)
						for {
							ev := <-ui.eventQueue
							if ev.Type == termbox.EventResize {
								if ev.Width < 99 || ev.Height < 35 {
									exitcode := ui.ShowResize()
									if exitcode == ResizeExitCodeQuit {
										return GameExitCodeQuit, nil
									}
									render(g, StatusLost, elapsed)
								} else {
									termbox.Sync()
									render(g, StatusLost, elapsed)
								}
							} else if ev.Type == termbox.EventKey {
								if ev.Key == termbox.KeyEnter {
									return GameExitCodeLoose, nil
								} else if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
									return GameExitCodeQuit, nil
								}
							}
						}
					}
					render(g, StatusRunning, elapsed)
				} else if ev.Ch == 'n' {
					return GameExitCodeNew, nil
				} else if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
					return GameExitCodeQuit, nil
				}
			}
		}
	}
}

func render(game *game.Game, status int, elapsed int) {
	terminalWidth, terminalHeight := termbox.Size()
	top := (terminalHeight-1)/2 - game.Height()/2
	left := terminalWidth/2 - (game.Width()*4+len("Elapsed Time: 999s")+1)/2
	if elapsed > 999 {
		elapsed = 999
	}

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	tbprint(left+1, top+0, termbox.ColorWhite, termbox.ColorBlack, strings.Repeat("____", game.Width()-1))
	tbprint(left+(game.Width()-1)*4+1, top+0, termbox.ColorWhite, termbox.ColorBlack, "___")
	for row := 0; row < game.Height(); row++ {
		tbprint(left, top+row+1, termbox.ColorWhite, termbox.ColorBlack, "|")
		for col := 0; col < game.Width(); col++ {
			square := game.GetSquare(col, row)
			var squareContent string
			var bgColor termbox.Attribute
			fgColor := termbox.ColorWhite
			if square.IsUncovered() {
				if !square.IsMine() {
					if square.CountOfMinesNeighbour() == 0 {
						squareContent = "___"
						bgColor = termbox.ColorGreen
					} else {
						squareContent = fmt.Sprintf("_%d_", square.CountOfMinesNeighbour())
						bgColor = termbox.ColorBlue
					}
				} else {
					squareContent = "_X_"
					bgColor = termbox.ColorRed
				}
			} else {
				squareContent = "___"
				bgColor = termbox.ColorBlack
			}
			if game.SelectedCol() == col && game.SelectedRow() == row {
				bgColor = termbox.ColorYellow
				fgColor = termbox.ColorBlack
			}
			tbprint(left+col*4+1, top+row+1, fgColor, bgColor, squareContent)
			tbprint(left+col*4+4, top+row+1, termbox.ColorWhite, termbox.ColorBlack, "|")
		}
	}
	_, height := termbox.Size()
	if status == StatusLost {
		tbprint(left+game.Width()*4+2, top+1, termbox.ColorWhite, termbox.ColorBlack, "You Lose!")
		tbprintCentered(height-1, termbox.ColorWhite, termbox.ColorBlack, "[Enter] - Main Menu  |  [q] - Quit")
	} else if status == StatusWon {
		tbprint(left+game.Width()*4+2, top+1, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("You Win! - %ds", elapsed))
		tbprintCentered(height-1, termbox.ColorWhite, termbox.ColorBlack, "[Enter] - Highscores  |  [q] - Quit")
	} else {
		tbprint(left+game.Width()*4+2, top+1, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("Elapsed Time: %ds", elapsed))
		tbprintCentered(height-1, termbox.ColorWhite, termbox.ColorBlack, "[Up]/[Down]/[Left]/[Right] - Navigate  |  [Enter] - Uncover Field  |  [n] - New Game  |  [q] - Quit")
	}
	termbox.Flush()
}
