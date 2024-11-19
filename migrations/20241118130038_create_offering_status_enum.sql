-- +goose Up
-- +goose StatementBegin
CREATE TYPE offering_status AS ENUM(
    'active',
    'inactive');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE offering_status;
-- +goose StatementEnd
