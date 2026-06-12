CREATE TABLE IF NOT EXISTS customers (
     id BIGSERIAL PRIMARY KEY,
     name VARCHAR(50) NOT NULL,
     phone VARCHAR(15) UNIQUE,
     address TEXT,
     company VARCHAR(50),
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW(),
     deleted_at TIMESTAMP
)