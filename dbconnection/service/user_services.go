package service

import (
	"fmt"
	"golang/model"
	"golang/userstore"
	"strings"
)

type UserSerice struct {
	store *store.UserStore
}


func NewUserServices(userStore *store.UserStore) *UserSerice{

	return &UserSerice{
		store: userStore,
	}
}

func (s *UserSerice) Register(name string, email string, pass string)(model.User, error){
	  
     name = strings.TrimSpace(name)
	 email = strings.TrimSpace(email)

	 if name == "" || email == "" || pass == ""{
		return  model.User{} , fmt.Errorf("name , email and password is required")
	 }

	 user , err := s.store.Create(model.User{
		Name: name,
		Email: email,
		Password_hash: pass,
	 })

	 if err != nil {
		return model.User{}, err
	 }

	 return  user , nil
}