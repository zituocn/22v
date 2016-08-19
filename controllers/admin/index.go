package admin

import (
	"github.com/astaxie/beego"
	"github.com/zituocn/VMovie/models"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type IndexHandel struct {
	baseController
}

type LoginHandel struct {
	baseController
}

func (this *IndexHandel) Main() {
	this.TplName = "admin/main.html"
}

func (this *IndexHandel) Left() {
	this.Data["username"] = this.username
	this.TplName = "admin/left.html"
}

func (this *IndexHandel) Right() {
	this.Data["hostname"], _ = os.Hostname()
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH
	this.Data["beegover"] = beego.VERSION
	this.Data["clientip"] = this.getClientIp()

	this.Data["systemver"] = beego.AppConfig.String("systemver")
	this.Data["developer"] = beego.AppConfig.String("developer")
	this.Data["servertime"] = this.FormatTime(time.Now(), "YYYY年MM月DD日 HH:mm:ss")

	//影片数量
	var movie models.MovieInfo
	count, _ := movie.Query().Count()
	this.Data["moviecount"] = count
	this.TplName = "admin/right.html"
}

//login post
func (this *LoginHandel) Login() {

	if this.IsPost() {
		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		var info string
		if len(username) == 0 || len(password) == 0 {
			info = "请填写登录帐号和密码..."
		} else {
			var user models.UserInfo
			user.Username = username
			if user.Read("username") != nil || user.Password != models.Md5(password) {
				info = "帐号或密码错误..."
			} else {
				user.Logintimes += 1
				user.Lastloginip = this.getClientIp()
				user.Lastlogintime = this.getTime()
				user.Update()
				authKey := models.Md5("samsong|" + user.Password)
				this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authKey)
				this.Redirect(this.admindir+"main", 302)
			}
		}
		this.Data["username"] = username
		this.Data["info"] = info
	}
	this.TplName = "admin/login.html"
}

//logout
func (this *LoginHandel) Logout() {
	this.Ctx.SetCookie("auth", "")
	this.Ctx.WriteString("<script>top.location.href='" + this.admindir + "'</script>")
}
