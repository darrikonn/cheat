package db

import (
	"cheat/cli/utils"
)

// AddCheat : adds a new cheat to the database
func AddCheat(command string, name string, description string, weight int) string {
	id := utils.GenerateRandomID()

	statement, err := database.Prepare(
		`INSERT INTO cheat (
      id, command, name, description, weight)
      VALUES (?, ?, ?, ?, ?)
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id, command, name, description, weight)
	if err != nil {
		panic(err)
	}

	return id
}

// DeleteCheat : deletes a cheat from the database
func DeleteCheat(id string) {
	statement, err := database.Prepare(
		`DELETE FROM cheat
     WHERE id = ?;
    `,
	)
	defer utils.Check(statement.Close)
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		panic(err)
	}
}
