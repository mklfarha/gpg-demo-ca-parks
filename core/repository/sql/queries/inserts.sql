
-- name: InsertPark :execresult
INSERT INTO park
(id,name,main_image,phone,hours,allows_dogs,links,status,created_at,updated_at,recreation_area_id,user_id)
VALUES
(?,?,?,?,?,?,?,?,?,?,?,?);

-- name: InsertRecreationArea :execresult
INSERT INTO recreation_area
(id,name,status,created_at,updated_at,user_id)
VALUES
(?,?,?,?,?,?);

-- name: InsertUser :execresult
INSERT INTO user
(id,name,email,password,status,created_at,updated_at)
VALUES
(?,?,?,?,?,?,?);

-- name: InsertFeature :execresult
INSERT INTO feature
(id,name,status,created_at,updated_at,user_id)
VALUES
(?,?,?,?,?,?);

-- name: InsertParkHasFeature :execresult
INSERT INTO park_has_feature
(id,details,status,created_at,updated_at,park_id,feature_id)
VALUES
(?,?,?,?,?,?,?);

-- name: InsertEvent :execresult
INSERT INTO event
(id,name,description,main_image,start_date,end_date,status,created_at,updated_at,park_id,user_id)
VALUES
(?,?,?,?,?,?,?,?,?,?,?);
