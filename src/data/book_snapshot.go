package data

import "strconv"

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
	query_sql := "select Bid,Bimage,Bprice,Bname,Sname,Bsalednum,Bcommentnum from Book,Shop where Book.Sid=Shop.Sid order by rand() limit 8"
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

func GetSnapshotsByKey(key string, page_now int, refer string)(booksnapshots []BookSnapshot, err error){
	query_sql := "select Bid,Bimage,Bprice,Bname,Sname,Bsalednum,Bcommentnum from Book,Shop where Book.Sid=Shop.Sid and (Bname like '%"+
				key + "%' or Bauthor like '%" + key +"%' or Bpublisher like '%"+ key +"%' or Sname like '%" + key +"%') order by "+ refer +" limit " + strconv.Itoa(bookSnapshotPageSize * (page_now - 1)) + "," + strconv.Itoa(bookSnapshotPageSize);
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


