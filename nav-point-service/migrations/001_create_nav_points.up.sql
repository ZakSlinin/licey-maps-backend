CREATE TABLE nav_points (
    id SERIAL PRIMARY KEY,
    orientation: TEXT NOT NULL,
    room: TEXT,
    type: TEXT,
    floor: INT
);