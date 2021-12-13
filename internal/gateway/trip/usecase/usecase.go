package usecase

import (
	"context"

	"snakealive/m/internal/models"
	"snakealive/m/pkg/errors"
	auth_service "snakealive/m/pkg/services/auth"
	sight_service "snakealive/m/pkg/services/sight"
	trip_service "snakealive/m/pkg/services/trip"
)

type TripGatewayUseCase interface {
	AddTrip(ctx context.Context, value *models.Trip, userID int) (*models.TripWithUserInfo, error)
	GetTripById(ctx context.Context, id int, userID int) (*models.TripWithUserInfo, error)
	DeleteTrip(ctx context.Context, id int, userID int) error
	UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.TripWithUserInfo, error)
	GetTripsByUser(ctx context.Context, id int) (*[]models.TripWithUserInfo, error)

	AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error)
	GetAlbumById(ctx context.Context, id int, userID int) (*models.Album, error)
	DeleteAlbum(ctx context.Context, id int, userID int) error
	UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album, userID int) (*models.Album, error)
	SightsByTrip(ctx context.Context, id int) (*[]models.TripSight, error)
	GetAlbumsByUser(ctx context.Context, id int) (*[]models.Album, error)

	AddTripUser(ctx context.Context, author int, tripId int, email string) error

	ShareLink(ctx context.Context, author int, tripId int) (string, error)
	AddUserByLink(ctx context.Context, author int, tripId int, uuid string) (string, error)

	GetUserInfo(ctx context.Context, ids []int64) (*[]models.UserInfo, error)
}

type tripGatewayUseCase struct {
	tripGRPC  tripGRPC
	sightGRPC sightGRPC
	authGRPC  authGRPC
}

func NewTripGatewayUseCase(grpc tripGRPC, sightGRPC sightGRPC, authGRPC authGRPC) TripGatewayUseCase {
	return &tripGatewayUseCase{
		tripGRPC:  grpc,
		sightGRPC: sightGRPC,
		authGRPC:  authGRPC,
	}
}

