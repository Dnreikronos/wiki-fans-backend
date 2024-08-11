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
		err = rows.Scan(&character.ID, &character.Age, &character.Name, &character.Description, &character.Height, &character.Abilities, &character.Team, &character.Coach, &character.Gender, &character.Weight, &character.Image, &character.GIF)
		if err != nil {
			continue
		}
		characters = append(characters, character)
	}
	return
}
