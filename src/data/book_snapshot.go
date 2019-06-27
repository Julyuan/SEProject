package data

import (
	"fmt"
)

type BookSnapshot struct {
	Bid string
	Bimage string
	Bprice float64
	Bname string
	Sname string
	Bsalednum int
	Bcommentnum int
}

func GetRandom8Books()(booksnapshots []BookSnapshot, err error){
	query_sql := "select Bid,Bimage,Bprice,Bname,Sname,Bsalednum,Bcommentnum from book,shop where book.Sid=shop.Sid order by rand() limit 8"
	rows, err := Db.Query(query_sql)
	if err != nil{
		fmt.Println("get random book出错")
		return
	}

	for rows.Next(){
		booksnapshot := BookSnapshot{}
		if err = rows.Scan(&booksnapshot.Bid, &booksnapshot.Bimage, &booksnapshot.Bprice, &booksnapshot.Bname,
			&booksnapshot.Sname, &booksnapshot.Bsalednum, &booksnapshot.Bcommentnum); err != nil {
			return
		}
		booksnapshots = append(booksnapshots, booksnapshot)
	}

	rows.Close()
	return
}

func GetBookSnapshotsByKey(key string, page_now int, refer string)(booksnapshots []BookSnapshot, err error){
	query_sql := "select Bid,Bimage,Bprice,Bname,Sname,Bsalednum,Bcommentnum from Book,Shop where Book.Sid=Shop.Sid and (Bname like '%"+
				key + "%' or Bauthor like '%" + key +"%' or Bpublisher like '%"+ key +"%' or Sname like '%" + key +"%') ";
	rows, err := Db.Query(query_sql)
	if err != nil{
		return
	}

	for rows.Next(){
		booksnapshot := BookSnapshot{}
		if err = rows.Scan(&booksnapshot.Bid, &booksnapshot.Bimage, &booksnapshot.Bprice, &booksnapshot.Bname,
			&booksnapshot.Sname, &booksnapshot.Bsalednum, &booksnapshot.Bcommentnum); err != nil {
			return
		}
		booksnapshots = append(booksnapshots, booksnapshot)
	}

	rows.Close()
	return
}

func GetBookSnapshotsByType(booktype string, pageNow int, refer string)(booksnapshots []BookSnapshot, err error){
	statement := "select Bid,Bimage,Bprice,Bname,Sname,Bsalednum,Bcommentnum from Book,Shop where Btype = ? and Book.Sid=Shop.Sid order by ? limit ?,?";
	rows, err := Db.Query(statement,booktype,refer,BookSnapshotPageSize*(pageNow-1),BookSnapshotPageSize)
	if err != nil{
		return
	}

	for rows.Next(){
		booksnapshot := BookSnapshot{}
		if err = rows.Scan(&booksnapshot.Bid, &booksnapshot.Bimage, &booksnapshot.Bprice, &booksnapshot.Bname,
			&booksnapshot.Sname, &booksnapshot.Bsalednum, &booksnapshot.Bcommentnum); err != nil {
			return
		}
		booksnapshots = append(booksnapshots, booksnapshot)
	}

	rows.Close()
	return
}

func SelectBookSnapshotBeanResultSet(sql string)(booksnapshots []BookSnapshot, err error){
	fmt.Println("Call SelectBookSnapshotBeanResultSet sql = " + sql)
	rows, err := Db.Query(sql)

	for rows.Next(){
		booksnapshot := BookSnapshot{}
		if err = rows.Scan(&booksnapshot.Bid, &booksnapshot.Bimage, &booksnapshot.Bprice, &booksnapshot.Bname,
			&booksnapshot.Sname, &booksnapshot.Bsalednum, &booksnapshot.Bcommentnum); err != nil {
			return
		}
		booksnapshots = append(booksnapshots, booksnapshot)
	}

	rows.Close()
	return
}
