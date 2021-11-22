package constants

const AddCookieQuery = `INSERT INTO Cookies ("hash", "user_id") VALUES ($1, $2)`
const GetCookieQuery = `SELECT U.id, U.name, U.surname, U.password, U.email, U.description, U.avatar
						FROM Users AS U JOIN Cookies AS C ON U.id = C.user_id
						WHERE C.hash = $1`
const DeleteCookieQuery = `DELETE FROM Cookies WHERE hash = $1`

const GetPlaceByIdQuery = `SELECT id, name, country, lat, lng, rating, tags, description, photos FROM Places WHERE id = $1`

const AddReviewQuery = `INSERT INTO public.reviews (title, text, rating, user_id, place_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

const GetReviewByIdQuery = `SELECT id, title, text, rating, user_id, place_id FROM Reviews WHERE id = $1`
const DeleteReviewQuery = `DELETE FROM Reviews WHERE id = $1`
const GetReviewAuthorQuery = `SELECT user_id FROM Reviews WHERE id = $1`

const AddTripQuery = `INSERT INTO Trips ("title", "description", "days", "user_id", "origin") 
						VALUES ($1, $2, $3, $4, $5) RETURNING id`
const GetTripQuery = `SELECT id, title, description FROM Trips WHERE id = $1`
const GetPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
								FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id WHERE tr.trip_id = $1
								ORDER BY tr.day, tr.order`
const GetTripPlaceCoord = `SELECT pl.id, pl.lat, pl.lng
FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id WHERE tr.trip_id = $1
ORDER BY  tr.order`

const UpdateTripQuery = `UPDATE Trips SET "title" = $1, "description" = $2, "days" = $3, "origin" = $4 WHERE id = $5`
const DeleteTripQuery = `DELETE FROM Trips WHERE id = $1`
const DeletePlacesForTripQuery = `DELETE FROM TripsPlaces WHERE trip_id = $1`
const AddPlacesForTripQuery = `INSERT INTO TripsPlaces ("trip_id", "place_id", "day", "order") VALUES ($1, $2, $3, $4)`
const GetTripAuthorQuery = `SELECT user_id FROM Trips WHERE id = $1`

const AddUserQuery = `INSERT INTO Users ("name", "surname", "password", "email", "description") VALUES ($1, $2, $3, $4, $5)`
const DeleteUserByIdQuery = `DELETE FROM Users WHERE id = $1`
const DeleteUserByEmailQuery = `DELETE FROM Users WHERE email = $1`
const GetUserByEmailQuery = `SELECT id, name, surname, password, email, avatar, description FROM Users WHERE email = $1`
const GetUserByIdQuery = `SELECT id, name, surname, password, email, avatar, description FROM Users WHERE id = $1`
const UpdateUserQuery = `UPDATE Users SET "name" = $1, "surname" = $2, "email" = $3, "password" = $4,
	"description" = $5 WHERE id = $6`
const AddAvatarQuery = `UPDATE Users SET "avatar" = $1 WHERE id = $2`
