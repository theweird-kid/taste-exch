-- name: CreateRecipe :one
INSERT INTO recipes (user_id, title, description, tags, ingredients, instructions, total_time, difficulty, servings, photo_url)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING id;

-- name: GetRecipeById :one
SELECT
    id,
    user_id,
    title,
    description,
    tags,
    ingredients,
    instructions,
    total_time,
    difficulty,
    servings,
    photo_url,
    created_at
FROM
    recipes
WHERE
    id = $1;

-- name: GetRecipes :many
SELECT
    title AS recipe_name,
    description AS recipe_description,
    photo_url,
    tags
FROM
    recipes
ORDER BY
    created_at DESC
LIMIT 10 OFFSET $1;
