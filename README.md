# MatchEd-CSC536-UoA
A software engineering project, matching faculty/educators to courses being taught in CS department.

## Team and Roles:

- Vedant Jadhav : Scrum master, Developer
- Krishna Prashanth Thummanapelly : Developer
- Ramya Ramachandran: Developer
- Adam Cunningham: Product Owner, Developer

## Start the services
    $ docker compose up

## Database setup
    $ psql -U postgres -h localhost -f db.sql

## Database viewing
    SELECT table_name FROM information_schema.tables WHERE table_schema = 'match_schema';
    SET search_path TO match_schema;
    \dt
    SELECT * FROM match_schema."user";

    note that "user" is a reserved keyword in postgresql, so quotes are needed

