-- Создание таблицы tg
CREATE TABLE if not exists tg (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    url TEXT NOT NULL UNIQUE
);