-- Remove static user groups
DELETE FROM user_groups
WHERE name = 'ADMIN'
OR name = 'STANDARD'
OR name = 'DONATOR';

-- Remove static workspace groups
DELETE FROM user_groups
WHERE name = 'OWNER'
OR name = 'ADMIN'
OR name = 'STANDARD';
