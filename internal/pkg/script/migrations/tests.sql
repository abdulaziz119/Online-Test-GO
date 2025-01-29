-- Tests table
CREATE TABLE tests (
                       id SERIAL PRIMARY KEY,
                       subject_id INTEGER REFERENCES subjects(id),
                       teacher_id INTEGER REFERENCES users(id),
                       title VARCHAR(255) NOT NULL,
                       description TEXT,
                       duration_minutes INTEGER NOT NULL,
                       pass_percent INTEGER NOT NULL CHECK (pass_percent BETWEEN 0 AND 100),
                       start_time TIMESTAMP NOT NULL,
                       end_time TIMESTAMP NOT NULL,
                       is_active BOOLEAN DEFAULT true,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP DEFAULT NULL,
                       CONSTRAINT valid_time_range CHECK (end_time > start_time)
);