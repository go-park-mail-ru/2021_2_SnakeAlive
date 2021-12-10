package repository

const AddReviewQuery = `INSERT INTO public.reviews (title, text, rating, user_id, place_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`

const GetReviewByIdQuery = `SELECT id, title, text, rating, user_id, place_id FROM Reviews WHERE id = $1`

const DeleteReviewQuery = `DELETE FROM Reviews WHERE id = $1`

const GetReviewAuthorQuery = `SELECT user_id FROM Reviews WHERE id = $1`

const GetReviewsByPlaceQuery = `SELECT id, title, text, rating, user_id FROM Reviews WHERE place_id = $1 LIMIT $2 OFFSET $3`
