CREATE TABLE IF NOT EXISTS admins (
     id BIGSERIAL PRIMARY KEY,
     name VARCHAR(50) NOT NULL,
     email VARCHAR(50) NOT NULL UNIQUE,
     phone VARCHAR(15) NOT NULL UNIQUE,
     photo TEXT,
     is_active CHECK,
     created_at TIMESTAMP DEFAULT NOW(),
     updated_at TIMESTAMP DEFAULT NOW(),
     deleted_at TIMESTAMP
)