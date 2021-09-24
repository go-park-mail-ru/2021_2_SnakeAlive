package storage

import ent "snakealive/m/entities"

var AuthDB = map[string]ent.User{
	"alex@mail.ru":   {Name: "Александра", Surname: "Волынец", Email: "alex@mail.ru", Password: "password"},
	"nikita@mail.ru": {Name: "Никита", Surname: "Черных", Email: "nikita@mail.ru", Password: "frontend123"},
	"ksenia@mail.ru": {Name: "Ксения", Surname: "Самойлова", Email: "ksenia@mail.ru", Password: "12345678"},
	"andrew@mail.ru": {Name: "Андрей", Surname: "Кравцов", Email: "andrew@mail.ru", Password: "000111000"},
}

var PlacesDB = map[string][]ent.Place{
	"Russia": {
		{
			Name: "Собор Василия Блаженного", Tags: []string{"Церкви и соборы"}, Photos: []string{"https://i.ibb.co/QmfGC6t/flsdkfj.jpg", "https://i.ibb.co/BCrP43N/photo0jpg.jpg"},
			Author: "Александра Волынец", Review: "Лучшее место для фоточек",
		},
		{
			Name: "Государственный Эрмитаж", Tags: []string{"Музеи исскуств"}, Photos: []string{"https://i.ibb.co/PZvRvCq/es-casi-imposible-tomar.jpg"},
			Author: "Никита Черных", Review: "Вкусно кормят",
		},
		{
			Name: "Сочи Парк", Tags: []string{"Парки равлечений"}, Photos: []string{"https://i.ibb.co/sWJzx0f/caption.jpg", "https://i.ibb.co/CHK7YP6/sdf.jpg"},
			Author: "Ксения Самойлова", Review: "Не хватает каруселей :(",
		},
		{
			Name: "Kazan Kremlin", Tags: []string{"Специализированные музеи", "Исторические достопримечательности"}, Photos: []string{"https://i.ibb.co/JvWrY0J/qol-shariff-mosque.jpg"},
			Author: "Андрей Кравцов", Review: "Я думал это бассеин",
		},
		{
			Name: "Ганина Яма мужской монастырь", Tags: []string{"Культурные объекты и достопримечательности"}, Photos: []string{"https://i.ibb.co/ZKHX76g/dfkjsk.jpg"},
			Author: "Никита Черных", Review: "Спасите меня пожалуйста",
		},
	},
}

var CookieDB = map[string]ent.User{}
