CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    genre_id INT REFERENCES genres(id) ON DELETE CASCADE,
    rating FLOAT,
    CONSTRAINT fk_genre FOREIGN KEY (genre_id) REFERENCES genres(id)
);