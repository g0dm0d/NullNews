package entity

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	User  string
}

type User struct {
	ID        int
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
