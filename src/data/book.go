package data

import "fmt"

func GetMaxBidBySid(sid int)(bid int, err error)  {
	fmt.Println("Call GetMaxBidBySid sid = " + string(sid))
	statement := "select max(Bid) from book where Sid = ?";
	err = Db.QueryRow(statement, sid).Scan(&bid)
	//fmt.Println("username = "+user.Urealname)
	return
}


func IncSaledNum(oid string)(err error){
	fmt.Println("Call IncSaledNum oid = " + string(oid))
	statement := "update Book,Transaction set Bsalednum = Bsalednum + 1 where Book.Bid in " +
		"(select Transaction.Bid from Transaction where Transaction.Oid = ?)";
	_, err = Db.Query(statement, oid)
	return
}

func IncCommentNum(bid string)(err error){
	fmt.Println("Call IncCommentNum bid = " + bid)
	statement := "update book set Bcommentnum = Bcommentnum + 1 where Bid = ?";
	_, err = Db.Exec(statement, bid)
	return
}


func GetRowCount(sql string)(count int, err error){
	err = Db.QueryRow(sql).Scan(&count)
	return
}

func GetPageCount(rowCount int, pageSize int)(int){
	if rowCount % pageSize== 0 {
		return rowCount / pageSize;
	} else {
		return rowCount/pageSize + 1;
	}
}

func GetBookSnapshotsPageCountByKey(key string)(pageCount int, err error){
	statement := "select count(*) from Book,Shop where Book.Sid = Shop.Sid and (Bname like '%"+key+"%' or Bauthor like '%" + key +
		"%' or Bpublisher like '%"+ key +"%' or Sname like '%"+ key +"%')";
	err = Db.QueryRow(statement).Scan(&pageCount)
	return
}

func GetBookSnapshotsPageCountByType(booktype string)(pageCount int, err error){
	statement := "select count(*) from Book where Btype = '"+ booktype +"'";
	err = Db.QueryRow(statement).Scan(&pageCount)
	return
}

