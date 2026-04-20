-- +migrate Up
-- Insert sample users
INSERT INTO users (username, email, first_name, last_name, is_active) VALUES
    ('johndoe', 'john@example.com', 'John', 'Doe', true),
    ('janedoe', 'jane@example.com', 'Jane', 'Doe', true),
    ('bobsmith', 'bob@example.com', 'Bob', 'Smith', false);

-- Insert some cached factorial values
INSERT INTO factorial_cache (n, result) VALUES
    (0, 1),
    (1, 1),
    (5, 120),
    (10, 3628800);

-- +migrate Down
DELETE FROM factorial_cache;
DELETE FROM users;
