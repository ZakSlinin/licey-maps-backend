-- Тестовые данные для навигационных точек
INSERT INTO nav_points (orientation, neighbours, room, type, floor) VALUES
('север', '{2, 3}', 'Главный вход', 'entrance', 1),
('восток', '{1, 4}', 'Коридор 1', 'corridor', 1),
('юг', '{1, 5}', 'Коридор 2', 'corridor', 1),
('запад', '{2, 6}', 'Коридор 3', 'corridor', 1),
('север', '{3, 7}', 'Коридор 4', 'corridor', 1),
('юг', '{4, 8}', 'Коридор 5', 'corridor', 1),
('восток', '{5, 9}', 'Коридор 6', 'corridor', 1),
('запад', '{6, 10}', 'Коридор 7', 'corridor', 1),
('север', '{7, 11}', 'Коридор 8', 'corridor', 1),
('юг', '{8, 12}', 'Коридор 9', 'corridor', 1);

-- Тестовые данные для точек
INSERT INTO points (name, env, nav_point) VALUES
('Главный вход', '{"entrance", "main"}', 1),
('Кабинет 101', '{"classroom", "math"}', 2),
('Кабинет 102', '{"classroom", "physics"}', 3),
('Кабинет 103', '{"classroom", "chemistry"}', 4),
('Кабинет 104', '{"classroom", "biology"}', 5),
('Кабинет 105', '{"classroom", "history"}', 6),
('Кабинет 106', '{"classroom", "literature"}', 7),
('Кабинет 107', '{"classroom", "geography"}', 8),
('Кабинет 108', '{"classroom", "art"}', 9),
('Кабинет 109', '{"classroom", "music"}', 10);
