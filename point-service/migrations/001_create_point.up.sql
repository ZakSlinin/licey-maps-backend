CREATE TABLE point (
    id SERIAL PRIMARY KEY,
    name TEXT,
    env TEXT[],
    nav_point INT
)