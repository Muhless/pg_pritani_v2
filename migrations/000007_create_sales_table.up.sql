CREATE TABLE
     IF NOT EXISTS sales (
          id BIGSERIAL PRIMARY KEY,
          customer_id BIGINT REFERENCES customers (id),
          total_price NUMERIC(10, 2) NOT NULL,
          paid_amount NUMERIC(10, 2) NOT NULL DEFAULT 0,
          remaining_amount NUMERIC(10, 2),
          status VARCHAR(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'paid', 'cancelled')),
          date TIMESTAMP NOT NULL,
          created_at TIMESTAMP DEFAULT NOW (),
          updated_at TIMESTAMP DEFAULT NOW (),
          deleted_at TIMESTAMP
     )