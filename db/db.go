package db

import (
	"database/sql"
	"fmt"
	"log"
	"phone/utils"
)

type DB struct {
	db *sql.DB
}

type phone struct {
	Phone_number string
	Id           int
}

func Open(driverName, dataSource string) (*DB, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Seed(tableName string) error {
	var phones = []string{
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	var normalizedPhones []string

	for _, phone := range phones {
		normalizedPhones = append(normalizedPhones, utils.Normalize(phone))
	}

	normalizedPhones = utils.RemoveDuplicatePhoneNumbers(normalizedPhones)

	for _, normalizedPhone := range normalizedPhones {
		if id, err := insertPhone(db, tableName, normalizedPhone); err != nil {
			return err
		} else {
			fmt.Printf(
				"inserted id=%d %s into table %s\n",
				id,
				normalizedPhone,
				tableName)
		}
	}
	return nil
}

func UpdatePhone(db *DB, tableName string, phone phone) error {
	statement := fmt.Sprintf(`UPDATE %s  SET VALUE($2) WHERE id=$1`, tableName)
	_, err := db.db.Exec(statement, phone.Id, phone.Phone_number)
	return err
}

func DeletePhone(db *DB, tableName string, id int) error {
	statement := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, tableName)
	_, err := db.db.Exec(statement, id)
	return err
}

func GetPhoneById(db *DB, tableName string, id int) (string, error) {
	statement := fmt.Sprintf(
		`SELECT * FROM %s WHERE id=$1`, tableName)
	var phone_number string
	err := db.db.QueryRow(statement, id).Scan(&id, &phone_number)
	if err != nil {
		return "", err
	}
	return phone_number, nil
}

func FindPhone(db *DB, tableName string, phoneNumber string) (*phone, error) {
	var p phone
	statement := fmt.Sprintf(
		`SELECT * FROM %s WHERE phone_number=$1`, tableName)
	err := db.db.QueryRow(statement, phoneNumber).Scan(&p.Id, &p.Phone_number)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetPhones(db *DB, tableName string) ([]phone, error) {
	statement := fmt.Sprintf(
		`SELECT id , phone_number FROM %s`, tableName)
	rows, err := db.db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []phone
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.Id, &p.Phone_number); err != nil {
			log.Fatal(err)
		}
		ret = append(ret, p)
	}
	return ret, nil
}

func insertPhone(db *DB, tableName, phone string) (int, error) {
	statement := fmt.Sprintf(
		`INSERT INTO %s (phone_number) VALUES($1) RETURNING id`, tableName)
	var id int
	err := db.db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ResetDB(driverName, dataSource, dbName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = resetDB(db, dbName)
	if err != nil {
		return err
	}
	return db.Close()
}

func resetDB(db *sql.DB, name string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + name)
	if err != nil {
		return err
	}
	return createDB(db, name)
}

func createDB(db *sql.DB, name string) error {
	_, err := db.Exec("CREATE DATABASE " + name)
	if err != nil {
		return err
	}
	return nil
}

func Migrate(driverName, dataSource, dbName, tableName string) error {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	err = createPhoneNumberTable(db, tableName)
	if err != nil {
		return err
	}
	return db.Close()
}

func createPhoneNumberTable(db *sql.DB, tableName string) error {
	statement := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s(
			id SERIAL,
			phone_number VARCHAR(255)
		) `, tableName)
	_, err := db.Exec(statement)
	return err

}
