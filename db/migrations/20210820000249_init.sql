-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS surveys (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS surveys;
-- +goose StatementEnd
