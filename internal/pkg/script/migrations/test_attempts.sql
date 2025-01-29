-- Test attempts table
CREATE TABLE test_attempts (
                               id SERIAL PRIMARY KEY,
                               user_id INTEGER REFERENCES users(id),
                               test_id INTEGER REFERENCES tests(id),
                               started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               finished_at TIMESTAMP,
                               score INTEGER,
                               status VARCHAR(20) CHECK (status IN ('started', 'finished')),
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               deleted_at TIMESTAMP DEFAULT NULL,
                               CONSTRAINT valid_attempt_time CHECK (finished_at > started_at)
);