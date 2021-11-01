package domain

//go:generate mockgen -source=review.go -destination=/mocks/mock.go
type Reviews []Review

type Review struct {
	Id         int    `json:"-"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Rating     int    `json:"rating"`
	User_id    int    `json:"user_id"`
	Place_id   int    `json:"place_id"`
	Created_at string `json:"created_at"`
}

type ReviewStorage interface {
	Add(value Review) error
	Get(id int) (Review, error)
	GetListByPlace(id int) (Reviews, error)
	Delete(id int) error
}

type ReviewUseCase interface {
	Add(review Review) (int, error)
	Get(id int) (Review, error)
	Delete(id int) error
	GetReviewsListByPlaceId(id int) (int, []byte)
}
