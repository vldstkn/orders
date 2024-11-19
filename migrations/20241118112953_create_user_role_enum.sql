-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM
    ('contractor',
     'customer',
     'admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE user_role;
-- +goose StatementEnd
