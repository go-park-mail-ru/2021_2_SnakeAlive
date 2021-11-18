package repository

const (
	getUserByEmailRequest = "SELECT id, password FROM Users WHERE email = ?;"
	getUserByIDRequest    = "SELECT name, surname, email, avatar, description FROM Users WHERE id =?;"
	createUserRequest     = "INSERT INTO Users (name, surname, email, password) VALUES (?,?,?,?) RETURNING id;"
	updateUserRequest     = "UPDATE Users WHERE id = ? SET "
	updateUserName        = "name=?"
	updateUserSurname     = "surname=?"
	updateUserPass        = "pass=?"
	updateUserEmail       = "email=?"
	updateUserDescription = "description=?"
	updateUserReturning   = "RETURNING name, surname, email, avatar, description;"
	createUserSession     = "INSERT INTO Cookies (user_id, hash) VALUES (?, ?);"
	validateUserSession   = "SELECT user_id FROM Cookies WHERE hash = ?;"
	removeUserSession     = "DELETE FROM Cookies WHERE hash = ?;"
)
