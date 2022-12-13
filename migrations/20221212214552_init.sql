-- +goose Up
-- +goose StatementBegin
CREATE TABLE client_types (
    id serial PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO client_types (name) VALUES ('Покупець');
INSERT INTO client_types (name) VALUES ('Постачальник');
INSERT INTO client_types (name) VALUES ('Співробітниик');

CREATE TABLE clients (
    id serial PRIMARY KEY,
    name text NOT NULL,
    client_type_id bigint NOT NULL REFERENCES client_types(id),
    tax_code varchar(10) NOT NULL DEFAULT '',
    address text NOT NULL DEFAULT '',
    phone_number text NOT NULL DEFAULT ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE clients;
DROP TABLE client_types;
-- +goose StatementEnd
