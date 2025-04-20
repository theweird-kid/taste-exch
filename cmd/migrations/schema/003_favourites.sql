-- +goose Up
CREATE TABLE favourites (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users (id) ON DELETE CASCADE,
    recipe_id INT REFERENCES recipes (id) ON DELETE CASCADE,
    UNIQUE (user_id, recipe_id) -- Ensures idempotency
);

-- +goose Down
DROP TABLE IF EXISTS favourites CASCADE;
