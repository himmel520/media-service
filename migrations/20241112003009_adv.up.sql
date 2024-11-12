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