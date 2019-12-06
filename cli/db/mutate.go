package db

import (
	"cheat/cli/exceptions"
	"cheat/cli/utils"
)

// AddCheat : adds a new cheat to the database
func AddCheat(name string, description string, weight int) string {
	statement, err := database.Prepare(
		`INSERT INTO cheat (
      name, description, weight)
      VALUES (?, ?, ?)
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(exceptions.CheatException("Could not prepare database statement to add a new cheat", err))
	}

	_, err = statement.Exec(name, description, weight)
	if err != nil {
		panic(exceptions.CheatException("Could not execute database statement to add a new cheat", err))
	}

	return name
}

// DeleteCheat : deletes a cheat from the database
func DeleteCheat(name string) {
	statement, err := database.Prepare(
		`DELETE FROM cheat
     WHERE name = ?;
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(exceptions.CheatException("Could not prepare database statement to delete <Cheat: "+name+">", err))
	}

	_, err = statement.Exec(name)
	if err != nil {
		panic(exceptions.CheatException("Could not execute database statement to delete <Cheat: "+name+">", err))
	}
}

// RenameCheat : renames cheat's name (PK) in the database
func RenameCheat(name string, newName string) _Cheat {
	statement, err := database.Prepare(
		`UPDATE cheat
		 SET name = ?
     WHERE name = ?;
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(exceptions.CheatException("Could not prepare database statement to rename <Cheat: "+name+"> to \""+newName+"\"", err))
	}

	_, err = statement.Exec(newName, name)
	if err != nil {
		panic(exceptions.CheatException("Could not execute database statement to rename <Cheat: "+name+"> to \""+newName+"\"", err))
	}

	return GetCheatByName(newName, false)
}

// EditCheat : edits cheat's attributes in the database
func EditCheat(name string, description string, weight int) {
	statement, err := database.Prepare(
		`UPDATE cheat
		 SET description = ?, weight = ?
     WHERE name = ?;
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(exceptions.CheatException("Could not prepare database statement to edit <Cheat: "+name+">", err))
	}

	_, err = statement.Exec(description, weight, name)
	if err != nil {
		panic(exceptions.CheatException("Could not execute database statement to edit <Cheat: "+name+">", err))
	}
}
