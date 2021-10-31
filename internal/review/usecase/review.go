package reviewUseCase

import (
	"snakealive/m/pkg/domain"
)

func NewReviewUseCase(reviewStorage domain.reviewStorage) domain.ReviewUseCase {
	return reviewUseCase{reviewStorage: reviewStorage}
}

type reviewUseCase struct {
	userStorage domain.ReviewStorage
}

// func (u userUseCase) GetByEmail(key string) (value domain.User, err error) {
// 	return u.userStorage.GetByEmail(key)
// }

// func (u userUseCase) GetById(id int) (value domain.User, err error) {
// 	return u.userStorage.GetById(id)
// }

func (u reviewUseCase) Add(review domain.Review) error {
	return u.userStorage.Add(review)
}

// func (u userUseCase) Update(id int, updatedUser domain.User) error {

// 	user, err := u.GetByEmail(updatedUser.Email)
// 	if err == nil && user.Id != id {
// 		return errors.New("user with this email already exists") // change later
// 	}

// 	return u.userStorage.Update(id, updatedUser)
// }

// func (u userUseCase) Delete(id int) error {
// 	return u.userStorage.Delete(id)
// }

// func (u userUseCase) Validate(user *domain.User) bool {
// 	if !govalidator.IsEmail(user.Email) ||
// 		!govalidator.StringLength(user.Password, cnst.MinPasswordLength, cnst.MaxPasswordLength) ||
// 		!govalidator.MaxStringLength(user.Email, cnst.MaxEmailLength) {
// 		return false
// 	}
// 	return true
// }

// func (u userUseCase) Login(user *domain.User) (int, error) {
// 	foundUser, err := u.GetByEmail(user.Email)
// 	if err != nil {
// 		log.Printf("error while login-GetByEmail")
// 		log.Print(err)
// 		return fasthttp.StatusNotFound, err
// 	}

// 	if foundUser.Password != user.Password {
// 		return fasthttp.StatusBadRequest, err
// 	}

// 	return fasthttp.StatusOK, err
// }

// func (u userUseCase) Registration(user *domain.User) (int, error) {
// 	_, err := govalidator.ValidateStruct(user)
// 	if err != nil {
// 		log.Printf("error while validating user")
// 		return fasthttp.StatusBadRequest, err
// 	}

// 	_, err = u.GetByEmail(user.Email)
// 	if err == nil {
// 		log.Printf("user with this email already exists")
// 		return fasthttp.StatusBadRequest, err
// 	}

// 	err = u.Add(*user)
// 	if err != nil {
// 		log.Printf("error while adding user")
// 		return fasthttp.StatusBadRequest, err
// 	}

// 	return fasthttp.StatusOK, err
// }

// func (u userUseCase) GetProfile(hash string, foundUser domain.User) (int, []byte) {
// 	response := map[string]string{"name": foundUser.Name, "surname": foundUser.Surname}
// 	bytes, err := json.Marshal(response)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 		return fasthttp.StatusOK, []byte("{}")
// 	}
// 	return fasthttp.StatusOK, bytes
// }

// func (u userUseCase) UpdateProfile(updatedUser *domain.User, foundUser domain.User, hash string) (int, []byte) {
// 	_, err := govalidator.ValidateStruct(updatedUser)
// 	if err != nil {
// 		log.Printf("error while validating user")
// 		return fasthttp.StatusBadRequest, nil
// 	}

// 	if err = u.Update(foundUser.Id, *updatedUser); err != nil {
// 		log.Printf("user with this email already exists")
// 		return fasthttp.StatusBadRequest, nil
// 	}

// 	response := map[string]string{"name": updatedUser.Name, "surname": updatedUser.Surname, "email": updatedUser.Email}
// 	bytes, err := json.Marshal(response)
// 	if err != nil {
// 		log.Printf("error while marshalling JSON: %s", err)
// 		return fasthttp.StatusOK, []byte("{}")
// 	}

// 	return fasthttp.StatusOK, bytes
// }

// func (u userUseCase) DeleteProfile(hash string, foundUser domain.User) int {
// 	u.Delete(foundUser.Id)
// 	return fasthttp.StatusOK
// }

// func (u userUseCase) DeleteUserByEmail(user domain.User) int {
// 	u.DeleteUserByEmail(user)
// 	return fasthttp.StatusOK
// }
