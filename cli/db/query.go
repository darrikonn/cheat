package db

import (
	"database/sql"

	"cheat/cli/exceptions"
	"cheat/cli/utils"
)

// GetCheatByName : returns a cheat by name from the database
func GetCheatByName(name string, ignoreCase bool) _Cheat {
	var cheat _Cheat
	row := database.QueryRow(`
    SELECT name, created, description, weight FROM cheat
    WHERE regexp($1, name, $2)
    ORDER BY name ASC;
    `,
		name, ignoreCase,
	)
	err := row.Scan(&cheat.Name, &cheat.Created, &cheat.Description, &cheat.Weight)

	switch err {
	case sql.ErrNoRows:
		panic(exceptions.CheatException("<Cheat: "+name+"> could not be found!", err))
	case nil:
		return cheat
	default:
		panic(exceptions.CheatException("Unknown exception occurred", err))
	}
}

// SearchCheats : query cheats from database by regex
func SearchCheats(searchString string, ignoreCase bool) []_Cheat {
	var cheats []_Cheat
	rows, err := database.Query(`
    SELECT name, created, description, weight FROM cheat
    WHERE
      regexp($1, name, $2)
      OR regexp($1, description, $2)
    ORDER BY
      weight DESC,
      description;
    `,
		searchString, ignoreCase,
	)
	defer utils.Check(rows.Close, "Could not close rows returned by regex search query")
	if err != nil {
		panic(exceptions.CheatException("Could not query database from regex: \""+searchString+"\"", err))
	}

	for rows.Next() {
		var cheat _Cheat
		err = rows.Scan(&cheat.Name, &cheat.Created, &cheat.Description, &cheat.Weight)
		if err != nil {
			panic(exceptions.CheatException("Could not scan database row from query", err))
		}

		cheats = append(cheats, cheat)
	}
	return cheats
}
