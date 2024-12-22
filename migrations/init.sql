create table products
(
    id       serial primary key,
    name     varchar(255) not null,
    price    float        not null,
    quantity integer      not null
);

INSERT INTO products (name, price, quantity)
VALUES ('Smartphone', 799.99, 50),
       ('Tablet', 499.99, 30),
       ('Headphones', 199.99, 100),
       ('Smartwatch', 299.99, 75),
       ('Camera', 899.99, 20),
       ('Printer', 149.99, 40),
       ('Monitor', 249.99, 60),
       ('Keyboard', 49.99, 150),
       ('Mouse', 29.99, 200),
       ('External Hard Drive', 99.99, 80);

-- Creating the orders table
CREATE TABLE orders
(
    order_id      SERIAL PRIMARY KEY, -- Unique identifier for the order
    customer_name VARCHAR(100),       -- Customer's name
    order_date    DATE,               -- Date of the order
    total_amount  DECIMAL(10, 2)      -- Total amount of the order
);

-- Inserting 10 records into the orders table
INSERT INTO orders (customer_name, order_date, total_amount)
VALUES ('John Smith', '2024-12-01', 1500.50),
       ('Peter Johnson', '2024-12-02', 2000.00),
       ('Olivia Brown', '2024-12-03', 750.75),
       ('Anna Davis', '2024-12-04', 3000.00),
       ('James Wilson', '2024-12-05', 1200.00),
       ('Emily Garcia', '2024-12-06', 500.25),
       ('Michael Martinez', '2024-12-07', 1000.00),
       ('Sophia Hernandez', '2024-12-08', 2500.00),
       ('David Lopez', '2024-12-09', 800.00),
       ('Emma Gonzalez', '2024-12-10', 1700.00);