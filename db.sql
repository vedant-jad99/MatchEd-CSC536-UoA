CREATE ROLE match_rw;

CREATE USER match_user WITH PASSWORD 'swifty';
GRANT match_rw TO match_user;

CREATE DATABASE match_db WITH OWNER = match_rw;

\connect match_db;
CREATE SCHEMA IF NOT EXISTS match_schema AUTHORIZATION match_rw;

\connect match_db match_user;
CREATE TABLE IF NOT EXISTS match_schema.user (
    id SERIAL PRIMARY KEY,
    name VARCHAR,
    email VARCHAR
);

-- User Role Table
CREATE TABLE IF NOT EXISTS match_schema.role (
    user_id SERIAL PRIMARY KEY,
    role VARCHAR NOT NULL
);

-- Authentication Table
CREATE TABLE IF NOT EXISTS match_schema.auth (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES match_schema.role(user_id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    auth_token VARCHAR NOT NULL UNIQUE
);

-- Course Table
CREATE TABLE IF NOT EXISTS match_schema.course (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    type VARCHAR NOT NULL
    -- Add other course-specific attributes here
);

-- Course Semester Table
CREATE TABLE IF NOT EXISTS match_schema.course_sem (
    id SERIAL PRIMARY KEY,
    course_id INT NOT NULL REFERENCES match_schema.course(id),
    semester VARCHAR NOT NULL,
    mandatory_level VARCHAR,
    timeslot VARCHAR
);

-- Preference Table
CREATE TABLE IF NOT EXISTS match_schema.preferences (
    id SERIAL PRIMARY KEY,
    semester VARCHAR NOT NULL,
    user_id INT NOT NULL,
    course_sem_id INT NOT NULL ,
    preference_level INT NOT NULL
);

-- Matching Iteration Table
CREATE TABLE IF NOT EXISTS match_schema.matching_iteration (
    id SERIAL PRIMARY KEY,
    triggered_by INT,
    status VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Matching Table
CREATE TABLE IF NOT EXISTS match_schema.matching (
    id SERIAL PRIMARY KEY,
    matching_iteration_id INT NOT NULL REFERENCES match_schema.matching_iteration(id),
    user_id INT NOT NULL REFERENCES match_schema.role(user_id),
    course_sem_id INT NOT NULL REFERENCES match_schema.course_sem(id),
    score DECIMAL NOT NULL
);

GRANT USAGE ON SCHEMA match_schema TO match_rw;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA match_schema TO match_rw;
GRANT USAGE, SELECT, UPDATE ON ALL SEQUENCES IN SCHEMA match_schema TO match_rw;



-- INSERT INTO match_schema.user (name, email) VALUES('Martin', 'martinmail@arizona.edu');
-- SELECT * FROM match_schema.user;