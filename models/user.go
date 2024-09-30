package models

import (
	"database/sql"
	"fmt"
	"ikct-ed/config"
	"ikct-ed/utility"
	"log"
)

type User struct {
	ID        int    `form:"id"`
	Name      string `form:"username" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Password  string `form:"password"`
	RoleId    int    `form:"role"`
	Status    bool   `form:"status"`
	LastLogin string
	UserAuth  UserAuth
}

type UserStorage struct {
	User *User
}

type UserAuth struct {
	UserID   int64
	JWTToken string `json:"jwt_token"`
}

// CheckUser
// Input: email
// Output: boolean, error
// Desc: This method returns true if there is a user exists with the given email.
func CheckUser(email string) (bool, error) {
	id := 0
	db, err := config.GetDB2()
	if err != nil {
		return false, err
	}
	defer db.Close()
	query := "SELECT id from users where LOWER(email)=$1"
	row := db.QueryRow(query, email)
	err = row.Scan(&id)
	if id > 0 {
		return true, err
	}
	return false, err
}

// New
// Input: User, passwordEncripted
// Output: error
// Desc: This method inserts a new record in the user table and also inserts a new record
// in the user2role table for mapping purposes.
func (u *User) New(newUser User, passwordEncripted string) error {
	db, err := config.GetDB2()
	if err != nil {
		return err
	}
	defer db.Close()
	//query to insert a new record in the users table
	query := `INSERT INTO users(
	name, email, password, status)
	VALUES ($1,$2,$3,$4) RETURNING id`
	var userId int
	err = db.QueryRow(query, newUser.Name, newUser.Email, passwordEncripted, newUser.Status).Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func GetAdminDetails(email string) (User, error) {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("GetAdminDetails: Failed while connecting with database with error: ", err)
		return User{}, err
	}
	defer db.Close()

	query := ` SELECT id, name, email, password, status FROM users WHERE email = $1`
	var (
		id       int
		name     sql.NullString
		emailId  string
		password string
		status   sql.NullBool
	)
	err = db.QueryRow(query, email).Scan(
		&id,
		&name,
		&emailId,
		&password,
		&status,
	)
	if err != nil {
		log.Println("GetAdminDetails: Failed while executing the query with error: ", err)
		return User{}, err
	}

	user := User{
		ID:       id,
		Name:     utility.SQLNullStringToString(name),
		Email:    emailId,
		Password: password,
		Status:   utility.SQLNullBoolToBool(status),
	}

	return user, nil
}

// StoreJwtSessionInDB METHOD, stores a new JWT session in the database. It connects to the database, executes an SQL INSERT
// query to save the JWT token along with its validity period (set to 5 years from the current timestamp) and the tutor's ID,
// and handles any errors that occur during the process.
func (userAuth UserAuth) StoreJwtSessionInDB() error {

	db, err := config.GetDB2()
	if err != nil {
		log.Println("[ERROR] StoreJwtSessionInDB: Failed to connect with database with error: ", err)
		return err
	}
	defer db.Close()
	query := ` INSERT INTO session(user_token,valid_until,user_id)VALUES($1,CURRENT_TIMESTAMP + INTERVAL '5 years',$2)`
	_, err = db.Exec(query, userAuth.JWTToken, userAuth.UserID)
	if err != nil {
		log.Println("[ERROR] StoreJwtSessionInDB: Failed to execute the query with error: ", err)
		return err
	}
	return nil
}

func GetUserProfileByToken(tokenString string) (User, error) {
	var (
		userID    int64
		user      User
		name      sql.NullString
		email     string
		password  string
		status    sql.NullBool
		lastLogin sql.NullTime
	)
	db, err := config.GetDB2()
	if err != nil {
		log.Println("[ERROR] GetUserProfileByToken: Failed to connect with database with error: ", err)
		return User{}, err
	}
	defer db.Close()

	query := ` SELECT 
				u.id,
				u.name,
				u.email,
				u.password,
				u.status,
				u.last_login
			FROM
				users as u,
				session as s
			WHERE
				u.id= s.user_id AND
				s.user_token = $1
				`
	fmt.Println("***query", query)
	err = db.QueryRow(query, tokenString).Scan(&userID, &name, &email, &password, &status, &lastLogin)
	if err != nil {
		log.Println("[ERROR] GetUserProfileByToken: Failed to execute the query with error: ", err)
		return User{}, err
	}
	user = User{
		ID:        int(userID),
		Name:      name.String,
		Email:     email,
		Password:  password,
		Status:    status.Bool,
		LastLogin: lastLogin.Time.Format("2006-01-02"),
	}
	return user, nil
}
