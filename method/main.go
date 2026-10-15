package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

type AllUser []User

// func (u User) GetName() string {
// 	return u.Name
// }

// func (u User) GetDetails() string{
// 	return fmt.Sprintf("Id : %d  Name: %s Email: %s " , u.ID,  u.Name, u.Email)
// }

// func (u *User) Rename(name string)  {

// 	u.Name = name

// }

func (us *AllUser) Add(u User) {

	*us = append(*us, u)

}

func (us AllUser) findUser(id int) (User , bool){
       
       for _ , us := range us{
		     
		  if(us.ID == id){
			return us , true
		  }
	   }

	   return  User{} , false 


}

func (us AllUser) FindbyEmail(email string) (User, bool){
      

	for _ , user := range us{
		  
		if(user.Email == email){
			return user, true
		}
	}

	return User{}, false
}

func (us *AllUser) DeleteUser(id int) bool{

	 for i , user := range *us {
             
		if user.ID == id{
             *us = append((*us)[:i] , (*us)[i+1:]... )

			 return true
		}
	 }
	  
	return false
}

func (us *AllUser) UpdateUser(id int, Name string , Email string) (User, bool){
  
     

    


	return User{}, false
}



func main() {

	var users AllUser

	users.Add(User{ID: 1, Name: "Pratush", Email: "pratush@gmail.com"})
		users.Add(User{ID: 2, Name: "Priya", Email: "Priya@gmail.com"})

			users.Add(User{ID: 3, Name: "Atul", Email: "Atul@gmail.com"})

				users.Add(User{ID: 4, Name: "Deepti", Email: "Deepti@gmail.com"})

	  fmt.Println(users.findUser(3))


	// fmt.Print(users)

}
