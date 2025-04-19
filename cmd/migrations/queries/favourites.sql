-- name: FavouriteRecipe :exec
INSERT INTO favourites (user_id, recipe_id)
VALUES ($1, $2)
ON CONFLICT (user_id, recipe_id) DO NOTHING;

-- name: UnfavouriteRecipe :exec
DELETE FROM favourites
WHERE user_id = $1 AND recipe_id = $2;
