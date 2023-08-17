CREATE TABLE studios (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birth_date DATE
);

CREATE TABLE directors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    birth_date DATE
);

CREATE TYPE rating AS ENUM ('PG-10', 'PG-13', 'PG-18');

CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    release_year INTEGER CHECK (release_year >= 1800),
    studio_id INTEGER REFERENCES studios(id),
    box_office DECIMAL(10,2),
    rating rating,
    UNIQUE (title, release_year)
);

CREATE TABLE movies_actors (
    movie_id INTEGER REFERENCES movies(id),
    actor_id INTEGER REFERENCES actors(id),
    PRIMARY KEY (movie_id, actor_id)
);

CREATE TABLE movies_directors (
    movie_id INTEGER REFERENCES movies(id),
    director_id INTEGER REFERENCES directors(id),
    PRIMARY KEY (movie_id, director_id)
);

INSERT INTO studios (name) VALUES
('Warner Bros.'),
('Universal Pictures'),
('20th Century Fox'),
('Paramount Pictures');

INSERT INTO actors (name, birth_date) VALUES
('Johnny Depp', '1963-06-09'),
('Natalie Portman', '1981-06-09'),
('Leonardo DiCaprio', '1974-11-11'),
('Scarlett Johansson', '1984-11-22');

INSERT INTO directors (name, birth_date) VALUES
('Steven Spielberg', '1946-12-18'),
('Christopher Nolan', '1970-07-30'),
('James Cameron', '1954-08-16'),
('Quentin Tarantino', '1963-03-27');

INSERT INTO movies (title, release_year, studio_id, box_office, rating) VALUES
('Pirates of the Caribbean', 2003, 1, 6543210.50, 'PG-13'),
('Black Swan', 2010, 2, 5432109.50, 'PG-18'),
('Titanic', 1997, 3, 9876543.21, 'PG-13'),
('Inception', 2010, 4, 8765432.10, 'PG-13');

INSERT INTO movies_actors (movie_id, actor_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 3), 
(4, 4);

INSERT INTO movies_directors (movie_id, director_id) VALUES
(1, 4),
(2, 1),
(3, 3),
(4, 2);

-- выборка фильмов с названием студии;
SELECT 
    movies.id AS movie_id,
    movies.title AS movie_title,
    studios.name AS studio_name
FROM 
    movies
JOIN 
    studios ON movies.studio_id = studios.id;

-- подсчёт количества фильмов со сборами больше 1000
SELECT 
    COUNT(*) AS number_of_movies
FROM 
    movies
WHERE 
    box_office > 1000;

-- подсчёт количества фильмов, имеющих дубли по названию
SELECT 
    title, 
    COUNT(*) AS count
FROM 
    movies
GROUP BY 
    title
HAVING 
    COUNT(*) > 1;
