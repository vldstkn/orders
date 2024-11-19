-- +goose Up
-- +goose StatementBegin
CREATE TABLE Offerings(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    status offering_status DEFAULT 'active',
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    price INT CHECK (price > 0 AND price < 25000),
    contractor_id INT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (contractor_id) REFERENCES Users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE offerings;
-- +goose StatementEnd
