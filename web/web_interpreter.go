package web

type Gameinfo struct {
	Name     string
	Type     string
	Origuser string
	Curruser string
	Note     string
}

func NewGameinfo() Gameinfo {
	info := Gameinfo{
		Name:     "Test",
		Type:     "Test4",
		Origuser: "Testo",
		Curruser: "Testi",
		Note:     "Test3",
	}
	return info
}
