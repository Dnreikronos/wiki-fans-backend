package models

import "github.com/Dnreikronos/wiki-fans-backend/db"

func Get(id int64) (character Character, err error) {

	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM character WHERE id=$1`, id)

	err = row.Scan(&character.Name, &character.Description, &character.Image, &character.Height, &character.Abilities, &character.Age, &character.Team, &character.Coach, &character.Gender, &character.Weight)

	return
}
