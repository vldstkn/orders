-- +goose Up
-- +goose StatementBegin
CREATE TABLE Orders(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    status order_status DEFAULT 'expectation',
    customer_id int NOT NULL,
    contractor_id int NOT NULL,
    offering_id int NOT NULL,
    CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES users(id),
    CONSTRAINT fk_contractor_id FOREIGN KEY (contractor_id) REFERENCES users(id),
    CONSTRAINT fk_offering_id FOREIGN KEY (offering_id) REFERENCES offerings(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS Orders cascade;
-- +goose StatementEnd
