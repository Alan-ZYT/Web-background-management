package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"code2/newsWeb/models"
	"encoding/base64"
)

type UserController struct {
	beego.Controller
}

//展示注册页面
func(this*UserController)ShowRegister(){
	this.TplName = "register.html"
}

//处理注册业务
func(this*UserController)HandleRegister(){
	//1.获取数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userName == "" || pwd == "" {
		beego.Error("传输数据不完整")
		this.TplName = "register.html"
		return
	}
	//处理数据
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	user.Pwd = pwd
	id,err := o.Insert(&user)
	if err != nil {
		beego.Error("用户注册失败")
		this.TplName = "register.html"
		return
	}
	beego.Info(id)
	//返回数据
	this.Redirect("/login",302)
}

//展示登录页面
func(this*UserController)ShowLogin(){
	//获取cookie数据，如果获取查到了，说明上一次记住用户名，不然的话，不记住用户名
	userName := this.Ctx.GetCookie("userName")
	//解密
	dec,_ := base64.StdEncoding.DecodeString(userName)
	if userName != ""{
		this.Data["userName"] = string(dec)
		this.Data["checked"] = "checked"
	}else{
		this.Data["userName"] = ""
		this.Data["checked"] = ""
	}

	this.TplName = "login.html"
}

//处理登录业务
func(this*UserController)HandleLogin(){
	//获取数据
	userName := this.GetString("userName")
	pwd := this.GetString("password")
	//校验数据
	if userName == "" || pwd == "" {
		beego.Error("传输数据不完整")
		this.TplName = "login.html"
		return
	}
	//处理数据
	o := orm.NewOrm()
	var user models.User
	user.Name = userName
	err := o.Read(&user,"Name")
	if err != nil {
		beego.Error("用户名不存在")
		this.TplName = "login.html"
		return
	}

	if user.Pwd != pwd{
		beego.Error("密码错误")
		this.TplName = "login.html"
		return
	}

	//实现记住用户名功能  上一次登陆成功以后，点击了记住用户名，下一次登陆的时候默认显示用户名
	remember:= this.GetString("remember")
	//给userName加密
	enc := base64.StdEncoding.EncodeToString([]byte(userName))
	if remember == "on"{
		this.Ctx.SetCookie("userName",enc,60)
	}else {
		this.Ctx.SetCookie("userName",userName,-1)
	}

	//session存储
	this.SetSession("userName",userName)
	//返回数据
	this.Redirect("/article/index",302)
}

//退出登录
func(this*UserController)Logout(){
	//删除session然后跳转到登录页面
	this.DelSession("userName")

	this.Redirect("/login",302)
}