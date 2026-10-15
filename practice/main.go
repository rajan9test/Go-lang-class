package main

import (
	"fmt"
	"strings"

	userModule "practice/user"
)

func main() {

	// db
	users := []userModule.User{
		{Id: 1, Name: "pratush", Email: "pratush@gmail.com"},
		{Id: 2, Name: "Aarav", Email: "aarav@gmail.com"},
	}

	//   fmt.Println(users)

	users = append(users, userModule.User{Id: 3, Name: "atul", Email: "atul@gmail.com"})

	ok := userModule.AddUser(&users, userModule.User{Id: 4, Name: "Sikha", Email: "sikha@gmail.com"})

	if ok {
		fmt.Println("user added in database")
	}
	// --------------------------------------------------------
	fmt.Println(users)

	user, found := userModule.FindById(users, 2)

	if found {
		fmt.Println(user)
	} else {
		fmt.Println("user didn't found")
	}

	user, found = userModule.FindById(users, 6)

	if found {
		fmt.Println(user)
	} else {
		fmt.Println("user didn't found")
	}

}

func emailNormalization(email string) (string, error) {

	//rajan  => ["ranjan"]
	//rajan@gmail.com   = ["rajan" "gmail.com"]   @hubspot.com@hubspot.com =

	// rajan@gmail.com =

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	emailPart := strings.SplitN(email, "@", 2)

	if len(emailPart) != 2 {
		return "", fmt.Errorf("this is not valid email")
	}

	local := emailPart[0]
	domain := emailPart[1]

	idx := strings.Index(local, "+")

	plusSplit := strings.Split(local, "+")
	// fmt.Println(plusSplit[len(plusSplit)-1])

	first := plusSplit[len(plusSplit)-1]

	nor := first + "@" + domain

	fmt.Println(idx)
	fmt.Println("==========================================================")
	fmt.Println(nor)
	return nor, nil

}
