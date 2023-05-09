DROP TABLE IF EXISTS watchlist;
CREATE TABLE watchlist (
                       id         INT AUTO_INCREMENT NOT NULL,
                       title      VARCHAR(128) NOT NULL,
                       mediaType  INT,
                       genre      VARCHAR(255),
                       streamingPlatform VARCHAR(255),
                       shortNote VARCHAR(255) NOT NULL,
                       PRIMARY KEY (`id`)
);

INSERT INTO watchlist
(title, mediaType, genre, streamingPlatform, shortNote)
VALUES
    ('Big Little Lies', 1, 'Drama', 'HBO', ''),
    ('Little Fires Everywhere', 1, 'Drama', 'Hulu', ''),
    ('Dead to Me (Netflix)', 1, 'Comedy', 'Netflix', ''),
    ('Reinventing Anna', 1, 'Drama', 'Netflix', 'Wanted to watch this');