CREATE TABLE IF NOT EXISTS products (
     id BIGSERIAL PRIMARY KEY,
     name VARCHAR(50) NOT NULL,
     type VARCHAR(30) NOT NULL,
     stock INTEGER,
     price NUMERIC(10, 2),
     photo TEXT,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW(),
     deleted_at TIMESTAMP
)