package models

import (
	"github.com/Dnreikronos/wiki-fans-backend/db"
)

func Insert(characters Character) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO character (name, description, image, height, abilities, age, team, coach, gender, weight) VALUES ($1, $2, $3) RETURN id`

	err = conn.QueryRow(sql, characters.Name, characters.Description, characters.Image, characters.Height,
		characters.Abilities, characters.Age, characters.Team, characters.Coach,
		characters.Gender, characters.Weight).Scan(&id)

	return
}
