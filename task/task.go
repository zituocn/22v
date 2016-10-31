package task

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	"github.com/zituocn/VMovie/models"
	"strings"
)

func init() {

	///添加向baidu站长平台push数据的task
	///默认为10分钟一次...
	time_con := beego.AppConfig.String("baidupush")
	if len(time_con) == 0 {
		time_con = "0 */10 * * * *"
	}
	baidu := toolbox.NewTask("baidu_push", time_con, PushBaidu)
	err := baidu.Run()
	if err != nil {
		beego.Error(err)
	}
	toolbox.AddTask("baidu_push", baidu)
	toolbox.StartTask()
	defer toolbox.StopTask()
}

///2016.10.26 sam增加
///添加向baidu站长平台push地址的方法
func PushBaidu() error {
	var (
		info     models.MovieInfo
		list     []*models.MovieInfo
		item     string
		push_url = "http://data.zz.baidu.com/urls?site=22v.net&token=jTGxCEp41rmXCVFT"
	)

	urls := []string{}
	ids := []int64{}

	//查询最近未push的数据
	query := info.Query()
	query = query.Filter("status__gte", 0)
	query = query.Filter("ispush", 0).OrderBy("id")
	query.Limit(20, 0).All(&list, "id") //一次20条数据...

	//构造url和ids
	if len(list) > 0 {
		for _, i := range list {
			ids = append(ids, i.Id)
			item = fmt.Sprintf("https://22v.net/v/%d/", i.Id)
			urls = append(urls, item)
		}
	}

	if len(ids) > 0 {
		o := orm.NewOrm()
		o.QueryTable("movie_info").Filter("id__in", ids).Update(orm.Params{
			"Ispush": 1,
		})
	}

	if len(urls) > 0 {
		//向baidu接口push信息
		req := httplib.Post(push_url)
		req.Body(strings.Join(urls, "\n"))
		result, _ := req.String()
		beego.Info("baidu api返回结果：" + result)
	}
	beego.Info("baidu push end")
	return nil
}
