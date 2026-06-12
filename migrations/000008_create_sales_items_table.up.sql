CREATE TABLE
     IF NOT EXISTS sales_items (
          id BIGSERIAL PRIMARY KEY,
          sales_id BIGINT NOT NULL REFERENCES sales (id),
          product_id BIGINT NOT NULL REFERENCES products (id),
          quantity INT NOT NULL,
          price NUMERIC(10, 2) NOT NULL,
          sub_total NUMERIC(10, 2) NOT NULL,
          created_at TIMESTAMP DEFAULT NOW (),
          updated_at TIMESTAMP DEFAULT NOW (),
          deleted_at TIMESTAMP
     )