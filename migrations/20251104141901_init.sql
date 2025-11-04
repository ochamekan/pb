-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS articles (
  id serial PRIMARY KEY,
  title text NOT NULL,
  body text NOT NULL,
  created_at timestamp DEFAULT now()
);

SELECT 'articles table succesfully created';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS articles;
-- +goose StatementEnd
