package main

import "fmt"

type Address struct {
	Region string
	Street string
	No     string
}

func (addr Address) Addr() string {
	return fmt.Sprintf("%s-%s-%s ", addr.Region, addr.Street, addr.No)
}

type User struct {
	ID   int
	Name string
	Addr Address
}

func (user User) User() string {
	return fmt.Sprintf("[%d]%s:%s", user.ID, user.Name, user.Addr)
}

func (user User) String() string {
	return fmt.Sprintf("[%d]%s:%s", user.ID, user.Name, user.Addr)
}

func main() {
	var u User = User{
		ID:   1,
		Name: "kk",
		Addr: Address{
			Region: "89328",
			Street: "iuiods",
			No:     "iods",
		},
	}

	fmt.Printf("%#v\n", u)
	fmt.Println(u.User())      //[1]kk:{89328 iuiods iods}
	fmt.Println(u.Addr.Region) //89328
	fmt.Println(u)             //[1]kk:{89328 iuiods iods}

}
