-- Создание таблицы colors
CREATE TABLE if not exists colors (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    hex VARCHAR(7) NOT NULL UNIQUE CHECK (char_length(hex) = 7)
);