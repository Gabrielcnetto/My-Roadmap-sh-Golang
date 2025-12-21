package main

type user struct {
	fname string
	lname string
	age   int
}

func getInfo(*user) {

}
func main() {
	var appUser = user{
		fname: "Habib",
		lname: "Sheikh",
		age:   21,
	}
	getInfo(&appUser)
}
