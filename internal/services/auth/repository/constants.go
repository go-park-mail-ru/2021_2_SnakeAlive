package repository

const (
	getUserByEmailRequest = "SELECT id, password FROM Users WHERE email = $1"
	getUserByIDRequest    = "SELECT name, surname, email, avatar, description FROM Users WHERE id = $1"
	createUserRequest     = "INSERT INTO Users (name, surname, email, password) VALUES ($1,$2,$3,$4) RETURNING id"
	updateUserRequest     = "UPDATE Users SET "
	updateUserName        = "name=$"
	updateUserSurname     = "surname=$"
	updateUserPass        = "password=$"
	updateUserEmail       = "email=$"
	updateUserDescription = "description=$"
	updateUserAvatar      = "avatar=$"
	whereStatement        = "WHERE id = $1"
	updateUserReturning   = "RETURNING name, surname, email, avatar, description"
	createUserSession     = "INSERT INTO Cookies (user_id, hash) VALUES ($1, $2)"
	validateUserSession   = "SELECT user_id FROM Cookies WHERE hash = $1"
	removeUserSession     = "DELETE FROM Cookies WHERE hash = $1"
	GetUserInfoQuery      = "SELECT id, name, surname, avatar FROM Users WHERE id = $1"
)
