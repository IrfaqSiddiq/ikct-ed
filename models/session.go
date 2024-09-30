package models

import (
	"ikct-ed/config"
	"log"
)

// CheckSession function to check for login history
func CheckSession(token string) (bool, error) {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("CheckSession: Failed while connecting with the database :", err)
		return false, err
	}
	defer db.Close()

	selectQuery := `SELECT 
				is_expire,
				user_id
				FROM
				session
				WHERE
				user_token = $1`
	var flag bool
	var userID string
	err = db.QueryRow(selectQuery, token).Scan(&flag, &userID)
	if err != nil {
		log.Println("CheckSession: failed:", err)
		return false, err
	}
	return !flag, err
}
