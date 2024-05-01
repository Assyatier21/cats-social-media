-- CREATE TABLE IF EXISTS match_cats (
--     id SERIAL PRIMARY KEY,
--     issued_by_id INT,
--     target_user_id INT,
--     match_cat_id INT,
--     user_cat_id INT,
--     message VARCHAR(120),
--     status VARCHAR(20),
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     deleted_at TIMESTAMP
-- );

ALTER TABLE match_cats
ADD COLUMN status VARCHAR(20) DEFAULT NULL;
