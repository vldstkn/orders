-- +goose Up
-- +goose StatementBegin
CREATE TYPE order_status AS ENUM(
    'expectation',
    'confirmed',
    'in_processing',
    'completed',
    'closed');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE order_status;
-- +goose StatementEnd
