package data

type OrderForm struct {
	Oid string
	Uid string
	Sid int
	Aid string
	Ototalbooksprice float64
	Ototaltransprice float64
	Stransprice float64
	Sicon string
	Sname string
	Osubmittime string
	Opaytime string
	Oreceivetime string
	Ofinishedtime string
	Ostatus int
	BookInCartList []BookInCart
}

func FindOrderFormByOid(oid string)(orderform OrderForm,err error){
	statement := "select Oid,Orderform.Uid,Orderform.Sid,Aid,Ototalbooksprice,Ototaltransprice,date_format(Osubmittime,'%Y-%c-%d %H:%i:%s'),date_format(Opaytime,'%Y-%c-%d %H:%i:%s'),date_format" +
		"(Oreceivetime,'%Y-%c-%d %H:%i:%s'),date_format(Ofinishedtime,'%Y-%c-%d %H:%i:%s'),Ostatus,Shop.Sicon,Shop.Sname from Orderform,Shop where Orderform.Oid=? and Orderform.Sid=Shop.Sid limit 1";

	err=Db.QueryRow(statement, oid).Scan(&orderform.Oid,)
	return
}

func GetOrderPageCountByUid(uid string)(count int, err error){
	statement := "select count(*) from Orderform where Uid='"+ uid +"'";
	rowcount, _ := GetRowCount(statement)
	count = GetPageCount(rowcount,ORDER_PAGESIZE)
	return
}

func FindOrderFormsByUidByPage(username string, pageNow int)(orderforms []OrderForm, err error){
	return
}

func FindBookListInOrderList()  {
	
}

func FindBookListInOrder(){
	
}

func AddOrder()  {
	
}

func DeleteOrder(){

}
func ModifyOrderListStatusByOids()  {
	
}

func ModifyOrderStatusByOid(){

}