package food

type Restaurant struct {
	Name string
	Time string
	MenuList map[string][]Menu
}

type Menu struct {
	Menu string
	Price string
}
