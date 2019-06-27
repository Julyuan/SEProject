package data
import "C"

type BookTradeRecord struct {
	Uavatar string
	Unickname string
	Tboughtnum int
	Treceivetime string
}

func NewBookTradeRecord() *BookTradeRecord{
	booktraderecord := new(BookTradeRecord)
	return booktraderecord
}

func GetBookTradeRecordByIdAndPage(bid string, pageNow int)(bookTradeRecords []BookTradeRecord, err error){
	statement := "select Uavatar,Unickname,Tboughtnum,date_format(Treceivetime,'%Y-%c-%d %H:%i:%s') from transaction," +
		"user where Bid=$1 and Tstatus >= $2 and Transaction.Uid = User.Uid order by Treceivetime DESC limit $3, $4";
	rows,err := Db.Query(statement, bid, TRADESTATUS_WAITCOMMENT, BookTradeRecordPageSize*(pageNow-1), BookTradeRecordPageSize)

	for rows.Next(){
		bookTradeRecord := BookTradeRecord{}
		if err = rows.Scan(&bookTradeRecord.Uavatar, &bookTradeRecord.Unickname, &bookTradeRecord.Tboughtnum, &bookTradeRecord.Treceivetime); err != nil {
			return
		}
		bookTradeRecords = append(bookTradeRecords, bookTradeRecord)
	}

	rows.Close()
	return
}