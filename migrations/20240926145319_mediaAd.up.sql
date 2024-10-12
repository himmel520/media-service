-- Создание таблицы logos
CREATE TABLE if not exists logos (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL UNIQUE,
    title VARCHAR(100) NOT NULL
);

-- Создание таблицы colors
CREATE TABLE if not exists colors (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    hex VARCHAR(7) NOT NULL UNIQUE CHECK (char_length(hex) = 7)
);

-- Создание таблицы tg
CREATE TABLE if not exists tg (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    url TEXT NOT NULL UNIQUE
);

-- Создание таблицы adv
CREATE TABLE if not exists adv (
    id SERIAL PRIMARY KEY,
    logos_id INTEGER NOT NULL REFERENCES logos(id),
    colors_id INTEGER NOT NULL REFERENCES colors(id),
    tg_id INTEGER NOT NULL REFERENCES tg(id),
    post VARCHAR(100) NOT NULL,
    title VARCHAR(40) NOT NULL,
    description VARCHAR(150) NOT NULL,
    priority SMALLINT NOT NULL CHECK (priority IN (1, 2, 3))
);