func (u *tripGatewayUseCase) AddTrip(ctx context.Context, value *models.Trip, userID int) (*models.TripWithUserInfo, error) {
	days := ProtoDaysFromPlaces(value.Sights)
	trip := &trip_service.Trip{
		Id:          int64(value.Id),
		Title:       value.Title,
		Description: value.Description,
		Sights:      days,
	}

	responce, err := u.tripGRPC.AddTrip(ctx,
		&trip_service.ModifyTripRequest{
			Trip:   trip,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	places := PlacesFromProtoDays(responce.Sights)
	albums := AlbumsFromProtoAlbums(responce.Albums)
	users, err := u.GetUserInfo(ctx, responce.Users)
	if err != nil {
		return nil, err
	}

	return &models.TripWithUserInfo{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
		Users:       *users,
	}, nil
}

func (u *tripGatewayUseCase) GetTripById(ctx context.Context, tripId int, userID int) (*models.TripWithUserInfo, error) {
	responce, err := u.tripGRPC.GetTrip(ctx, &trip_service.TripRequest{
		TripId: int64(tripId),
		UserId: int64(userID),
	})
	if err != nil {
		return nil, err
	}

	if responce.Id == 0 {
		return nil, errors.TripNotFound
	}

	places := PlacesFromProtoDays(responce.Sights)
	albums := AlbumsFromProtoAlbums(responce.Albums)
	users, err := u.GetUserInfo(ctx, responce.Users)
	if err != nil {
		return nil, err
	}

	return &models.TripWithUserInfo{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
		Users:       *users,
	}, nil
}

func (u *tripGatewayUseCase) DeleteTrip(ctx context.Context, id int, userID int) error {
	_, err := u.tripGRPC.DeleteTrip(ctx, &trip_service.TripRequest{
		TripId: int64(id),
		UserId: int64(userID),
	})
	return err
}

func (u *tripGatewayUseCase) UpdateTrip(ctx context.Context, id int, updatedTrip *models.Trip, userID int) (*models.TripWithUserInfo, error) {
	days := ProtoDaysFromPlaces(updatedTrip.Sights)
	trip := &trip_service.Trip{
		Id:          int64(id),
		Title:       updatedTrip.Title,
		Description: updatedTrip.Description,
		Sights:      days,
	}

	responce, err := u.tripGRPC.UpdateTrip(ctx,
		&trip_service.ModifyTripRequest{
			Trip:   trip,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	places := PlacesFromProtoDays(responce.Sights)
	albums := AlbumsFromProtoAlbums(responce.Albums)
	users, err := u.GetUserInfo(ctx, responce.Users)
	if err != nil {
		return nil, err
	}

	return &models.TripWithUserInfo{
		Id:          int(responce.Id),
		Title:       responce.Title,
		Description: responce.Description,
		Sights:      places,
		Albums:      albums,
		Users:       *users,
	}, nil
}

func (u *tripGatewayUseCase) AddAlbum(ctx context.Context, album *models.Album, userID int) (*models.Album, error) {
	protoAlbum := &trip_service.Album{
		Id:          int64(album.Id),
		Title:       album.Title,
		TripId:      int64(album.TripId),
		Description: album.Description,
		Photos:      album.Photos,
	}

	responce, err := u.tripGRPC.AddAlbum(ctx,
		&trip_service.ModifyAlbumRequest{
			Album:  protoAlbum,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
	}, nil
}

func (u *tripGatewayUseCase) GetAlbumById(ctx context.Context, id int, userID int) (*models.Album, error) {
	responce, err := u.tripGRPC.GetAlbum(ctx, &trip_service.AlbumRequest{
		AlbumId: int64(id),
		UserId:  int64(userID),
	})
	if err != nil {
		return nil, err
	}

	if responce.Id == 0 {
		return nil, errors.TripNotFound
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
		Photos:      responce.Photos,
	}, nil
}

func (u *tripGatewayUseCase) DeleteAlbum(ctx context.Context, id int, userID int) error {
	_, err := u.tripGRPC.DeleteAlbum(ctx, &trip_service.AlbumRequest{
		AlbumId: int64(id),
		UserId:  int64(userID),
	})
	return err
}

func (u *tripGatewayUseCase) UpdateAlbum(ctx context.Context, id int, updatedAlbum *models.Album, userID int) (*models.Album, error) {
	album := &trip_service.Album{
		Id:          int64(id),
		Title:       updatedAlbum.Title,
		Description: updatedAlbum.Description,
		Photos:      updatedAlbum.Photos,
	}

	responce, err := u.tripGRPC.UpdateAlbum(ctx,
		&trip_service.ModifyAlbumRequest{
			Album:  album,
			UserId: int64(userID),
		},
	)
	if err != nil {
		return nil, err
	}

	return &models.Album{
		Id:          int(responce.Id),
		Title:       responce.Title,
		TripId:      int(responce.TripId),
		UserId:      int(responce.Author),
		Description: responce.Description,
		Photos:      responce.Photos,
	}, nil
}

func (u *tripGatewayUseCase) SightsByTrip(ctx context.Context, id int) (*[]models.TripSight, error) {
	ids, err := u.tripGRPC.SightsByTrip(ctx, &trip_service.SightsRequest{
		TripId: int64(id),
	})
	if err != nil {
		return nil, err
	}

	sights, err := u.sightGRPC.GetSightsByIDs(ctx, &sight_service.GetSightsByIDsRequest{Ids: ids.Ids})
	if err != nil {
		return nil, err
	}

	adapted := make([]models.TripSight, len(sights.Sights))
	for i, sight := range sights.Sights {
		adapted[i] = models.TripSight{
			Id:  int(sight.Id),
			Lng: sight.Lng,
			Lat: sight.Lat,
		}
	}

	return &adapted, nil
}

func (u *tripGatewayUseCase) GetTripsByUser(ctx context.Context, id int) (*[]models.TripWithUserInfo, error) {
	protoTrips, err := u.tripGRPC.GetTripsByUser(ctx, &trip_service.ByUserRequest{UserId: int64(id)})
	if err != nil {
		return nil, err
	}

	var trips []models.TripWithUserInfo
	for _, trip := range protoTrips.Trips {
		places := PlacesFromProtoDays(trip.Sights)
		albums := AlbumsFromProtoAlbums(trip.Albums)

		users, err := u.GetUserInfo(ctx, trip.Users)
		if err != nil {
			return nil, err
		}

		trips = append(trips, models.TripWithUserInfo{
			Id:          int(trip.Id),
			Title:       trip.Title,
			Description: trip.Description,
			Albums:      albums,
			Sights:      places,
			Users:       *users,
		})
	}

	return &trips, nil
}

func (u *tripGatewayUseCase) GetAlbumsByUser(ctx context.Context, id int) (*[]models.Album, error) {
	protoAlbums, err := u.tripGRPC.GetAlbumsByUser(ctx, &trip_service.ByUserRequest{UserId: int64(id)})
	if err != nil {
		return nil, err
	}

	var albums []models.Album
	for _, album := range protoAlbums.Albums {
		albums = append(albums, models.Album{
			Id:          int(album.Id),
			Title:       album.Title,
			TripId:      int(album.TripId),
			UserId:      int(album.Author),
			Description: album.Description,
			Photos:      album.Photos,
		})
	}
	return &albums, nil
}

func (u *tripGatewayUseCase) AddTripUser(ctx context.Context, author int, tripId int, email string) error {
	user, err := u.authGRPC.GetUserByEmail(ctx, &auth_service.UserEmailRequest{Email: email})
	if err != nil {
		return err
	}

	_, err = u.tripGRPC.AddTripUser(ctx, &trip_service.AddTripUserRequest{
		TripId: int64(tripId),
		UserId: int64(user.Id),
		Author: int64(author),
	})
	return err
}

func (u *tripGatewayUseCase) ShareLink(ctx context.Context, author int, tripId int) (string, error) {
	link, err := u.tripGRPC.ShareLink(ctx, &trip_service.ShareRequest{
		TripId: int64(tripId),
		UserId: int64(author),
	})
	if err != nil {
		return "", err
	}

	return link.Link, nil
}

func (u *tripGatewayUseCase) AddUserByLink(ctx context.Context, author int, tripId int, uuid string) (string, error) {

	responce, err := u.tripGRPC.AddUserByLink(ctx,
		&trip_service.AddByShareRequest{
			TripId: int64(tripId),
			UserId: int64(author),
			Uuid:   uuid,
		},
	)
	if err != nil {
		return "", err
	}

	return responce.Link, nil
}

func (u *tripGatewayUseCase) GetUserInfo(ctx context.Context, ids []int64) (*[]models.UserInfo, error) {
	users := make([]models.UserInfo, 0)
	for _, id := range ids {
		user, err := u.authGRPC.GetUserInfo(ctx, &auth_service.GetUserRequest{Id: id})
		if err != nil {
			return nil, err
		}

		users = append(users, models.UserInfo{
			Id:      int(user.UserId),
			Name:    user.Name,
			Surname: user.Surname,
			Avatar:  user.Image,
		})
	}
	return &users, nil
}

func ProtoDaysFromPlaces(places []models.Place) []*trip_service.Sight {
	var protoDays []*trip_service.Sight
	for _, sight := range places {
		protoSight := trip_service.Sight{
			Id:          int64(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        sight.Tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int64(sight.Day),
		}
		protoDays = append(protoDays, &protoSight)
	}
	return protoDays
}

func PlacesFromProtoDays(sights []*trip_service.Sight) []models.Place {
	var places []models.Place
	for _, sight := range sights {
		placesSight := models.Place{
			Id:          int(sight.Id),
			Name:        sight.Name,
			Country:     sight.Country,
			Rating:      sight.Rating,
			Tags:        sight.Tags,
			Description: sight.Description,
			Photos:      sight.Photos,
			Day:         int(sight.Day),
		}
		places = append(places, placesSight)
	}
	return places
}

func ProtoAlbumsFromAlbums(albums []models.Album) []*trip_service.Album {
	var protoAlbums []*trip_service.Album
	for _, album := range albums {
		protoAlbum := trip_service.Album{
			Id:          int64(album.Id),
			Title:       album.Title,
			Description: album.Description,
			Photos:      album.Photos,
		}
		protoAlbums = append(protoAlbums, &protoAlbum)
	}
	return protoAlbums
}

func AlbumsFromProtoAlbums(protoAlbums []*trip_service.Album) []models.Album {
	var albums []models.Album
	for _, album := range protoAlbums {
		modelAlbum := models.Album{
			Id:          int(album.Id),
			Title:       album.Title,
			Description: album.Description,
			Photos:      album.Photos,
		}
		albums = append(albums, modelAlbum)
	}
	return albums
}
