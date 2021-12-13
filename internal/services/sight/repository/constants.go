package repository

const (
	GetSightByIdQuery       = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE id = $1`
	GetSightsByCountryQuery = `SELECT id, name, description, tags, photos, country, avg(rating) over (partition by id) as total_rating, lng, lat FROM Places WHERE country = $1 ORDER BY total_rating DESC LIMIT 10`
	GetSightsByIDs          = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE id = any($1)`
	SearchSights            = `SELECT id, name, country, rating, tags, description, photos, lng, lat FROM Places WHERE name LIKE $1 OFFSET $2 LIMIT $3`
	GetSightsByTag          = `SELECT id, name, country, rating, array(
								select name
								from Tags t
								where t.id = ANY(tags)
								) as tag_arr, description, photos, lng, lat, FROM Places WHERE $1 = ANY(tags)`
	GetTags = `SELECT id, name FROM tags`
)
