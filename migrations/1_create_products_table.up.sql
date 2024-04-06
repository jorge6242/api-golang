-- +migrate Up
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    precio NUMERIC NOT NULL
);
