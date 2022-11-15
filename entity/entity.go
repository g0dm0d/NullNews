package entity

type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"text"`
	Author string `json:"author"`
}

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Permission int    `json:"permission"`
	Password   string `json:"password"`
}
