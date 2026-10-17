package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "modernc.org/sqlite"
)
func Connect() (*sql.DB , error){
     

	// task :- we have to impmlement env helper fuction and this connection string should be in .env
	dns := "postgresql://neondb_owner:npg_cGPzB58UeklI@ep-billowing-glade-b4jhb8lj-pooler.c-6.us-east-2.aws.neon.tech/neondb?sslmode=require&channel_binding=require"

    database , err := sql.Open("pgx", dns)

	if err != nil{
		return  nil , fmt.Errorf("connection error :  %w", err)
	}

	err = database.Ping();

	if err != nil{
		 return nil, fmt.Errorf("database : %w" , err)
	}

	return database, nil

}


func MigrateinPostgres(database *sql.DB) error{
	 

	const userTable = `CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT  NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    create_at DATE DEFAULT CURRENT_TIMESTAMP,
    update_at DATE DEFAULT CURRENT_TIMESTAMP
	);`

	const college = `
	CREATE TABLE IF NOT EXISTS college (
    id INTEGER PRIMARY KEY,
	userId INTEGER,
    name TEXT  NOT NULL
	);`
     

	_, err := database.Exec(userTable)

	if err != nil{
		log.Println("user Table not create ", err);
		return err
	}
	

	_, err = database.Exec(college)

	if err != nil{
		log.Println("user college not create ", err);
		return err
	}
  return  nil
}
