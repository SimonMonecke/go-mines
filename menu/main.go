package menu

type ItemId int

const (
	MenuItemEasy ItemId = iota
	MenuItemNormal
	MenuItemHard
	MenuItemHighscores
	MenuItemExit
)

type Item struct {
	isSelected bool
	id         ItemId
}

func (mi *Item) IsSelected() bool {
	return mi.isSelected
}

func (mi *Item) GetId() ItemId {
	return mi.id
}

type Menu struct {
	items []Item
}

func New() *Menu {
	return &Menu{
		items: []Item{
			Item{isSelected: false, id: MenuItemEasy},
			Item{isSelected: true, id: MenuItemNormal},
			Item{isSelected: false, id: MenuItemHard},
			Item{isSelected: false, id: MenuItemHighscores},
			Item{isSelected: false, id: MenuItemExit}}}
}

func (m *Menu) GetItems() []Item {
	return m.items
}

func (m *Menu) GetSelectedItemId() ItemId {
	var selectedId ItemId
	for i := 0; i < len(m.items); i++ {
		if m.items[i].isSelected {
			selectedId = m.items[i].id
		}
	}
	return selectedId
}

func (m *Menu) SelectPrevItem() {
	var selectedMenuItemIndex int
	for i := 0; i < len(m.items); i++ {
		if m.items[i].isSelected {
			m.items[i].isSelected = false
			selectedMenuItemIndex = i
			break
		}
	}
	if selectedMenuItemIndex == 0 {
		m.items[len(m.items)-1].isSelected = true
	} else {
		m.items[selectedMenuItemIndex-1].isSelected = true
	}
}

func (m *Menu) SelectNextItem() {
	var selectedMenuItemIndex int
	for i := 0; i < len(m.items); i++ {
		if m.items[i].isSelected {
			m.items[i].isSelected = false
			selectedMenuItemIndex = i
			break
		}
	}
	if selectedMenuItemIndex == len(m.items)-1 {
		m.items[0].isSelected = true
	} else {
		m.items[selectedMenuItemIndex+1].isSelected = true
	}
}
