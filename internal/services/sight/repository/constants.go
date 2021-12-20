package repository

import "strconv"

const (
	GetSightByIdQuery = `SELECT pl.id, pl.name, ct.translated, pl.rating, array(
							select name
							from Tags t
							where t.id = ANY(tags)
							) as tag_arr, pl.description,
       						pl.photos, pl.lng, pl.lat FROM Places AS pl
							JOIN Countries AS ct
							ON ct.name = pl.country WHERE pl.id = $1`
	GetSightsByCountryQuery = `SELECT id, name, description, array(
								select name
								from Tags t
								where t.id = ANY(tags)
								) as tag_arr, photos, country, avg(rating) over (partition by id) as total_rating, lng, lat FROM Places WHERE country = $1 ORDER BY total_rating DESC LIMIT 10`
	GetSightsByIDs = `SELECT id, name, country, rating, array(
								select name
								from Tags t
								where t.id = ANY(tags)
								) as tag_arr, description, photos, lng, lat FROM Places WHERE id = any($1)`
	SearchSights = `SELECT id, name, country, rating, array(
								select name
								from Tags t
								where t.id = ANY(tags)
								) as tag_arr, description, photos, lng, lat FROM Places WHERE `
	Tags           = "tags && $"
	Offset         = "OFFSET $"
	Limit          = "LIMIT $"
	GetSightsByTag = `SELECT id, name, country, rating, array(
								select name
								from Tags t
								where t.id = ANY(tags)
								) as tag_arr, description, photos, lng, lat FROM Places WHERE $1 = ANY(tags)`
	GetTags = `SELECT id, name FROM tags`
)

func SearchStatement(pos int) string {
	return "(tsq @@ phraseto_tsquery('russian',LOWER($" + strconv.Itoa(pos) + ")) OR LOWER(name) LIKE $" + strconv.Itoa(pos+1) + ")"
}

func Country(pos int) string {
	return "country = any($" + strconv.Itoa(pos) + ")"
}
