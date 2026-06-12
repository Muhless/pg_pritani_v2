CREATE TABLE
     IF NOT EXISTS payments (
          id BIGSERIAL PRIMARY KEY,
          sales_id BIGINT NOT NULL REFERENCES sales (id),
          amount NUMERIC(10, 2) NOT NULL,
          method VARCHAR(20) NOT NULL,
          date TIMESTAMP NOT NULL,
          note TEXT,
          created_at TIMESTAMP DEFAULT NOW (),
          updated_at TIMESTAMP DEFAULT NOW (),
          deleted_at TIMESTAMP
     )