CREATE TABLE IF NOT EXISTS suppliers (
     id BIGSERIAL PRIMARY KEY,
     name VARCHAR(50) NOT NULL,
     email VARCHAR(50),
     phone VARCHAR(15) NOT NULL UNIQUE,
     address TEXT,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW(),
     deleted_at TIMESTAMP
)