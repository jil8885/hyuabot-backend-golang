package food

type Restaurant struct {
	Name string
	Time string	`firestore:"time,omitempty"`
	MenuList map[string][]Menu `firestore:"menuList,omitempty"`
}

type Menu struct {
	Menu string
	Price string
}
