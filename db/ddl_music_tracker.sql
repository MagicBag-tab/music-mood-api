/*
 * DDL for music tracker database
 * Sarah Rachel Estrada Bonilla - 24347
*/

-- Tabla de artistas
CREATE TABLE artists (
    id          UNIQUE NOT NULL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    country     VARCHAR(100),
    image_path  VARCHAR(500),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabla de álbumes
CREATE TABLE albums (
    id           UNIQUE NOT NULL PRIMARY KEY,
    artist_id    UNIQUE NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    title        VARCHAR(255) NOT NULL,
    release_date DATE,
    cover_path   VARCHAR(500),
    created_at   TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabla de canciones
CREATE TABLE songs (
    id          UNIQUE NOT NULL PRIMARY KEY,
    artist_id   UNIQUE NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    album_id    UNIQUE REFERENCES albums(id) ON DELETE SET NULL,
    title       VARCHAR(255) NOT NULL,
    mood        VARCHAR(50) NOT NULL CHECK (mood IN ('happy','sad','energetic','calm','angry','relaxed')),
    source      VARCHAR(20) NOT NULL DEFAULT 'manual' CHECK (source IN ('manual','spotify')),
    spotify_id  VARCHAR(255) UNIQUE,
    image_path  VARCHAR(500),
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tabla de géneros
CREATE TABLE genres (
    id   UNIQUE NOT NULL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

-- Tabla de canciones y géneros
CREATE TABLE song_genres (
    song_id  UNIQUE NOT NULL PRIMARY KEY,
    genre_id UNIQUE NOT NULL PRIMARY KEY REFERENCES genres(id) ON DELETE CASCADE,
    song_id  UUID NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    genre_id UUID NOT NULL REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (song_id, genre_id)
);

-- Tabla de ratings
CREATE TABLE ratings (
    id         UNIQUE NOT NULL PRIMARY KEY,
    song_id    UNIQUE NOT NULL REFERENCES songs(id) ON DELETE CASCADE,
    score      INT NOT NULL CHECK (score >= 1 AND score <= 5),
    comment    TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Index para búsqueda y almacenamiento
CREATE INDEX idx_songs_mood      ON songs(mood);
CREATE INDEX idx_songs_source    ON songs(source);
CREATE INDEX idx_songs_title     ON songs(title);
CREATE INDEX idx_songs_artist_id ON songs(artist_id);