CREATE SCHEMA IF NOT EXISTS auth_service;

CREATE TABLE IF NOT EXISTS auth_service.users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email VARCHAR(255) UNIQUE NOT NULL,
    token_hash VARCHAR(255),
    ip_address VARCHAR(45)
);

INSERT INTO auth_service.users (email, ip_address) VALUES ('user1@mail.ru', '1.1.1.1');
INSERT INTO auth_service.users (email, ip_address) VALUES ('user2@mail.ru', '2.2.2.2');
INSERT INTO auth_service.users (email, ip_address) VALUES ('user3@mail.ru', '3.3.3.3');

CREATE USER api_user WITH PASSWORD 'api_password';

GRANT USAGE ON SCHEMA auth_service TO api_user;

GRANT SELECT, INSERT, UPDATE ON auth_service.users TO api_user;