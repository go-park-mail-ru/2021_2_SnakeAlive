CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = NOW();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TABLE Users (
  id SERIAL NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  surname TEXT NOT NULL,
  email TEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  avatar TEXT,
  description TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON Users
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE Cookies (
  id SERIAL NOT NULL PRIMARY KEY,
  hash TEXT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE Places (
  id SERIAL NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  country TEXT NOT NULL,
  rating REAL NOT NULL,
  description TEXT,
  tags TEXT[],
  photos TEXT[]
);

CREATE TABLE Trips (
  id SERIAL NOT NULL PRIMARY KEY,
  title TEXT NOT NULL,
  origin INT,
  description TEXT,
  days INT NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (origin) REFERENCES places(id) ON DELETE CASCADE
);

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON Trips
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

CREATE TABLE TripsPlaces (
  id SERIAL NOT NULL PRIMARY KEY,
  place_id INT NOT NULL,
  trip_id INT NOT NULL,
  day INT NOT NULL,
  "order" INT NOT NULL,
  CONSTRAINT fk_trip FOREIGN KEY(trip_id) REFERENCES trips(id) ON DELETE CASCADE,
  CONSTRAINT fk_place FOREIGN KEY(place_id) REFERENCES places(id) ON DELETE CASCADE
);

CREATE TABLE Reviews (
  id SERIAL NOT NULL PRIMARY KEY,
  title TEXT,
  text TEXT,
  rating INT NOT NULL,
  user_id INT NOT NULL,
  place_id INT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY (place_id) REFERENCES places(id) ON DELETE CASCADE
);

CREATE TABLE Countries (
  id SERIAL NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  photo TEXT
);

INSERT INTO Countries ("name", "description", "photo")
VALUES ('Россия', 'Россия – крупнейшая страна мира, расположенная в Восточной Европе и Северной Азии и омываемая водами Тихого и Северного Ледовитого океанов.', 'russia.jpeg');
INSERT INTO Countries ("name", "description", "photo")
VALUES ('Германия', 'Германия – государство в Западной Европе с лесами, реками, горными хребтами и пляжными курортами Северного моря.', 'germany.jpeg');
INSERT INTO Countries ("name", "description", "photo")
VALUES ('США', ' Соединенные Штаты Америки – государство, состоящее из 50 штатов, занимает значительную часть Северной Америки. ', 'usa.jpeg');
INSERT INTO Countries ("name", "description", "photo")
VALUES ('Великобритания', 'Великобритания (официальное название – Соединенное Королевство Великобритании и Северной Ирландии) – островное государство на северо-западе Европы, состоящее из Англии, Шотландии, Уэльса и Северной Ирландии. ', 'uk.jpeg');

INSERT INTO Users ("name", "surname", "password", "email", "avatar")
VALUES ('Алексадра', 'Волынец', 'password', 'alex@mail.ru', 'test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "avatar")
VALUES ('Никита', 'Черных', 'frontend123', 'nikita@mail.ru', 'test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "avatar")
VALUES ('Ксения', 'Самойлова', '12345678', 'ksenia@mail.ru', 'test.jpeg');
INSERT INTO Users ("name", "surname", "password", "email", "avatar")
VALUES ('Андрей', 'Кравцов', '000111000', 'andrew@mail.ru', 'test.jpeg');

INSERT INTO Places ("name", "country", "rating", "description", "tags", "photos")
VALUES ('Цирк', 'Россия', 5, 'Коровы крутые очень! Огонь!',
ARRAY['Опасно', 'Есть парковка'], ARRAY['test.jpeg', 'test.jpeg']);
INSERT INTO Places ("name", "country", "rating", "description", "tags", "photos")
VALUES ('Кафе', 'Россия', 2.3, 'Обслуживание на 2-',
ARRAY['Еда', 'Есть парковка'], ARRAY['test.jpeg', 'test.jpeg']);
INSERT INTO Places ("name", "country", "rating", "description", "tags", "photos")
VALUES ('Церковь', 'Россия', 4.5, 'Самое святое место после Иерусалима',
ARRAY['Святые места'], ARRAY['test.jpeg']);
INSERT INTO Places ("name", "country", "rating", "description", "tags", "photos")
VALUES ('Музей', 'Россия', 3.5, 'Огонь!',
ARRAY['Исскуство', '18+'], ARRAY['test.jpeg', 'test.jpeg']);

INSERT INTO public.reviews (id, title, text, rating, user_id, place_id, created_at) VALUES (DEFAULT, 'title', 'text', 10, 1, 1, DEFAULT);
INSERT INTO public.reviews (id, title, text, rating, user_id, place_id, created_at) VALUES (DEFAULT, 'title2', 'text2', 11, 1, 1, DEFAULT);
INSERT INTO public.reviews (id, title, text, rating, user_id, place_id, created_at) VALUES (DEFAULT, 'title3', 'text3', 12, 1, 2, DEFAULT);

