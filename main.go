package main

import "fmt"

func User() {
	var username string = "manish"
	fmt.Println(username)
	fmt.Fprint("the username type is = %T \n", username)

	var isLoggedIn bool = "true"
	fmt.Println("username is loggedin =  ", isLoggedIn)
}
