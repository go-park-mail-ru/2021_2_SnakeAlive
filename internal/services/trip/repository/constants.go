package repository

const (
	AddTripQuery = `INSERT INTO Trips ("title", "description", "days", "user_id", "origin") VALUES ($1, $2, $3, $4, $5) 
							RETURNING id`
	GetTripQuery         = `SELECT id, title, description FROM Trips WHERE id = $1`
	GetPlaceForTripQuery = `SELECT pl.id, pl.name, pl.tags, pl.description, pl.rating, pl.country, pl.photos, tr.day
								FROM TripsPlaces AS tr JOIN Places AS pl ON tr.place_id = pl.id WHERE tr.trip_id = $1
								ORDER BY tr.day, tr.order`
	UpdateTripQuery          = `UPDATE Trips SET "title" = $1, "description" = $2, "days" = $3, "origin" = $4 WHERE id = $5`
	DeleteTripQuery          = `DELETE FROM Trips WHERE id = $1`
	DeletePlacesForTripQuery = `DELETE FROM TripsPlaces WHERE trip_id = $1`
	AddPlacesForTripQuery    = `INSERT INTO TripsPlaces ("trip_id", "place_id", "day", "order") VALUES ($1, $2, $3, $4)`
	GetTripAuthorQuery       = `SELECT user_id FROM Trips WHERE id = $1`
)
