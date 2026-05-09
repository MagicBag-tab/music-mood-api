/*
 * DML for music tracker database
 * Basado en gustos y vibes de Sarah 
*/

-- =========================
-- INSERT ARTISTS
-- =========================
INSERT INTO artists (id, name, country, image_path) VALUES
(1, 'Taylor Swift', 'United States', '/images/artists/taylor.jpg'),
(2, 'Arctic Monkeys', 'United Kingdom', '/images/artists/arctic_monkeys.jpg'),
(3, 'Mitski', 'Japan/United States', '/images/artists/mitski.jpg'),
(4, 'Stray Kids', 'South Korea', '/images/artists/stray_kids.jpg'),
(5, 'Chase Atlantic', 'Australia', '/images/artists/chase_atlantic.jpg'),
(6, 'BTS', 'South Korea', '/images/artists/bts.jpg'),
(7, 'Olivia Rodrigo', 'United States', '/images/artists/olivia.jpg'),
(8, 'Ariana Grande', 'United States', '/images/artists/ariana.jpg'),
(9, 'TXT', 'South Korea', '/images/artists/txt.jpg'),
(10, 'The Neighbourhood', 'United States', '/images/artists/neighbourhood.jpg');

-- =========================
-- INSERT ALBUMS
-- =========================
INSERT INTO albums (id, artist_id, title, release_date, cover_path) VALUES
(1,  1,  'folklore',                    '2020-07-24', '/covers/folklore.jpg'),
(2,  1,  'Midnights',                   '2022-10-21', '/covers/midnights.jpg'),
(3,  2,  'AM',                          '2013-09-09', '/covers/am.jpg'),
(4,  3,  'Be the Cowboy',               '2018-08-17', '/covers/bethecowboy.jpg'),
(5,  4,  'NOEASY',                      '2021-08-23', '/covers/noeasy.jpg'),
(6,  5,  'Phases',                      '2019-06-28', '/covers/phases.jpg'),
(7,  6,  'Love Yourself: Tear',         '2018-05-18', '/covers/tear.jpg'),
(8,  7,  'GUTS',                        '2023-09-08', '/covers/guts.jpg'),
(9,  8,  'eternal sunshine',            '2024-03-08', '/covers/eternal_sunshine.jpg'),
(10, 9,  'The Chaos Chapter: FREEZE',   '2021-05-31', '/covers/freeze.jpg'),
(11, 10, 'I Love You.',                 '2013-04-19', '/covers/ily.jpg');

-- =========================
-- INSERT GENRES
-- =========================
INSERT INTO genres (id, name) VALUES
(1,  'Pop'),
(2,  'K-Pop'),
(3,  'Alternative'),
(4,  'Indie'),
(5,  'Rock'),
(6,  'R&B'),
(7,  'Bedroom Pop'),
(8,  'Synthpop'),
(9,  'Emo Pop'),
(10, 'Hip Hop');

-- =========================
-- INSERT SONGS
-- =========================
INSERT INTO songs (id, artist_id, album_id, title, mood, source, spotify_id, image_path) VALUES
(1,  1,  1,    'cardigan',                'sad',      'spotify', '4R2kfaDFgnULkiCo3gahGb', '/songs/cardigan.jpg'),
(2,  1,  2,    'Maroon',                  'calm',     'spotify', '3yNGQUlHxMWMXmNHjQK2Ag', '/songs/maroon.jpg'),
(3,  2,  3,    'Do I Wanna Know?',        'calm',     'spotify', '5FVd6KXrgO9B3JPmC8OPst', '/songs/doiwannaknow.jpg'),
(4,  2,  3,    'R U Mine?',               'energetic','spotify', '4AnjggyFCSpAg6xDDFds9V', '/songs/rumine.jpg'),
(5,  3,  4,    'Nobody',                  'sad',      'spotify', '4zRoS7poMHCGkFoVERJCUl', '/songs/nobody.jpg'),
(6,  4,  5,    'Thunderous',              'energetic','spotify', '6nCfRBFKiMGJDDpjOUkZVi', '/songs/thunderous.jpg'),
(7,  4,  5,    'Red Lights',              'angry',    'spotify', '7sWRnu1oNHeYVLlJJKdFN6', '/songs/redlights.jpg'),
(8,  5,  6,    'Swim',                    'relaxed',  'spotify', '6e8ISVQiIzPUjcCUMZ7MXI', '/songs/swim.jpg'),
(9,  6,  7,    'Fake Love',               'sad',      'spotify', '6jG2YzhxptolDzLHTGXMfo', '/songs/fakelove.jpg'),
(10, 7,  8,    'vampire',                 'sad',      'spotify', '1kuGVB7EU95pJObxwvfwUs', '/songs/vampire.jpg'),
(11, 7,  8,    'bad idea right?',         'happy',    'spotify', '3IX0yuEVvjchTBWVCRRhX3', '/songs/badidea.jpg'),
(12, 8,  9,    'we can''t be friends',    'calm',     'spotify', '4Dy8JWMBMGbGCCACbfgWM3', '/songs/wcbf.jpg'),
(13, 9,  10,   '0X1=LOVESONG',            'angry',    'spotify', '3Wly5OKT2HqQWgxDZkMDgO', '/songs/lovesong.jpg'),
(14, 10, 11,   'Sweater Weather',         'relaxed',  'spotify', '6YJ4EgMp4PEBKIRoXRNkBz', '/songs/sweaterweather.jpg'),
(15, 3,  4,    'Washing Machine Heart',   'sad',      'spotify', '1SYKGjFSrm2f3s9P0h9ggi', '/songs/wmh.jpg');

-- =========================
-- INSERT SONG_GENRES
-- =========================
INSERT INTO song_genres (song_id, genre_id) VALUES
(1,  1), (1,  4),
(2,  1), (2,  8),
(3,  3), (3,  5),
(4,  5),
(5,  4),
(6,  2), (6,  10),
(7,  2), (7,  6),
(8,  6), (8,  7),
(9,  2), (9,  1),
(10, 9), (10, 1),
(11, 1),
(12, 1), (12, 8),
(13, 2), (13, 5),
(14, 7), (14, 3),
(15, 4);

-- =========================
-- INSERT RATINGS
-- =========================
INSERT INTO ratings (song_id, score, comment) VALUES
(1,  5, 'Perfect for rainy nights and overthinking'),
(2,  5, 'Literal midnight main character vibes'),
(3,  5, 'Timeless obsession'),
(4,  4, 'Makes me feel cooler than I actually am'),
(5,  5, 'Emotionally devastating in the best way'),
(6,  4, 'Gym arc soundtrack'),
(7,  5, 'Too intense but addictive'),
(8,  5, 'Late night drive energy'),
(9,  5, 'Classic heartbreak anthem'),
(10, 5, 'The drama. The pain. The vocals.'),
(11, 4, 'Chaotic but fun'),
(12, 5, 'Soft sadness but elegant'),
(13, 5, 'Sounds like surviving your villain arc'),
(14, 5, 'Ultimate comfort song'),
(15, 4, 'Mentally unstable but aesthetic');

-- =========================
-- RESET SEQUENCES
-- =========================
SELECT setval('artists_id_seq', (SELECT MAX(id) FROM artists));
SELECT setval('albums_id_seq',  (SELECT MAX(id) FROM albums));
SELECT setval('songs_id_seq',   (SELECT MAX(id) FROM songs));
SELECT setval('genres_id_seq',  (SELECT MAX(id) FROM genres));
SELECT setval('ratings_id_seq', (SELECT MAX(id) FROM ratings));