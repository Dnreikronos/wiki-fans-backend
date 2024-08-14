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

	sql := `INSERT INTO character (id, age, name, description, height, abilities, team, coach, gender, weight, image) VALUES ($1, $2, $3) RETURN id`

	err = conn.QueryRow(sql, characters.Age, characters.Name, characters.Description,
		characters.Height, characters.Abilities, characters.Team, characters.Coach,
		characters.Gender, characters.Weight, characters.Image, characters.GIF).Scan(&id)

	return
}
