package admin

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/nfnt/resize"
	"golang.org/x/net/context"
	"image"
	"image/jpeg"
	"io/ioutil"
	"os"
	"qiniupkg.com/api.v7/kodo"
	"strings"
)

var (
	imgdir = "./upload/"
)

type UploadHandle struct {
	baseController
}

//上传页面
func (this *UploadHandle) UpLoadPage() {
	obj := this.GetString("obj")
	this.Data["obj"] = obj
	this.TplName = "admin/upload.html"

}

//七牛上传handel
func (this *UploadHandle) QiniuUpLoadFile() {
	if this.IsPost() {
		//远程服务器地址
		imgserver := beego.AppConfig.String("imgserver")
		var err error
		file, fileHead, err := this.GetFile("uploadfile")
		if err != nil {
			this.Ctx.WriteString("post err:" + err.Error())
			return
		}
		defer file.Close()
		obj := this.GetString("obj")
		filename := fileHead.Filename
		split_part := strings.Split(filename, ".")
		ext := "." + strings.ToLower(split_part[1])

		//save location file
		filename = this.GetGUID() + ext
		tempfilename := "temp_+filename"
		data, err := ioutil.ReadAll(file)
		err = ioutil.WriteFile(imgdir+tempfilename, data, 0777)

		if err != nil {
			this.Ctx.WriteString("save to file err:" + err.Error())
			return
		} else {
			//图片大小控制
			WaterMark(imgdir+tempfilename, imgdir+filename, ext)
			//开始上传到qiniu服务器
			kodo.SetMac("jgzDpnMK3mprXixO9zX3fJlmp2hp0EtJ07kU_NR9", "hVaszp4-IequAWsKh6pBVkmarWvs8RCnqZZzbNzj")
			zone := 0
			c := kodo.New(zone, nil)
			bucket := c.Bucket("aimeiju")
			ctx := context.Background()
			localFile := imgdir + filename
			err := bucket.PutFile(ctx, nil, filename, localFile, nil)
			if err != nil {
				this.Ctx.WriteString("upload imgserver err:" + err.Error())
				return
			} else {
				if len(obj) == 0 {
					obj = "photo"
				}
				this.Ctx.WriteString("<a href='../upload/add'>重新上传..</a><script>window.parent.setphoto('" + obj + "','" + imgserver + filename + "');</script>") //输出远程服务器地址
			}
		}

	} else {
		this.Ctx.WriteString("请求错误...")
		return
	}
}
func WaterMark(tempfilepath string, newfilepath string, ext string) bool {
	imgb, _ := os.Open(tempfilepath)
	img, _, err := image.Decode(imgb)
	img = Resize(img)
	defer imgb.Close()
	out, err := os.Create(newfilepath)
	defer out.Close()
	jpeg.Encode(out, img, &jpeg.Options{100})
	os.Remove(tempfilepath)
	if err != nil {
		return false
	} else {

		return true
	}
}

func Resize(img image.Image) (m image.Image) {
	var maximage uint
	maximage = 600
	if img.Bounds().Dx() > 600 {
		return resize.Resize(maximage, 0, img, resize.MitchellNetravali)
	} else {
		return img
	}
}
