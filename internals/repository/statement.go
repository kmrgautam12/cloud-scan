package repository

const (
	createUser                = `INSERT INTO "user"(username,name,password) VALUES($1,$2,$3)`
	checkUserExistsByUsername = `SELECT EXISTS(SELECT 1 FROM "user" WHERE username=$1)`
	getUserByUsername         = `SELECT username,name,password FROM "user" where username=$1`
)
