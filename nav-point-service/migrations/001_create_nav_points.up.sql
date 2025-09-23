CREATE TABLE nav_points (
    id SERIAL PRIMARY KEY,
    orientation TEXT NOT NULL,
    neighbours INT[],
    room TEXT,
    type TEXT,
    floor INT
);