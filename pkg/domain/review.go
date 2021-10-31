package domain

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
	Get(key string) (Reviews, error)
	// GetById(id int) (value User, err error)
	// GetByEmail(key string) (value User, err error)
	// Delete(id int) error
	// Update(id int, value User) error
	// DeleteByEmail(user User) error
}

type ReviewUseCase interface {
	Add(review Review) error
	Get(key string) error
	GetReviewsListByName(param string) (int, []byte)
	// GetById(id int) (value User, err error)
	// GetByEmail(key string) (value User, err error)
	// Delete(id int) error
	// Update(id int, updatedUser User) error
	// Validate(user *User) bool
	// Login(user *User) (int, error)
	// Registration(user *User) (int, error)
	// GetProfile(hash string, user User) (int, []byte)
	// UpdateProfile(updatedUser *User, foundUser User, hash string) (int, []byte)
	// DeleteProfile(hash string, foundUser User) int
	// DeleteUserByEmail(user User) int
}
