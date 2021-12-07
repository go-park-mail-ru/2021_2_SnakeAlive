package repository

const (
	AddTripQuery = `INSERT INTO Trips ("title", "description", "origin") VALUES ($1, $2, $3, $4) 
							RETURNING id`
	AddTripUserQuery     = `INSERT INTO TripsUsers ("trip_id", "user_id") VALUES ($1, $2)`
	GetTripUsersQuery    = `SELECT user_id FROM TripsUsers WHERE trip_id = $1`
	GetTripQuery         = `SELECT id, title, description FROM Trips WHERE id = $1`
	GetPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
								FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id 
								WHERE tr.trip_id = $1
								ORDER BY tr.day, tr.order`
	UpdateTripQuery          = `UPDATE Trips SET "title" = $1, "description" = $2, "origin" = $3 WHERE id = $4`
	DeleteTripQuery          = `DELETE FROM Trips WHERE id = $1`
	DeletePlacesForTripQuery = `DELETE FROM TripsPlaces WHERE trip_id = $1`
	AddPlacesForTripQuery    = `INSERT INTO TripsPlaces ("trip_id", "place_id", "day", "order") VALUES ($1, $2, $3, $4)`
	GetTripAuthorQuery       = `SELECT user_id FROM Trips WHERE id = $1`

	AddAlbumQuery = `INSERT INTO Albums ("title", "description", "trip_id", "author", "photos") VALUES ($1, $2, $3, $4, $5) 
								RETURNING id`
	GetAlbumQuery = `SELECT a.id, a.title, a.description, a.trip_id, a.author, a.photos
								FROM Albums AS a
								WHERE a.id = $1`
	UpdateAlbumQuery    = `UPDATE Albums SET "title" = $1, "description" = $2, "photos" = $3 WHERE id = $4`
	DeleteAlbumQuery    = `DELETE FROM Albums WHERE id = $1`
	GetAlbumAuthorQuery = `SELECT author FROM Albums WHERE id = $1`
	GetAlbumPhotosQuery = `SELECT photos FROM Albums WHERE id = &1`
	AddAlbumPhotosQuery = `UPDATE Albums SET "photos" = $1 WHERE id = $2`

	SightsByTripQuery = `SELECT place_id FROM TripsPlaces AS tp WHERE trip_id = $1 ORDER BY day, tp.order`
	TripsByUserQuery  = `SELECT tr.id, tr.description, tr.title FROM Trips AS tr
							JOIN TripsUsers tu on tr.id = tu.trip_id
							WHERE  tu.user_id = $1 GROUP BY tr.id`
	AlbumsByUserQuery    = `SELECT id, title, description, trip_id, author, photos FROM Albums WHERE author = $1`
	GetAlbumsByTripQuery = `SELECT id, title, description, photos FROM Albums WHERE trip_id = $1`
)
