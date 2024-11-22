CREATE SCHEMA IF NOT EXISTS music_library;

CREATE TABLE IF NOT EXISTS songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_name VARCHAR(255) NOT NULL,
    release_date DATE,
    text TEXT,
    link VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);