package admin

import (
	"github.com/zituocn/VMovie/models"
	"strings"
)

type UserHandel struct {
	baseController
}

//更新密码
func (this *UserHandel) ChangePass() {

	this.Data["AdminDir"] = this.admindir
	this.TplName = "admin/changepassword.html"

}

func (this *UserHandel) SavePass() {
	var (
		ypass   string
		newpass string
		epass   string
		info    models.UserInfo
	)

	ypass = strings.TrimSpace(this.GetString("ypass"))
	newpass = strings.TrimSpace(this.GetString("newpass"))
	epass = strings.TrimSpace(this.GetString("epass"))

	if len(ypass) == 0 || len(newpass) == 0 || len(epass) == 0 {
		this.showmsg("请填写完整密码信息...")
		return
	}
	if newpass != epass {
		this.showmsg("新密码两次输入不一样，请重试...")
		return
	}
	info.Id = this.userid
	info.Read()
	if info.Password != models.Md5(ypass) {
		this.showmsg("原密码不正确，请返回重新输入...")
		return
	}
	info.Password = models.Md5(newpass)
	err := info.Update("password")
	if err == nil {
		this.showmsg("密码更新成功...", this.admindir+"user/changepassword/")
	} else {
		this.showmsg("密码更新失败，请重试...")
	}
}
