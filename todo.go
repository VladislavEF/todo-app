package todo

type TodoList struct {
	Id          int    `json:"-"`
	Titel       string `json:"title"`
	Description string `json:"description"`
}
type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"-"`
	Titel       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
