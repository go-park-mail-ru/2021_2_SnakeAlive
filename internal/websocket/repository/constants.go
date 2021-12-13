package repository

const (
	GetTripQuery         = `SELECT id, title, description FROM Trips WHERE id = $1`
	GetPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
	FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id 
	WHERE tr.trip_id = $1
	ORDER BY tr.day, tr.order`
	GetAlbumsByTripQuery = `SELECT id, title, description, photos FROM Albums WHERE trip_id = $1`
	GetTripUsersQuery    = `SELECT user_id FROM TripsUsers WHERE trip_id = $1`
)
