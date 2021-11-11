package users



type address struct {
	Region  string
	streect string
	no      string
}


type User struct {
	ID       int
	Name     string
	birthday string
	addr     address
}


