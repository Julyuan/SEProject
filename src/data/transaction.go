package data

type Transaction struct {
	Tid 		int
	Bid 		string
	Oid 		string
	Uid 		string
	Tstatus		int
	Tboughtnum	int
	Tmark		int
	Tcomment	string
	Tsubmittime	string
	Tpaytime	string
	Treceivetime string
	Tcommenttime	string
}


func NewTransaction() *Transaction{
	transaction := new(Transaction)
	transaction.Tid = 0
	transaction.Bid = ""
	transaction.Oid = ""
	transaction.Tstatus = TRADESTATUS_WAITPAY
	transaction.Tboughtnum = 0
	transaction.Tmark = 0
	transaction.Tcomment = ""
	transaction.Tsubmittime = ""
	transaction.Tpaytime = ""
	transaction.Treceivetime = ""
	transaction.Tcommenttime = ""
	return transaction
}

func AddTrade(){

}

func ModifyTradeStatusByOids()  {

}

func ModifyTradeStatusByOid(){
	
}

func UpdateTransactionByOidByBid(){

}

func FindTransactionStatusListByOid()  {
	
}