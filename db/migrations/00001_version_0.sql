-- +goose Up
-- +goose StatementBegin
CREATE TABLE dogs (
    id SERIAL PRIMARY KEY,
    breed VARCHAR(100) NOT NULL,
    sub_breed VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd


-- +goose Down
-- +goose StatementBegin
DROP TABLE dogs;
-- +goose StatementEnd