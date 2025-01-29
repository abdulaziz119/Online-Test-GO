-- User answers table
CREATE TABLE user_answers (
                              id SERIAL PRIMARY KEY,
                              attempt_id INTEGER REFERENCES test_attempts(id),
                              question_id INTEGER REFERENCES questions(id),
                              answer_id INTEGER REFERENCES answers(id),
                              answered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                              deleted_at TIMESTAMP DEFAULT NULL
);