CREATE TABLE points (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    env TEXT[],
    nav_point INT,
    FOREIGN KEY (nav_point) REFERENCES nav_points(id)
);
