package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/amren1254/gin-docker/model"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	username = "root" //os.Getenv("DB_USERNAME")
	password = "root" //os.Getenv("DB_PASSWORD")
	hostname = "127.0.0.1:3306"
	dbname   = "power_zone"
)

type DatabaseRepository struct {
	DB *sql.DB
}

func NewDatabaseRepository(database *sql.DB) *DatabaseRepository {
	return &DatabaseRepository{
		DB: database,
	}
}

func dataSourceName(databaseName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func (db *DatabaseRepository) InitDatabaseConnection() *sql.DB {
	var err error
	newDb, err := sql.Open("mysql", dataSourceName(dbname))
	if err != nil {
		log.Printf("Error while opening connection with database")
		return nil
	}
	//defer DB.Close()

	newDb.SetMaxOpenConns(10)
	newDb.SetMaxIdleConns(10)
	newDb.SetConnMaxLifetime(time.Minute * 5)
	return newDb
}

func (db *DatabaseRepository) InsertUser(user model.UserDetails) {
	query := "INSERT INTO user_details(username,password) VALUES(?,?)"
	insert, err := db.DB.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	//encrypt password before storing
	encryptedPassword, err := encrypt(user.Password)
	if err != nil {
		log.Println(err)
	}
	response, err := insert.Exec(user.Username, encryptedPassword)
	insert.Close()
	if err != nil {
		log.Println(err)
	}
	log.Println(response.RowsAffected())
}

func (db *DatabaseRepository) retrieveUserDetails(user model.UserDetails) (retrievedUserCredential model.UserDetails, err error) {

	err = db.DB.QueryRow("select username, password from user_details where username = ?", user.Username).
		Scan(&retrievedUserCredential.Username, &retrievedUserCredential.Password)
	if err != nil {
		log.Println(err)
		return retrievedUserCredential, err
	}
	return retrievedUserCredential, nil
}
