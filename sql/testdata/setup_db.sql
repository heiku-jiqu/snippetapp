CREATE DATABASE test_snippetapp;

CREATE USER test_web WITH PASSWORD 'pass';
GRANT CREATE ON SCHEMA public TO test_web;
GRANT DELETE, UPDATE, INSERT, SELECT ON ALL TABLES IN SCHEMA public TO test_web;
ALTER DEFAULT PRIVILEGES IN SCHEMA public
  GRANT USAGE, SELECT ON SEQUENCES TO test_web;
