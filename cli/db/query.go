package db

import (
	"database/sql"

	"cheat/cli/models"
	"cheat/cli/utils"
)

type cheat = models.Cheat

// GetCheatByID : returns a cheat by id from the database
func GetCheatByID(id string) cheat {
	var cheat cheat
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

// SearchCheats : query cheats from database by regex
func SearchCheats(searchString string) []cheat {
	var cheats []cheat
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
	defer utils.Check(rows.Close)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var cheat cheat
		err = rows.Scan(&cheat.ID, &cheat.Created, &cheat.Command, &cheat.Name, &cheat.Description, &cheat.Weight)
		if err != nil {
			panic(err)
		}

		cheats = append(cheats, cheat)
	}
	return cheats
}

// GetCheats : gets all cheats from database
func GetCheats() []cheat {
	var cheats []cheat
	rows, err := database.Query(`
    SELECT * FROM cheat
    ORDER BY weight DESC;
    `,
	)
	defer utils.Check(rows.Close)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var cheat cheat
		err = rows.Scan(&cheat.ID, &cheat.Created, &cheat.Command, &cheat.Name, &cheat.Description, &cheat.Weight)
		if err != nil {
			panic(err)
		}

		cheats = append(cheats, cheat)
	}
	return cheats
}
