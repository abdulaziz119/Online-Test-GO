-- Answers table
CREATE TABLE answers (
                         id SERIAL PRIMARY KEY,
                         question_id INTEGER REFERENCES questions(id),
                         answer_text TEXT NOT NULL,
                         is_correct BOOLEAN NOT NULL,
                         order_number INTEGER NOT NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         deleted_at TIMESTAMP DEFAULT NULL,
                         UNIQUE (question_id, order_number)
);