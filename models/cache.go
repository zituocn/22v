package models

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"time"
)

var cc cache.Cache

func init() {
	var err error
	cc, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		beego.Info(err)
	}
}

func SetCache(key string, value interface{}) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cache:cc is nil")
	}
	err = cc.Put(key, data, 3600*time.Second)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cache:cc is nil")
	}
	data := cc.Get(key)
	if data == nil {
		return errors.New("cache:缓存不存在" + key)
	}
	err := Decode(data.([]byte), to)
	return err
}

func RemoveCache(key string) error {
	if cc == nil {
		return errors.New("cache:cc is nil")
	}

	err := cc.Delete(key)
	if err != nil {
		return errors.New("cache:cahce 删除失败")
	} else {
		return nil
	}

}

// --------------------
// Encode
// 用gob进行数据编码
//
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// -------------------
// Decode
// 用gob进行数据解码
//
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
