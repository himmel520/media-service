-- Создание таблицы logos
CREATE TABLE if not exists logos (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL UNIQUE,
    title VARCHAR(100) NOT NULL
);