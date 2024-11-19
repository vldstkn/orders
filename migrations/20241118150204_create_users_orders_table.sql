-- +goose Up
-- +goose StatementBegin
CREATE TABLE users_orders (
    user_id int,
    order_id int,
    PRIMARY KEY (user_id, order_id),
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id),
    CONSTRAINT fk_order FOREIGN KEY (order_id) REFERENCES orders(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_orders cascade;
-- +goose StatementEnd
