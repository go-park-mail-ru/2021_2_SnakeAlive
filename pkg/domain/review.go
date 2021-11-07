package domain

type Reviews []Review

type Review struct {
	Id        int    `json:"-"`
	Title     string `json:"title"`
	Text      string `json:"text"`
	Rating    int    `json:"rating"`
	UserId    int    `json:"user_id"`
	PlaceId   int    `json:"place_id"`
	CreatedAt string `json:"created_at"`
}

type ReviewStorage interface {
	Add(value Review, userId int) (int, error)
	Get(id int) (Review, error)
	GetListByPlace(id int) (Reviews, error)
	Delete(id int) error
	GetReviewAuthor(id int) int
}

type ReviewUseCase interface {
	Add(review Review, user User) (int, []byte, error)
	Get(id int) (Review, error)
	Delete(id int) error
	GetReviewsListByPlaceId(id int) (int, []byte)
	CheckAuthor(user User, id int) bool
	SanitizeReview(review Review) Review
}
