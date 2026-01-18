INSERT INTO users (username, name, password, email, phone_number)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, name, password, email, phone_number, created_at;