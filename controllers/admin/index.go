package admin

import (
	"github.com/astaxie/beego"
	"os"
	"runtime"
)

type IndexController struct {
	baseController
}

func (this *IndexController) Index() {

	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username

	this.TplNames = this.moduleName + "/index/index.html"
}

func (this *IndexController) Main() {

	this.Data["hostname"], _ = os.Hostname()
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	

	this.display()
}
