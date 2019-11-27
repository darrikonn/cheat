package db

import (
	"database/sql"

	"cheat/cli/models"
)

type Cheat = models.Cheat

func GetCheatById(id string) Cheat {
	var cheat Cheat
	row := database.QueryRow(`
    SELECT * FROM cheat
    WHERE id LIKE ('%' || $1 || '%')
    ORDER BY weight;
    `,
		id,
	)
	err := row.Scan(&cheat.ID, &cheat.Created, &cheat.Command, &cheat.Name, &cheat.Description, &cheat.Weight)

	switch err {
	case sql.ErrNoRows:
		panic(err)
	case nil:
		return cheat
	default:
		panic(err)
	}
}

func SearchCheats(searchString string) []Cheat {
	var cheats []Cheat
	rows, err := database.Query(`
    SELECT * FROM cheat
    WHERE
      command regexp $1
      OR name regexp $1
      OR description regexp $1
    ORDER BY weight DESC;
    `,
		searchString,
	)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var cheat Cheat
		err = rows.Scan(&cheat.ID, &cheat.Created, &cheat.Command, &cheat.Name, &cheat.Description, &cheat.Weight)
		if err != nil {
			panic(err)
		}

		cheats = append(cheats, cheat)
	}
	return cheats
}

func GetCheats() []Cheat {
	var cheats []Cheat
	rows, err := database.Query(`
    SELECT * FROM cheat
    ORDER BY weight DESC;
    `,
	)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var cheat Cheat
		err = rows.Scan(&cheat.ID, &cheat.Created, &cheat.Command, &cheat.Name, &cheat.Description, &cheat.Weight)
		if err != nil {
			panic(err)
		}

		cheats = append(cheats, cheat)
	}
	return cheats
}
