package db

import (
	"cheat/cli/utils"
)

func AddCheat(command string, name string, description string, weight int) string {
	id := utils.GenerateRandomID()

	statement, err := database.Prepare(
		`INSERT INTO cheat (
      id, command, name, description, weight)
      VALUES (?, ?, ?, ?, ?)
    `,
	)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id, command, name, description, weight)
	if err != nil {
		panic(err)
	}

	return id
}

func DeleteCheat(id string) {
	statement, err := database.Prepare(
		`DELETE FROM cheat
     WHERE id = ?;
    `,
	)
	defer statement.Close()
	if err != nil {
		panic(err)
	}

	_, err = statement.Exec(id)
	if err != nil {
		panic(err)
	}
}
