-- Questions table
CREATE TABLE questions (
                           id SERIAL PRIMARY KEY,
                           test_id INTEGER REFERENCES tests(id),
                           question_text TEXT NOT NULL,
                           points INTEGER DEFAULT 1,
                           question_type VARCHAR(20) CHECK (question_type IN ('single', 'multiple')),
                           order_number INTEGER NOT NULL,
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                           deleted_at TIMESTAMP DEFAULT NULL,
                           UNIQUE (test_id, order_number)
);