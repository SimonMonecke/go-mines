package userinterface

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

const (
	ResizeExitCodeOk ExitCode = iota
	ResizeExitCodeQuit
)

func (ui *UserInterface) ShowResize() ExitCode {
	termbox.Sync()
	renderResize()

	for {
		ev := <-ui.eventQueue

		if ev.Type == termbox.EventResize {
			if ev.Width >= 99 && ev.Height >= 35 {
				termbox.Sync()
				return ResizeExitCodeOk
			}
			termbox.Sync()
			renderResize()
		} else if ev.Type == termbox.EventKey {
			if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return ResizeExitCodeQuit
			}
		}
	}
}

func renderResize() {
	terminalWidth, terminalHeight := termbox.Size()

	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	tbprintCentered(terminalHeight/2-1, termbox.ColorBlack, termbox.ColorWhite, "GO-MINES")
	tbprintCentered(terminalHeight/2+1, termbox.ColorWhite, termbox.ColorBlack, fmt.Sprintf("Min Terminal Size: 99x35. Actual Size %dx%d", terminalWidth, terminalHeight))

	tbprintCentered(terminalHeight-1, termbox.ColorWhite, termbox.ColorBlack, "[q] - Quit")
	termbox.Flush()
}
