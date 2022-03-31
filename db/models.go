package db

const USERS_TABLE = `
	CREATE TABLE IF NOT EXISTS users(
		userId INTEGER PRIMARY KEY AUTO_INCREMENT,
		firstName VARCHAR(20),
		lastName VARCHAR(20),
		otherName VARCHAR(20) DEFAULT NULL,
		gender VARCHAR(1),
		address VARCHAR(100),
		email VARCHAR(50) UNIQUE,
		phoneNum VARCHAR(20),
		otherNum VARCHAR(20) DEFAULT NULL,
		kinName VARCHAR(60),
		kinNumber VARCHAR(20),
		kinRelationship VARCHAR(50),
		passwordHash VARCHAR(200)
	);

`

const ACCOUNTS_TABLE = `
	CREATE TABLE IF NOT EXISTS accounts(
		acctId INTEGER PRIMARY KEY AUTO_INCREMENT,
		currentBal INTEGER DEFAULT NULL,
		userId INTEGER DEFAULT NULL,
		acctType VARCHAR(1),
		acctNum INTEGER,
		FOREIGN KEY(userId) REFERENCES users(userId) ON DELETE SET NULL
	)

`

const TRANSACTIONS_TABLE = `
	CREATE TABLE IF NOT EXISTS transactions(
		txnId INTEGER PRIMARY KEY AUTO_INCREMENT,
		senderId INTEGER,
		receiverId INTEGER,
		amount INTEGER,
		txnTime TIMESTAMP,
		FOREIGN KEY(senderId) REFERENCES accounts(acctId) ON DELETE SET NULL,
		FOREIGN KEY(receiverId) REFERENCES accounts(acctId) ON DELETE SET NULL
	)

`

type UserStruct struct {
	UserId          int    `bson:"userId" json:"userId"`
	FirstName       string `bson:"firstName" json:"firstName"`
	LastName        string `bson:"lastName" json:"lastName"`
	OtherName       string `bson:"otherName" json:"otherName"`
	Email           string `bson:"email" json:"email"`
	Gender          string `bson:"gender" json:"gender"`
	PhoneNum        string `bson:"phoneNum" json:"phoneNum"`
	OtherNum        string `bson:"otherNum,omitempty" json:"otherNum,omitempty"`
	Address         string `bson:"address" json:"address"`
	KinName         string `bson:"kinName" json:"kinName"`
	KinNumber       string `bson:"kinNumber" json:"kinNumber"`
	KinRelationship string `bson:"kinRelationship" json:"kinRelationship"`
	PasswordHash    string `bson:"passwordHash,omitempty" json:"passwordHash,omitempty"`
}

type AcctStruct struct {
	AcctId     int    `bson:"acctId" json:"acctId"`
	UserId     int    `bson:"userId" json:"userId"`
	CurrentBal int    `bson:"currentBal" json:"currentBal"`
	AcctType   string `bson:"acctType" json:"acctType"`
	AcctNum    int    `bson:"acctNum" json:"acctNum"`
}

type TxnStruct struct {
	TxnId      int    `bson:"txnId" json:"txnId"`
	SenderId   int    `bson:"senderId" json:"senderId"`
	ReceiverId int    `bson:"receiverId" json:"receiverId"`
	Amount     int    `bson:"amount" json:"amount"`
	TxnTime    string `bson:"txnTime" json:"txnTime"`
}
