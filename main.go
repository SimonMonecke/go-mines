package main

import (
	"log"
	"math/rand"
	"time"

	"fmt"

	termbox "github.com/nsf/termbox-go"
	"github.com/smonecke/go-mines/convert"
	"github.com/smonecke/go-mines/game"
	"github.com/smonecke/go-mines/generate"
	"github.com/smonecke/go-mines/highscore"
	"github.com/smonecke/go-mines/menu"
	"github.com/smonecke/go-mines/userinterface"
)

type state bool

const (
	running state = true
	quitted state = false
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ui := userinterface.Create()
	terminalWidth, terminalHeight := termbox.Size()
	if terminalWidth < 99 || terminalHeight < 35 {
		ui.Destroy()
		fmt.Printf("Error - Min Terminal Size: 99x35. Actual Size: %dx%d\n", terminalWidth, terminalHeight)
	} else {
		var err error

		for gameState := running; gameState == running && err == nil; {
			m := menu.New()

			switch ui.ShowMenu(m) {
			case userinterface.MenuExitCodeEasy:
				stringMap := generate.EasyStringMap()
				gameState, err = playTheGame(stringMap, ui, game.Easy)
			case userinterface.MenuExitCodeNormal:
				stringMap := generate.NormalStringMap()
				gameState, err = playTheGame(stringMap, ui, game.Normal)
			case userinterface.MenuExitCodeHard:
				stringMap := generate.HardStringMap()
				gameState, err = playTheGame(stringMap, ui, game.Hard)
			case userinterface.MenuExitCodeHighscores:
				gameState, err = showTheHighscores(ui)
			case userinterface.MenuExitCodeExit:
				gameState = quitted
			default:
				log.Fatal("unknown state")
			}
		}
		ui.Destroy()
		if err != nil {
			fmt.Printf("Error - %s\n", err)
		}
	}
}

func playTheGame(stringMap [][]string, ui userinterface.UserInterface, mode game.Mode) (state, error) {
	board := convert.FromStringMapToBoard(stringMap)
	g := game.New(&board, mode)
	var newState state
	exitCode, err := ui.ShowGame(g)
	if err != nil {
		return quitted, err
	}

	switch exitCode {
	case userinterface.GameExitCodeWin:
		newState, err = showTheHighscores(ui)
		if err != nil {
			return quitted, err
		}
	case userinterface.GameExitCodeLoose:
		newState = running
	case userinterface.GameExitCodeNew:
		newState = running
	case userinterface.GameExitCodeQuit:
		newState = quitted
	default:
		log.Fatal("unknown state")
	}
	return newState, nil
}

func showTheHighscores(ui userinterface.UserInterface) (state, error) {
	hl, err := highscore.LoadList()
	if err != nil {
		return quitted, err
	}
	var newState state

	switch ui.ShowHighscores(hl) {
	case userinterface.HighscoreExitCodeMenu:
		newState = running
	case userinterface.HighscoreExitCodeQuit:
		newState = quitted
	default:
		log.Fatal("unknown state")
	}
	return newState, nil
}
