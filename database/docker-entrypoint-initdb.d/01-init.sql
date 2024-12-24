CREATE TABLE IF NOT EXISTS users
(
    id       INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name     VARCHAR(100) NOT NULL,
    email    VARCHAR(254) NOT NULL UNIQUE,
    password VARCHAR(128) NOT NULL,
    CHECK ( name <> '' ),
    CHECK ( email <> '' ),
    CHECK ( password <> '' )
);


CREATE TABLE IF NOT EXISTS musicians
(
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR(254) NOT NULL,
    CHECK ( name <> '' )
);

CREATE TABLE IF NOT EXISTS albums
(
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name        VARCHAR(100) NOT NULL,
    cover_file  BYTEA        NOT NULL,
    musician_id INT          NOT NULL
        REFERENCES musicians (id)
            ON DELETE CASCADE,
    CHECK ( name <> '' ),
    CHECK ( length(cover_file) > 0 )

);

CREATE TABLE IF NOT EXISTS tracks
(
    id       INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payload   BYTEA        NOT NULL,
    name     VARCHAR(100) NOT NULL,
    album_id INT          NOT NULL
        REFERENCES albums (id)
            ON DELETE CASCADE,
    CHECK ( payload <> '' ),
    CHECK ( length(payload) > 0 ),
    CHECK ( name <> '' )
);
