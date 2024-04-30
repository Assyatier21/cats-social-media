CREATE TABLE IF NOT EXISTS cats (
    id SERIAL PRIMARY KEY,
    user_id INTEGER, -- REFERENCES users(id),
    name VARCHAR(255),
    race VARCHAR(255),
    sex VARCHAR(255),
    age INTEGER,
    description VARCHAR(255),
    images TEXT[],
    is_already_matched BOOLEAN,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);