-- Delete duplicate entries
DELETE FROM groups 
WHERE id NOT IN (
    SELECT MIN(id)
    FROM groups
    GROUP BY name
); 