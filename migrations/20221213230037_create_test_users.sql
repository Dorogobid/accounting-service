-- +goose Up
-- +goose StatementBegin
INSERT INTO clients (
    name, 
    client_type_id, 
    tax_code, 
    address, 
    phone_number
) VALUES (
    'Приватне підприємство "Лідер"', 1, '22156234', '', '0509410248'
);
INSERT INTO clients (
    name, 
    client_type_id, 
    tax_code, 
    address, 
    phone_number
) VALUES (
    'Приватне підприємство "Аском"', 2, '99999903', 'м. Запоріжжя', ''
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM clients;
-- +goose StatementEnd
