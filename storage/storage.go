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
			Name: "Собор Василия Блаженного", Tags: []string{"Церкви и соборы"}, Photos: []string{""},
			Author: "Александра Волынец", Review: "Лучшее место для фоточек",
		},
		{
			Name: "Государственный Эрмитаж", Tags: []string{"Музеи исскуств"}, Photos: []string{""},
			Author: "Никита Черных", Review: "Вкусно кормят",
		},
		{
			Name: "Сочи Парк", Tags: []string{"Парки равлечений"}, Photos: []string{""},
			Author: "Ксения Самойлова", Review: "Не хватает каруселей :(",
		},
		{
			Name: "Kazan Kremlin", Tags: []string{"Специализированные музеи", "Исторические достопримечательности"}, Photos: []string{""},
			Author: "Андрей Кравцов", Review: "Я думал это бассеин",
		},
		{
			Name: "Ганина Яма мужской монастырь", Tags: []string{"Культурные объекты и достопримечательности"}, Photos: []string{""},
			Author: "Никита Черных", Review: "Спасите меня пожалуйста",
		},
	},
}

var CookieDB = map[string]ent.User{}
