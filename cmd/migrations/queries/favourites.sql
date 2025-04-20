-- name: FavouriteRecipe :exec
INSERT INTO favourites (user_id, recipe_id)
VALUES ($1, $2)
ON CONFLICT (user_id, recipe_id) DO NOTHING;

-- name: GetFavouriteRecipes :many
SELECT
    r.id AS recipe_id,
    r.title AS recipe_name,
    r.description AS recipe_description,
    r.photo_url,
    r.tags
FROM recipes r
JOIN favourites f ON r.id = f.recipe_id
WHERE f.user_id = $1;

-- name: UnfavouriteRecipe :exec
DELETE FROM favourites
WHERE user_id = $1 AND recipe_id = $2;
