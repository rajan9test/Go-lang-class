package store

import (
	"database/sql"
	"fmt"
	"golang/model"
)

type UserStore struct {
	db *sql.DB
}

func NewUserStore( db *sql.DB) *UserStore{
	return &UserStore{
		db : db,
	}
}

func (s *UserStore) Create(user model.User) (model.User , error){
        
      err := s.db.QueryRow(
		`INSERT INTO users (name , email , password_hash)
		values($1, $2 , $3)
        RETURNING ID `,
		 user.Name , user.Email , user.Password_hash,
	  ).Scan(&user.ID)

	  if err != nil {
		 return model.User{} , fmt.Errorf("Insert Postgres : %w", err)
	  }

	return user,nil
}