package models

// Cheat : mapping of the cheat entity stored in the database
type Cheat struct {
	ID          string
	Created     string
	Command     string
	Name        string
	Description string
	Weight      int
}
