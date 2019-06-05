package data

import (
	"fmt"
	"time"
)

type User struct{
	Uid    			string
	Upassword  		string
	Uavatar  		string
	Unickname 		string
	Urealname		string
	Usex			string
	Uaddress		string
	Uphone			string
	Uemail          string
	Ubirthday		string
}

type Session struct {
	Id     int
	Uuid   string
	Uemail  string
	Uid int
	CreatedAt  time.Time
}

func (user *User) CreateSession()(session Session, err error){
	statement := "insert into sessions (uuid, Uemail, Uid, created_at) "+
		"values ($1, $2, $3, $4) returning id, uuid, Uemail, Uid, created_at"

	stmt, err := Db.Prepare(statement)

	if err != nil{
		return
	}

	defer  stmt.Close()

	err = stmt.QueryRow(createUUID(), user.Uemail, user.Uid, time.Now()).
		Scan(&session.Id, &session.Uuid, &session.Uid, &session.Uemail, &session.CreatedAt)

	return
}

func (user *User) Session() (session Session, err error){
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1", user.Uid).
		Scan(&session.Id, &session.Uuid, &session.Uemail, &session.Uid, &session.CreatedAt)
	return
}

func (session *Session) Check()(valid bool, err error){
	fmt.Println("call check function")
	fmt.Println("uuid is equal to",session.Uuid)
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Uemail, &session.Uid, &session.CreatedAt)

	if err != nil {
		fmt.Println("valid is false")
		valid = false
		return
	}

	if session.Id != 0{
		fmt.Println("valid is true")
		valid = true

	}

	return
}

func(session *Session) DeleteByUUID()(err error){
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)

	if err != nil{
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

func (session *Session)User()(user User, err error){
	user = User{}
	err = Db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", session.Uid).
		Scan(&user.Uid, &user., &user.Uemail, &user.CreatedAt)
	if(err != nil){
		fmt.Println(err.Error())
	}
	return
}

func SessionDeleteAll()(err error){
	statement := "delete from sessions"
	_, err = Db.Exec(statement)
	return
}


func (user *User) Create() (err error){

	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, "+
		"$3, $4, $5) returning id, uuid, created_at"

	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(createUUID(),user.Name,user.Email, Encrypt(user.Password), time.Now()).Scan(&user.Id, &user.Uuid, &user.CreatedAt)
	return
}


func (user *User)Delete() (err error){
	statement := "delete from users where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update users set name = $2, email = $3 where id = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id, user.Name, user.Email)
	return
}

func UserDeleteAll() (err error){
	statement := "delete from users"
	_, err = Db.Exec(statement)
	return
}

func Users() (users []User, err error){
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil{
		return
	}

	for rows.Next(){
		user := User{}
		if err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return
		}
		users = append(users, user)
	}

	rows.Close()
	return
}

func UserByEmail(email string) (user User, err error){
	fmt.Println("Call UserByEmail email = "+email)
	statement := "SELECT id, uuid, name, email, password, created_at FROM "+
		"users WHERE email = $1"
	user = User{}
	err = Db.QueryRow(statement, email).Scan(&user.Id, &user.Uuid, &user.Name,
		&user.Email, &user.Password, &user.CreatedAt)
	fmt.Println("username = "+user.Name)
	return
}

// Get a single user given the UUID
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE uuid = $1", uuid).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}