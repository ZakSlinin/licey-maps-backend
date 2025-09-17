CREATE TABLE nav_points (
    id SERIAL PRIMARY KEY,
    nav_point_id: INT,
    orientation: TEXT NOT NULL,
    room: TEXT,
    type: TEXT,
    floor: INT
);