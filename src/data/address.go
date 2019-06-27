package data

import "fmt"

type Address struct {
	Aid string
	Uid string
	Areceivername string
	Aaddress string
	Acode string
	Aphone string
	Afixphone string
	Aprovince string
	Acity string
	Atown string
	Acheck int
}


func GetAddressByAid(aid string) (address Address, err error){
	fmt.Println("Call GetAddressByAid aid = " + aid)
	statement := "SELECT * FROM "+
		"address WHERE Aid = $1"
	address = Address{}
	err = Db.QueryRow(statement, aid).Scan(&address.Aid, &address.Aid, &address.Areceivername, &address.Aaddress,
		&address.Acode, &address.Aphone, &address.Afixphone, &address.Aprovince, &address.Acity, &address.Atown, &address.Acheck)
	//fmt.Println("username = "+user.Urealname)
	return
}


func GetAddressesByUid(uid string) (addresses []Address, err error){
	fmt.Println("Call GetAddressByAid uid = " + uid)
	statement := "select * from address where Uid = $1"

	rows,err := Db.Query(statement, uid)

	for rows.Next(){
		address := Address{}
		if err = rows.Scan(&address.Aid, &address.Uid, &address.Areceivername, &address.Aaddress,
			&address.Acode, &address.Aphone, &address.Afixphone, &address.Aprovince, &address.Acity,
			&address.Atown, &address.Acheck); err != nil {
			return
		}
		addresses = append(addresses, address)
	}

	rows.Close()
	return
}

func AddAddress(address Address)(err error){
	statement := "insert into address (Aid, Uid, Areceivername, Aaddress, Acode, Aphone, " +
		"Afixphone, Aprovince, Acity, Atown, Acheck) values (?,?,?,?,?,?,?,?,?,?,?)"

	_, err = Db.Exec(statement, address.Aid, address.Uid, address.Areceivername, address.Acode,
		address.Aphone, address.Afixphone, address.Aprovince, address.Acity, address.Atown, address.Acheck)

	return
}

func ModifyAddress(address Address)(index int, err error){
	return
}

func GetAddressCount(uid string)(count int, err error){
	count = MaxAddressNum

	statement := "select count(*) from address where Uid = ?"
	err = Db.QueryRow(statement, uid).Scan(&count)
	return
}

func GetMaxAidByUid(uid string)(maxaid string, err error){
	statement := "select max(Aid) from address where Uid = ?"
	err = Db.QueryRow(statement,uid).Scan(&maxaid)
	return
}

func DeleteAddress(username string, aid string)(newCheckedAid string, err error)  {
	statement := "delete from Address where Aid=?"
	_, err = Db.Exec(statement, aid)

	if err!=nil{
		return
	}
	statement = "select Aid from Address where Uid=? limit 1"
	err = Db.QueryRow(statement,username).Scan(&newCheckedAid)
	return
}

func CheckAddress(username string, aid string)(err error){
	statement := "update Address set Acheck=? where Uid=? and Acheck=?"
	_, err = Db.Exec(statement, AddressUnchecked, username, AddressChecked)

	if err!=nil{
		return
	}
	statement = "update Address set Acheck=? where Aid=?"
	_, err = Db.Exec(statement, AddressChecked, aid)

	return
}