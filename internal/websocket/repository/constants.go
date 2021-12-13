package repository

const (
	getTripQuery         = `SELECT id, title, description FROM Trips WHERE id = $1`
	getPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
	FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id 
	WHERE tr.trip_id = $1
	ORDER BY tr.day, tr.order`
	getAlbumsByTripQuery = `SELECT id, title, description, photos FROM Albums WHERE trip_id = $1`
	getTripUsersQuery    = `SELECT user_id FROM TripsUsers WHERE trip_id = $1`
)
