package data

import "C"
import "fmt"

type BookDetail struct {
	Bid string
	Bimage string
	Bprice float64
	Bname string
	Sid int
	Sname string
	Sicon string
	Bsalednum int
	Bcommentnum int
	Bauthor string
	Bpublisher string
	Bpublishdate string
	Boriprice float64
	Stransprice float64
	Sactivity string
	Tmark_AVG int
	Bversion int
	Bpagenum int
	Bwordnum int
	Bprintdate string
	Bsize string
	Bpaperstyle string
	Bprintnum int
	Bpackage string
	Bisbn string
	Bcontentsummary string
	Bauthorsummary string
	Bmediacomment string
	Btastecontent string
	Bstocknum int
	Buploaddate string
	Btype string
}


func AddBook(book BookDetail)(err error){
	statement := "insert into Book values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)";
	_, err = Db.Exec(statement,book.Bid, book.Bimage, book.Bprice, book.Bname, book.Sid, book.Bauthor,
		book.Bpublisher, book.Bpublishdate, book.Bsalednum, book.Bcommentnum, book.Boriprice, book.Bversion,
		book.Bwordnum, book.Bprintdate, book.Bsize, book.Bpaperstyle, book.Bprintnum, book.Bpackage,
		book.Bisbn, book.Bcontentsummary, book.Bauthorsummary, book.Bmediacomment, book.Btastecontent,
		)
	return
}

func GetBookDetailById(bid string)(bookDetail BookDetail, err error)  {
	fmt.Println("Call GetBookDetailById bid = " + bid)
	statement := "select book.Bid, Bimage, Bprice, Bname, shop.Sid, Sname, Bsalednum, Bcommentnum, Bauthor, Bpublisher, Bpublishdate, Boriprice, shop.Stransprice, Sactivity, Bversion, Bpagenum, Bwordnum, Bprintdate, Bsize, Bpaperstyle, Bprintnum, Bpackage, " +
		"Bisbn, Bcontentsummary, Bauthorsummary, Bmediacomment, Btastecontent,shop.Sicon from book,shop,transaction where book.Bid = ? and book.Bid = transaction.Bid and Tstatus >= ? and book.Sid=shop.Sid group by book.Bid limit 1";
	bookDetail = BookDetail{}
	err = Db.QueryRow(statement, bid, 1).Scan(&bookDetail.Bid, &bookDetail.Bimage, &bookDetail.Bprice, &bookDetail.Bname,
		&bookDetail.Sid, &bookDetail.Sname, &bookDetail.Bsalednum, &bookDetail.Bcommentnum, &bookDetail.Bauthor, &bookDetail.Bpublisher,
		&bookDetail.Bpublishdate,&bookDetail.Boriprice, &bookDetail.Stransprice, &bookDetail.Sactivity, &bookDetail.Bversion,
		&bookDetail.Bpagenum, &bookDetail.Bwordnum, &bookDetail.Bprintdate, &bookDetail.Bsize, &bookDetail.Bpaperstyle, &bookDetail.Bprintnum, &bookDetail.Bpackage,
		&bookDetail.Bisbn, &bookDetail.Bcontentsummary, &bookDetail.Bauthorsummary, &bookDetail.Bmediacomment, &bookDetail.Btastecontent, &bookDetail.Sicon)
	//fmt.Println("username = "+user.Urealname)
	fmt.Println(err)
	return
}

