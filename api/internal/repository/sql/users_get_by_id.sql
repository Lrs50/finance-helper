SELECT id, username, name, password, email, phone_number, created_at
FROM users
WHERE id = $1;

