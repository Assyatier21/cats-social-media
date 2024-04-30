CREATE TABLE IF NOT EXISTS match_cats (
    id SERIAL PRIMARY KEY,
    issued_by_id INT,
    target_user_id INT,
    match_cat_id INT,
    user_cat_id INT,
    message VARCHAR(120),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
