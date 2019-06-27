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
	Uid     string
	CreatedAt  time.Time
}

func (user *User) CreateSession()(session Session, err error){
	statement := "insert into sessions (uuid, Uemail, Uid, created_at) "+
		"values (?, ?, ?, ?)"

	stmt, err := Db.Prepare(statement)

	if err != nil{
		return
	}

	defer  stmt.Close()
	uuid := createUUID()
	_,err = stmt.Exec(uuid, user.Uemail, user.Uid, time.Now())

	statement1 := "select id, uuid, Uemail, Uid, created_at from sessions where uuid=? "
	stmt1,err := Db.Prepare(statement1)

	if err != nil{
		return
	}
	defer stmt1.Close()

	err	= stmt1.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Uid, &session.Uemail, &session.CreatedAt)
	return
}

func (user *User) Session() (session Session, err error){
	session = Session{}
	err = Db.QueryRow("SELECT id, uuid, Uemail, Uid, created_at FROM sessions WHERE Uid = $1", user.Uid).
		Scan(&session.Id, &session.Uuid, &session.Uemail, &session.Uid, &session.CreatedAt)
	return
}

func (session *Session) Check()(valid bool, err error){
	fmt.Println("call check function")
	//fmt.Println("uuid is equal to",session.Uuid)
	//fmt.Println("SELECT id, uuid, Uemail, Uid, created_at FROM sessions WHERE uuid = "+session.Uuid)
	err = Db.QueryRow("SELECT id, uuid, Uemail, Uid, created_at FROM sessions WHERE uuid = "+"\""+session.Uuid+"\"").
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
	err = Db.QueryRow("SELECT Uid, Urealname, Uemail FROM user WHERE Uid = ?", session.Uid).
		Scan(&user.Uid, &user.Urealname, &user.Uemail)
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
	statement := "insert into user (Uid, Upassword) values (?, ?)"
	fmt.Print("user.Uid ",user.Uid," user.Upassword ",user.Upassword,"\n")
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Uid,user.Upassword)
	return
}


func (user *User)Delete() (err error){
	statement := "delete from user where Uid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uid)
	return
}

// Update user information in the database
func (user *User) Update() (err error) {
	statement := "update user set Ureadlname = ?, Uemail = ? where Uid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uid, user.Urealname, user.Uemail)
	return
}

func UserDeleteAll() (err error){
	statement := "delete from user"
	_, err = Db.Exec(statement)
	return
}

func Users() (users []User, err error){
	rows, err := Db.Query("SELECT id, uuid, name, email, password, created_at FROM user")
	if err != nil{
		return
	}

	for rows.Next(){
		user := User{}
		if err = rows.Scan(&user.Uid, &user.Urealname, &user.Uemail, &user.Upassword); err != nil {
			return
		}
		users = append(users, user)
	}

	rows.Close()
	return
}

func UserByEmail(email string) (user User, err error){
	fmt.Println("Call UserByEmail email = "+email)
	statement := "SELECT Uid,  Urealname, Uemail, Upassword FROM "+
		"users WHERE Uemail = $1"
	user = User{}
	err = Db.QueryRow(statement, email).Scan(&user.Uid, &user.Urealname,
		&user.Uemail, &user.Upassword)
	fmt.Println("username = "+user.Urealname)
	return
}

// Get a single user given the UUID
func UserByUid(Uid string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT Uid, Urealname, Uemail, Upassword FROM user WHERE Uid = ?", Uid).
		Scan(&user.Uid, &user.Urealname, &user.Uemail, &user.Upassword)
	return
}