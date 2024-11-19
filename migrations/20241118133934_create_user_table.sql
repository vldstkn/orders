-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Users(
   id SERIAL PRIMARY KEY,
   created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
   email VARCHAR(150) UNIQUE NOT NULL,
   password VARCHAR(300) NOT NULL,
   name VARCHAR(50) NOT NULL,
   role user_role DEFAULT 'customer',
   rating FLOAT DEFAULT 0.0,
   number_completed_orders INT DEFAULT 0
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Users cascade;
-- +goose StatementEnd
