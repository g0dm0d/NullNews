package db

type User struct {
	ID        int
	Username  string
	Firstname string
	Lastname  string
	Password  string
	Email     string
	Admin     bool
}

type News struct {
	ID    int
	Title string
	Text  string
	User  int
}
