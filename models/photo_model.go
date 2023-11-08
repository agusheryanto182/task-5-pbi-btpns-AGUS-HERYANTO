package models

type Photo struct {
	ID       int
	Title    string
	Caption  string
	PhotoURL string
	UserID   User
}
