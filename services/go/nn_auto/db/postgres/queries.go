package postgres

const (
	createUserQuery = `
	INSERT INTO users (username, password_hash, email) VALUES ($1, $2, $3) RETURNING id;
`
	setRefreshQuery = `
	UPDATE users
	SET refresh_token = $1
	WHERE id = $2
	`
	getUserByLoginQuery = `
	SELECT id, username, password_hash, email, refresh_token 
	FROM users
	WHERE username = $1;
	`
	getUserByRefreshQuery = `
	SELECT id, username, password_hash, email, refresh_token 
	FROM users
	WHERE refresh_token = $1;
	`
)
