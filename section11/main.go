package main

import "fmt"

type T struct {
	User
}

type User struct {
	Name string
	Age int
	// X, Y int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// こちらのやり方が望ましい。ポインタ型。
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	var user1 User
	fmt.Println(user1)

	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)

	user2.Name = "user2"
	user2.Age = 10
	fmt.Println(user2)

	user3 := User{Name: "user3"}
	user3.SayName()

	user3.SetName2("A")
	user3.SayName()

	t := T{User: User{Name: "user1", Age: 100}}
	fmt.Println(t)

	fmt.Println(t.User)

}