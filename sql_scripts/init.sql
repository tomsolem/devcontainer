CREATE DATABASE inventory;

\c inventory;

CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    type VARCHAR(100) NOT NULL,
    quantity INT NOT NULL DEFAULT 0
);

INSERT INTO equipment (name, type, quantity) VALUES
('Hammer', 'Tool', 10),
('Laptop', 'Electronics', 5),
('Desk', 'Furniture', 20);