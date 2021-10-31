package entities

import "snakealive/m/pkg/domain"

var AuthDB = map[string]domain.User{
	"alex@mail.ru":   {Name: "Александра", Surname: "Волынец", Email: "alex@mail.ru", Password: "password"},
	"nikita@mail.ru": {Name: "Никита", Surname: "Черных", Email: "nikita@mail.ru", Password: "frontend123"},
	"ksenia@mail.ru": {Name: "Ксения", Surname: "Самойлова", Email: "ksenia@mail.ru", Password: "12345678"},
	"andrew@mail.ru": {Name: "Андрей", Surname: "Кравцов", Email: "andrew@mail.ru", Password: "000111000"},
}

var PlacesDB = map[string]domain.Places{
	"Russia": {
		{
			Name: "Собор Василия Блаженного", Tags: []string{"Церкви и соборы"},
			Photos: []string{"https://i.ibb.co/QmfGC6t/flsdkfj.jpg", "https://i.ibb.co/BCrP43N/photo0jpg.jpg"},
			Author: "Александра Волынец", Review: "Лучшее место для фоточек",
		},
		{
			Name: "Государственный Эрмитаж", Tags: []string{"Музеи исскуств"},
			Photos: []string{"https://i.ibb.co/PZvRvCq/es-casi-imposible-tomar.jpg"},
			Author: "Никита Черных", Review: "Вкусно кормят",
		},
		{
			Name: "Сочи Парк", Tags: []string{"Парки равлечений"},
			Photos: []string{"https://i.ibb.co/sWJzx0f/caption.jpg", "https://i.ibb.co/CHK7YP6/sdf.jpg"},
			Author: "Ксения Самойлова", Review: "Не хватает каруселей :(",
		},
		{
			Name: "Kazan Kremlin", Tags: []string{"Специализированные музеи", "Исторические достопримечательности"},
			Photos: []string{"https://i.ibb.co/JvWrY0J/qol-shariff-mosque.jpg"},
			Author: "Андрей Кравцов", Review: "Я думал это бассеин",
		},
		{
			Name: "Ганина Яма мужской монастырь", Tags: []string{"Культурные объекты и достопримечательности"},
			Photos: []string{"https://i.ibb.co/ZKHX76g/dfkjsk.jpg"},
			Author: "Никита Черных", Review: "Спасите меня пожалуйста",
		},
	},
	"Nicaragua": {
		{
			Name: "Puerto Salvador Allende", Tags: []string{"Культурные объекты и достопримечательности"},
			Photos: []string{"https://i.ibb.co/dPF3qpQ/puerto-salvador-allende.jpg"},
			Author: "Никита Черных", Review: "Мое самое любимое место во всем Никарагуа!!!! Советую посетить всем, всем!!",
		},
		{
			Name: "Plaza de la Revolucion", Tags: []string{"Культурные объекты и достопримечательности"},
			Photos: []string{"https://i.ibb.co/jzQ6KNp/getlstd-property-photo.jpg"},
			Author: "Александра Волынец", Review: "Красивые закаты! Вкусное мороженое...",
		},
		{
			Name: "Teatro Nacional Ruben Dario", Tags: []string{"Театры"},
			Photos: []string{"https://i.ibb.co/0jyh8W4/teatro-nacional-ruben.jpg"},
			Author: "Ксения Самойлова", Review: "Такого Шекспира я еще не видела! Кошки были очень милые",
		},
		{
			Name: "Metrocentro", Tags: []string{"Торговые центры"},
			Photos: []string{"https://i.ibb.co/7W00vQk/photo1jpg.jpg", "https://i.ibb.co/f0tzdc1/img-20170412-125048-largejpg.jpg"},
			Author: "Андрей Кравцов", Review: "Обокрали! Обманули! Ставлю -100/1",
		},
		{
			Name: "The National Palace of Culture", Tags: []string{"Архитектурные достопримечательности", "Правительственные здания"},
			Photos: []string{"https://i.ibb.co/d5rpG9T/national-palace-of-culture.jpg", "https://i.ibb.co/L1jwLTR/caption.jpg"},
			Author: "Андрей Кравцов", Review: "Цена за вход: 15$ с человека. Как то маловато для такого здания...",
		},
	},
}

//var CookieDB = map[string]int{} // string - cookie hash; int - user id
