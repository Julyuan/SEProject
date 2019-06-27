package data

type Shop struct {
	Sid 		int
	Sname 		string
	Uid   		string
	Sicon		string
	Ssummary	string
	Sactivity	string
	Stransprice	float64
}

func (shop *Shop) AddShop(err error){
	statement := "insert into shop (Uid) values (?)"

	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(shop.Uid).Scan()
	return
}

func (shop *Shop) FindShopBySid(sid int){
	statement := "select * from shop where Sid = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		return
	}

	defer stmt.Close()
	Db.QueryRow(statement,shop.Uid).
		Scan(&shop.Uid, &shop.Sname, &shop.Uid, &shop.Sicon,
			&shop.Ssummary, &shop.Sactivity, &shop.Stransprice)
}

func GetSidByUid(uid string)(sid int, err error){
	statement := "select Sid from shop where Uid = ?";
	err = Db.QueryRow(statement, uid).Scan(&sid)
	return
}

func UpdateShop(sicon string)(err error){
	statement := "update shop set Sname = ? , Sicon = ? , Ssummary = ? ,Sactivity = ? ,Stransprice = ? where Sid = ?"

	_, err = Db.Exec(statement)
	return
}


func GetSiconBySid(sid int)(sicon string, err error){
	statement := "select Sicon from Shop where Sid = ?";
	err = Db.QueryRow(statement, sid).Scan(&sicon)

	return
}
