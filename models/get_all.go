package models

import "github.com/Dnreikronos/wiki-fans-backend/db"

func GetAll() (characters []Character, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM character`)
	if err != nil {
		return
	}
	for rows.Next() {
		var character Character
		err = rows.Scan(&character.Name, &character.Description, &character.Image, &character.Height, &character.Abilities, &character.Age, &character.Team, &character.Coach, &character.Gender, &character.Weight)
	}
	return
}
