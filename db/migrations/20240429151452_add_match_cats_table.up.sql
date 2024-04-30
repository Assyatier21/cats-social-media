CREATE TABLE IF NOT EXISTS match_cats (
    id SERIAL PRIMARY KEY,
    issued_by_id INTEGER, -- REFERENCES users(id),
    target_user_id INTEGER, -- REFERENCES users(id),
    user_cat_id INTEGER, -- REFERENCES cats(id),
    match_cat_id INTEGER, -- REFERENCES cats(id),
    message VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);