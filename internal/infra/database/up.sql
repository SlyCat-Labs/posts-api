DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id VARCHAR(36) PRIMARY KEY,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

DROP TABLE IF EXISTS posts;

CREATE TABLE posts (
    id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(32) NOT NULL,
    content VARCHAR(255) NOT NULL,
    author_id VARCHAR(36) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,
    
    FOREIGN KEY(author_id) REFERENCES users(id)
);