package userinterface

import (
	"github.com/nsf/termbox-go"
	"github.com/smonecke/go-mines/menu"
)

const (
	MenuExitCodeEasy ExitCode = iota
	MenuExitCodeNormal
	MenuExitCodeHard
	MenuExitCodeHighscores
	MenuExitCodeExit
)

func (ui *UserInterface) ShowMenu(m *menu.Menu) ExitCode {
	renderMenu(m)
	for {
		ev := <-ui.eventQueue
		if ev.Type == termbox.EventResize {
			if ev.Width < 99 || ev.Height < 35 {
				exitcode := ui.ShowResize()
				if exitcode == ResizeExitCodeQuit {
					return MenuExitCodeExit
				}
				renderMenu(m)
			} else {
				termbox.Sync()
				renderMenu(m)
			}
		} else if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyArrowDown {
				m.SelectNextItem()
				renderMenu(m)
			} else if ev.Key == termbox.KeyArrowUp {
				m.SelectPrevItem()
				renderMenu(m)
			} else if ev.Key == termbox.KeyEnter {
				menuItemId := m.GetSelectedItemId()
				if menuItemId == menu.MenuItemEasy {
					return MenuExitCodeEasy
				} else if menuItemId == menu.MenuItemNormal {
					return MenuExitCodeNormal
				} else if menuItemId == menu.MenuItemHard {
					return MenuExitCodeHard
				} else if menuItemId == menu.MenuItemHighscores {
					return MenuExitCodeHighscores
				} else if menuItemId == menu.MenuItemExit {
					return MenuExitCodeExit
				}
			} else if ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD {
				return MenuExitCodeExit
			}
		}
	}
}

func menuItemIdToString(menuItemId menu.ItemId) string {
	var itemString string
	if menuItemId == menu.MenuItemEasy {
		itemString = "EASY"
	} else if menuItemId == menu.MenuItemNormal {
		itemString = "NORMAL"
	} else if menuItemId == menu.MenuItemHard {
		itemString = "HARD"
	} else if menuItemId == menu.MenuItemHighscores {
		itemString = "HIGHSCORES"
	} else if menuItemId == menu.MenuItemExit {
		itemString = "EXIT"
	}
	return itemString
}

func renderMenu(m *menu.Menu) {
	_, terminalHeight := termbox.Size()
	termbox.Clear(termbox.ColorWhite, termbox.ColorBlack)
	menuItems := m.GetItems()
	countOfLines := 10
	yOffset := (terminalHeight-1)/2 - countOfLines/2

	tbprintCentered(yOffset, termbox.ColorBlack, termbox.ColorWhite, "GO-MINES")
	for i, mi := range menuItems {
		if mi.IsSelected() {
			if i == len(menuItems)-1 {
				tbprintCentered(yOffset+i+5, termbox.ColorBlack, termbox.ColorWhite, menuItemIdToString(mi.GetId()))
			} else if i == len(menuItems)-2 {
				tbprintCentered(yOffset+i+4, termbox.ColorBlack, termbox.ColorWhite, menuItemIdToString(mi.GetId()))
			} else {
				tbprintCentered(yOffset+i+3, termbox.ColorBlack, termbox.ColorWhite, menuItemIdToString(mi.GetId()))
			}
		} else {
			if i == len(menuItems)-1 {
				tbprintCentered(yOffset+i+5, termbox.ColorWhite, termbox.ColorBlack, menuItemIdToString(mi.GetId()))
			} else if i == len(menuItems)-2 {
				tbprintCentered(yOffset+i+4, termbox.ColorWhite, termbox.ColorBlack, menuItemIdToString(mi.GetId()))
			} else {
				tbprintCentered(yOffset+i+3, termbox.ColorWhite, termbox.ColorBlack, menuItemIdToString(mi.GetId()))
			}
		}
	}
	_, height := termbox.Size()
	tbprintCentered(height-1, termbox.ColorWhite, termbox.ColorBlack, "[Up]/[Down] - Navigate  |  [Enter] - Select Item  |  [q] - Quit")
	termbox.Flush()
}
