package data
import "C"

type BookComment struct {
	Uavatar string
	Unickname string
	Tcomment string
	Tcommenttime string
}


func GetBookCommentByIdAndPage(bid string, pageNow int)(bookComments []BookComment, err error){
	statement := "select Uavatar,Unickname,Tcomment,date_format(Tcommenttime,'%Y-%c-%d %H:%i:%s') " +
		"from Transaction,User where Bid= $1 and Transaction.Uid=User.Uid and Tcomment != '' " +
		"order by Tcommenttime DESC limit $2 , $3";
	rows,err := Db.Query(statement, bid, BookCommentPageSize*(pageNow-1), BookCommentPageSize)

	for rows.Next(){
		bookComment := BookComment{}
		if err = rows.Scan(&bookComment.Uavatar, &bookComment.Unickname, &bookComment.Tcomment, &bookComment.Tcommenttime); err != nil {
			return
		}
		bookComments = append(bookComments, bookComment)
	}

	rows.Close()
	return
}


