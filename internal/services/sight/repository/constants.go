package repository

const (
	GetSightByIdQuery       = `SELECT id, name, country, rating, tags, description, photos FROM Places WHERE id = $1`
	GetSightsByCountryQuery = `SELECT id, name, tags, photos, country, avg(rating) over (partition by id) as total_rating FORM Places WHERE country = $1 ORDER BY total_rating DESC LIMIT 10`
)
