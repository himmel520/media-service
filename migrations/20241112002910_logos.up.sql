DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'image_type') THEN
        CREATE TYPE image_type AS ENUM ('adv', 'logo');
    END IF;
END $$;


-- Создание таблицы images
CREATE TABLE if not exists images (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    title VARCHAR(100) NOT NULL,
    type image_type,
    UNIQUE (type, url)
);