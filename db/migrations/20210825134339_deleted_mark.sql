-- +goose Up
-- +goose StatementBegin
ALTER TABLE surveys
ADD COLUMN IF NOT EXISTS deleted BOOLEAN;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE surveys
DROP COLUMN IF EXISTS deleted;
-- +goose StatementEnd
