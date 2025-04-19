-- name: LikeRecipe :exec
INSERT INTO likes (user_id, recipe_id)
VALUES ($1, $2)
ON CONFLICT (user_id, recipe_id) DO NOTHING;

-- name: UnlikeRecipe :exec
DELETE FROM likes
WHERE user_id = $1 AND recipe_id = $2;

-- name: GetLikesCount :one
SELECT COUNT(*) AS likes_count
FROM likes
WHERE recipe_id = $1;
