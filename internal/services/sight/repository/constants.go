package repository

const (
	GetSightByIdQuery       = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE id = $1`
	GetSightsByCountryQuery = `SELECT id, name, tags, photos, country, avg(rating) over (partition by id) as total_rating, lng, lat FROM Places WHERE country = $1 ORDER BY total_rating DESC LIMIT 10`
	GetSightsByIDs          = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE id = any($1)`
	SearchSights            = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE name LIKE $1 OFFSET $2 LIMIT $3`
)
