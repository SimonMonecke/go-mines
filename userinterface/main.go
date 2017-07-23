package userinterface

import (
	"github.com/nsf/termbox-go"
)

type ExitCode int

type UserInterface struct {
	eventQueue chan termbox.Event
}

func Create() UserInterface {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()
	return UserInterface{eventQueue: eventQueue}
}

func (ui *UserInterface) Destroy() {
	termbox.Close()
}

func tbprintCentered(y int, fg, bg termbox.Attribute, msg string) {
	terminalwidth, _ := termbox.Size()
	x := terminalwidth/2 - len(msg)/2
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}
