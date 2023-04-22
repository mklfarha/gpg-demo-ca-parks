
-- name: UpdatePark :exec
UPDATE park
SET
name = ?, main_image = ?, phone = ?, hours = ?, allows_dogs = ?, links = ?, status = ?, created_at = ?, updated_at = ?, recreation_area_id = ?, user_id = ?
WHERE id = ?;

-- name: UpdateRecreationArea :exec
UPDATE recreation_area
SET
name = ?, status = ?, created_at = ?, updated_at = ?, user_id = ?
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE user
SET
name = ?, email = ?, password = ?, status = ?, created_at = ?, updated_at = ?
WHERE id = ?;

-- name: UpdateFeature :exec
UPDATE feature
SET
name = ?, status = ?, created_at = ?, updated_at = ?, user_id = ?
WHERE id = ?;

-- name: UpdateParkHasFeature :exec
UPDATE park_has_feature
SET
details = ?, status = ?, created_at = ?, updated_at = ?, park_id = ?, feature_id = ?
WHERE id = ?;

-- name: UpdateEvent :exec
UPDATE event
SET
name = ?, description = ?, main_image = ?, start_date = ?, end_date = ?, status = ?, created_at = ?, updated_at = ?, park_id = ?, user_id = ?
WHERE id = ?;

