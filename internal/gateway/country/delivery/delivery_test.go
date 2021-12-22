package delivery

import (
	mock_repository "snakealive/m/internal/gateway/country/repository/mock"
)

type Test struct {
	Prepare func(repo *mock_repository.MockCountryStorageMockRecorder)
	Run     func(d CountryDelivery)
}

//func prepare(t *testing.T) (d CountryDelivery, repo *mock_repository.MockCountryStorageMockRecorder) {
//	ctrl := gomock.NewController(t)
//	repo = mock_repository.NewMockCountryStorage(ctrl)
//
//	return NewCountryDelivery(nil, nil)
//}
