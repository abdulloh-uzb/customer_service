CREATE TABLE if not exists addresses (
    id bigserial PRIMARY KEY,
    street varchar NOT NULL,
    district varchar NOT NULL,
    customer_id int REFERENCES customers(id) NOT NULL
);
