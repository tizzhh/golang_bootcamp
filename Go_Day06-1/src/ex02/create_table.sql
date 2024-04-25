CREATE TABLE IF NOT EXISTS article (
    id SERIAL PRIMARY KEY,
    post_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title VARCHAR(100) NOT NULL,
    article_text TEXT NOT NULL
);

INSERT INTO article (title, article_text)
VALUES ('Test title', 'this is an article text');