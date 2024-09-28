package models

import (
	"ikct-ed/config"
	"log"
)

func GetImageData(id int64) ([]byte, error) {

	var imageData []byte

	db, err := config.GetDB2()
	if err != nil {
		log.Println("GetImageData: Failed while connecting with database with error: ", err)
		return []byte{}, err
	}
	defer db.Close()

	query := ` SELECT profile_pic FROM student_financial_info WHERE id= $1`

	err = db.QueryRow(query, id).Scan(&imageData)

	if err != nil {
		log.Println("GetImageData: Failed while executing the query with error: ", err)
		return []byte{}, err
	}
	return imageData, nil
}
