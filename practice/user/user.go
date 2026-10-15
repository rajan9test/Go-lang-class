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


func UpdateUser(user *[]User, id int, value User) bool {
    for i := 0; i < len(*user); i++ {
        if (*user)[i].Id == id {
            (*user)[i] = value
            return true
        }
    }

    return false
}
























































































































































































































func DeleteUser(user *[]User, id int) bool {
    for i := 0; i < len(*user); i++ {
        if (*user)[i].Id == id {
            *user = append((*user)[:i], (*user)[i+1:]...)
            return true
        }
    }

    return false
}

