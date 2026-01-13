DROP TABLE IF EXISTS book;
CREATE TABLE book (
  id      SERIAL PRIMARY KEY,
  title   VARCHAR(128) NOT NULL,
  author  VARCHAR(255) NOT NULL,
  price   DECIMAL(5,2) NOT NULL
);

INSERT INTO book
  (title, author, price)
VALUES
  ('Clean Code', 'Uncle Bob', 21.12),
  ('The Pragmatic Programmer', 'David Thomas', 31.99),
  ('Designing Data-Intensive Applications', 'Martin Kleppmann', 37.00),
  ('Code Complete', 'Steve McConnell', 40.01);
