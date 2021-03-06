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
  CONSTRAINT fk_trip FOREIGN KEY(trip_id) REFERENCES trips(id),
  CONSTRAINT fk_place FOREIGN KEY(place_id) REFERENCES places(id)
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