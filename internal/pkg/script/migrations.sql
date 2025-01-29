-- Users table
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       full_name VARCHAR(100) NOT NULL,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(20) CHECK (role IN ('admin', 'teacher', 'student')),
                       is_active BOOLEAN DEFAULT true,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP DEFAULT NULL
);

-- Subjects table
CREATE TABLE subjects (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(100) NOT NULL,
                          description TEXT,
                          is_active BOOLEAN DEFAULT true,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          deleted_at TIMESTAMP DEFAULT NULL
);

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

-- Indexes for better performance
CREATE INDEX idx_tests_subject ON tests(subject_id);
CREATE INDEX idx_tests_teacher ON tests(teacher_id);
CREATE INDEX idx_questions_test ON questions(test_id);
CREATE INDEX idx_answers_question ON answers(question_id);
CREATE INDEX idx_attempts_user ON test_attempts(user_id);
CREATE INDEX idx_attempts_test ON test_attempts(test_id);
CREATE INDEX idx_user_answers_attempt ON user_answers(attempt_id);

-- Add basic user roles
INSERT INTO users (full_name, username, password, role) VALUES
                                                            ('System Admin', 'admin', 'hashed_password', 'admin'),
                                                            ('Test Teacher', 'teacher', 'hashed_password', 'teacher'),
                                                            ('Test Student', 'student', 'hashed_password', 'student');

-- Add sample subject
INSERT INTO subjects (name, description) VALUES
                                             ('Mathematics', 'Basic mathematics course'),
                                             ('Physics', 'Basic physics course');