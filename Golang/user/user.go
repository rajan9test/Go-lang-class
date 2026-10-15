package user

import "fmt"


func  FindById(user []User, id int) (User, bool) {
	for i := 0; i < len(user); i++ { // main (users) => adduser(users)

		if user[i].Id == id {
			return user[i], true
		}

	}
	userNil := User{}
	return userNil, false

}

func AddUser(user *[]User, SingleUser User) bool {

	// how i can add anything in slice
	*user = append(*user, SingleUser)

	fmt.Println(*user)

	return true
}

func UpdatUser(user *[]User, id int, value User) bool {



	return false
}

func DeleteUser(user *[]User, id int) bool{
     


	return false
}
