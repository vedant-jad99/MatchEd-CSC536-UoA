-- Create the role if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_roles WHERE rolname = 'match_rw') THEN
        CREATE ROLE match_rw;
    END IF;
END $$;

-- Create the user if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_user WHERE usename = 'match_user') THEN
        CREATE USER match_user WITH PASSWORD 'swifty';
    END IF;
END $$;

-- Grant the role to the user if it's not already granted
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_user u
                   JOIN pg_catalog.pg_auth_members am ON u.usesysid = am.member
                   JOIN pg_catalog.pg_roles r ON am.roleid = r.oid
                   WHERE u.usename = 'match_user' AND r.rolname = 'match_rw') THEN
        GRANT match_rw TO match_user;
    END IF;
END $$;

-- Create the database if it doesn't exist
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_catalog.pg_database WHERE datname = 'match_db') THEN
        CREATE DATABASE match_db WITH OWNER = match_rw;
    END IF;
END $$;

-- Switch to the match_db database
\connect match_db;

-- Create the schema if it doesn't exist
CREATE SCHEMA IF NOT EXISTS match_schema AUTHORIZATION match_rw;

-- Reconnect as match_user
\connect match_db match_user;

-- User Table
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
    number VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    location VARCHAR NOT NULL,
    semesters VARCHAR NOT NULL
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