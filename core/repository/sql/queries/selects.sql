


-- name: FetchParkByID :many
SELECT * FROM park
WHERE 
     id = ?  
    ;

-- name: FetchParkByStatus :many
SELECT * FROM park
WHERE 
     status = ?  
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaID :many
SELECT * FROM park
WHERE 
     recreation_area_id = ?  
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaIDAndStatus :many
SELECT * FROM park
WHERE 
     recreation_area_id = ? AND status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchParkByStatusOrderedByCreatedAtASC :many
SELECT * FROM park
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkByStatusOrderedByCreatedAtDESC :many
SELECT * FROM park
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkByStatusOrderedByUpdatedAtASC :many
SELECT * FROM park
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM park
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkByRecreationAreaIDOrderedByCreatedAtASC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaIDOrderedByCreatedAtDESC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkByRecreationAreaIDOrderedByUpdatedAtASC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaIDOrderedByUpdatedAtDESC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtASC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaIDAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkByRecreationAreaIDAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM park
WHERE 
     recreation_area_id = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchPark :many
SELECT * FROM park
WHERE 
    name like ? 
    LIMIT ?, ?;








-- name: FetchRecreationAreaByID :many
SELECT * FROM recreation_area
WHERE 
     id = ?  
    ;

-- name: FetchRecreationAreaByStatus :many
SELECT * FROM recreation_area
WHERE 
     status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchRecreationAreaByStatusOrderedByCreatedAtASC :many
SELECT * FROM recreation_area
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchRecreationAreaByStatusOrderedByCreatedAtDESC :many
SELECT * FROM recreation_area
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchRecreationAreaByStatusOrderedByUpdatedAtASC :many
SELECT * FROM recreation_area
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchRecreationAreaByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM recreation_area
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchRecreationArea :many
SELECT * FROM recreation_area
WHERE 
    name like ? 
    LIMIT ?, ?;








-- name: FetchUserByID :many
SELECT * FROM user
WHERE 
     id = ?  
    ;

-- name: FetchUserByEmail :many
SELECT * FROM user
WHERE 
     email = ?  
    LIMIT ?, ?;

-- name: FetchUserByStatus :many
SELECT * FROM user
WHERE 
     status = ?  
    LIMIT ?, ?;

-- name: FetchUserByEmailAndStatus :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchUserByEmailOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByEmailOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchUserByEmailAndStatusOrderedByCreatedAtASC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchUserByEmailAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchUserByEmailAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM user
WHERE 
     email = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchFeatureByID :many
SELECT * FROM feature
WHERE 
     id = ?  
    ;

-- name: FetchFeatureByStatus :many
SELECT * FROM feature
WHERE 
     status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchFeatureByStatusOrderedByCreatedAtASC :many
SELECT * FROM feature
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchFeatureByStatusOrderedByCreatedAtDESC :many
SELECT * FROM feature
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchFeatureByStatusOrderedByUpdatedAtASC :many
SELECT * FROM feature
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchFeatureByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM feature
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    



-- name: SearchFeature :many
SELECT * FROM feature
WHERE 
    name like ? 
    LIMIT ?, ?;








-- name: FetchParkHasFeatureByID :many
SELECT * FROM park_has_feature
WHERE 
     id = ?  
    ;

-- name: FetchParkHasFeatureByStatus :many
SELECT * FROM park_has_feature
WHERE 
     status = ?  
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkID :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ?  
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkIDAndStatus :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ? AND status = ?  
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureID :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ?  
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureIDAndStatus :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ? AND status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchParkHasFeatureByStatusOrderedByCreatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByStatusOrderedByCreatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkHasFeatureByStatusOrderedByUpdatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkHasFeatureByParkIDOrderedByCreatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkIDOrderedByCreatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkHasFeatureByParkIDOrderedByUpdatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkIDOrderedByUpdatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkIDAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByParkIDAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     park_id = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkHasFeatureByFeatureIDOrderedByCreatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureIDOrderedByCreatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureIDOrderedByUpdatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureIDAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchParkHasFeatureByFeatureIDAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM park_has_feature
WHERE 
     feature_id = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    










-- name: FetchEventByID :many
SELECT * FROM event
WHERE 
     id = ?  
    ;

-- name: FetchEventByStatus :many
SELECT * FROM event
WHERE 
     status = ?  
    LIMIT ?, ?;

-- name: FetchEventByParkID :many
SELECT * FROM event
WHERE 
     park_id = ?  
    LIMIT ?, ?;

-- name: FetchEventByParkIDAndStatus :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    LIMIT ?, ?;



    

    
    
-- name: FetchEventByStatusOrderedByStartDateASC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchEventByStatusOrderedByStartDateDESC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByStatusOrderedByEndDateASC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY end_date ASC
    LIMIT ?, ?;

-- name: FetchEventByStatusOrderedByEndDateDESC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY end_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByStatusOrderedByCreatedAtASC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchEventByStatusOrderedByCreatedAtDESC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchEventByStatusOrderedByUpdatedAtASC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchEventByStatusOrderedByUpdatedAtDESC :many
SELECT * FROM event
WHERE 
     status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchEventByParkIDOrderedByStartDateASC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDOrderedByStartDateDESC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDOrderedByEndDateASC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY end_date ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDOrderedByEndDateDESC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY end_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDOrderedByCreatedAtASC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDOrderedByCreatedAtDESC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDOrderedByUpdatedAtASC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDOrderedByUpdatedAtDESC :many
SELECT * FROM event
WHERE 
     park_id = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    

    
    
-- name: FetchEventByParkIDAndStatusOrderedByStartDateASC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY start_date ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDAndStatusOrderedByStartDateDESC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY start_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDAndStatusOrderedByEndDateASC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY end_date ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDAndStatusOrderedByEndDateDESC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY end_date DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDAndStatusOrderedByCreatedAtASC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY created_at ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDAndStatusOrderedByCreatedAtDESC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY created_at DESC
    LIMIT ?, ?;

        
-- name: FetchEventByParkIDAndStatusOrderedByUpdatedAtASC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY updated_at ASC
    LIMIT ?, ?;

-- name: FetchEventByParkIDAndStatusOrderedByUpdatedAtDESC :many
SELECT * FROM event
WHERE 
     park_id = ? AND status = ?  
    ORDER BY updated_at DESC
    LIMIT ?, ?;

        
    







