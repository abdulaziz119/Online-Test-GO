-- Users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       full_name VARCHAR(100) NOT NULL,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(20) CHECK (role IN ('admin', 'teacher', 'student')),
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       is_active BOOLEAN DEFAULT true
);