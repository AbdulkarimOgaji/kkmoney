package api

const (
	getUsers = `
	SELECT userId, firstName, lastName,
		IFNULL(otherName, ""), email, phoneNum,
		IFNULL(otherNum, ""), gender, address,
		kinName, kinNumber, kinRelationship, createdTime
	FROM users
	`
	getUserById = `
	SELECT userId, firstName, lastName, IFNULL(otherName, ""),
		email, phoneNum, IFNULL(otherNum, ""), gender, address,
		kinName, kinNumber, kinRelationship, createdTime
	FROM users
	WHERE userId = ?
	`
	updateUser = `
	UPDATE users
	SET firstName = NULLIF(?, ""),
		lastName = NULLIF(?, ""),
		otherName = NULLIF(?, ""),
		email = NULLIF(?, ""),
		phoneNum = NULLIF(?, ""),
		otherNum = NULLIF(?, ""),
		gender = NULLIF(?, ""),
		address = NULLIF(?, ""),
		kinName = NULLIF(?, ""),
		kinNumber = NULLIF(?, ""),
		kinRelationship = NULLIF(?, "")
	WHERE userId = ?
	`

	createUser = `
	INSERT INTO users(
		firstName, lastName, otherName, email, phoneNum,
		otherNum, gender, address, kinName, kinNumber, kinRelationship, createdTime, passwordHash
			) 
		values(NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""),
		NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""),  NULLIF(?, ""))
	`

	deleteUser = `
	DELETE FROM users
	WHERE userId = ?
	`
)
const (
	getAccts = `
	SELECT * FROM accounts
	`

	getUserAccts = `
	SELECT * FROM accounts
	WHERE userId = ?
`
	getAcctById = `
	SELECT * FROM accounts
	WHERE acctId = ?
`
	updateAcct = `
	UPDATE accounts
	SET acctType = NULLIF(?, "")
	WHERE acctId = ?
	`

	createAcct = `
	INSERT INTO accounts(userId, acctType, acctNum, createdTime)
	values(NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""), NULLIF(?, ""))
	`

	deleteAcct = `
	DELETE FROM accounts
	WHERE acctId = ?
`
)
