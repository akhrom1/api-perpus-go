-- +migrate Up
CREATE TABLE loan_items (
    id SERIAL PRIMARY KEY,
    loan_id INT NOT NULL REFERENCES loans(id) ON DELETE CASCADE,
    book_id INT NOT NULL REFERENCES books(id) ON DELETE RESTRICT
);

-- +migrate Down
DROP TABLE IF EXISTS loan_items;
