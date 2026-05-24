CREATE TABLE IF NOT EXISTS users (
     id BIGSERIAL PRIMARY KEY,
     username VARCHAR(50) NOT NULL,
     password TEXT NOT NULL,
     role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'employee')),
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW(),
     deleted_at TIMESTAMP
)