CREATE TABLE reviews (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    score INT CHECK (score >= 1 AND score <= 10)
);