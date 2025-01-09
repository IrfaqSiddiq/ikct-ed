package models

import (
	"ikct-ed/config"
	"log"
)

func AddSchool(school string) error {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("AddSchool: Failed while connecting with database with error: ", err)
		return err
	}
	defer db.Close()
	query := ` INSERT INTO schools(name)VALUES($1)`

	_, err = db.Exec(query, school)
	if err != nil {
		log.Println("AddSchool: Failed while executing the query with error: ", err)
		return err
	}
	
	return nil
}
func DeleteSchool(schoolID int64) error {
	db, err := config.GetDB2()
	if err != nil {
		log.Println("DeleteSchool: Failed while connecting with database with error: ", err)
		return err
	}
	defer db.Close()
	query := ` DELETE from schools WHERE id = $1`
	_,err = db.Exec(query, schoolID)
	if err!= nil {
        log.Println("DeleteSchool: Failed while executing the query with error: ", err)
        return err
    }
	return nil
}
