package main



import (
	_ "code2/newsWeb/controllers"
	_ "code2/newsWeb/routers"
	"github.com/astaxie/beego"
	_ "code2/newsWeb/models"
)

func main() {
	//给视图函数建立映射
	beego.AddFuncMap("prePage",PrePage)
	beego.AddFuncMap("nextPage",NextPage)
	beego.Run()
}

func PrePage(pageNum int)int{
	if pageNum <= 1{
		return 1
	}
	return pageNum - 1
}

func NextPage(pageNum int,pageCount float64)int{
	if pageNum >= int(pageCount){
		return int(pageCount)
	}
	return pageNum + 1
}