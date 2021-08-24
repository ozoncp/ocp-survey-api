-- +goose Up
-- +goose StatementBegin
CREATE TABLE surveys (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL,
  link TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE surveys;
-- +goose StatementEnd
