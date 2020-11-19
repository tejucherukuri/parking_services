package models

import (
	"reflect"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ParkslotOrm struct {
	orm.Ormer
}

func NewOrm() *ParkslotOrm {
	db := orm.NewOrm()
	return &ParkslotOrm{db}
}

func (this *ParkslotOrm) InsertB(md interface{}) (i int64, e error) {
	curTime := time.Now()

	reflect.ValueOf(md).Elem().FieldByName("CreatedDate").Set(reflect.ValueOf(curTime))
	reflect.ValueOf(md).Elem().FieldByName("LastModified").Set(reflect.ValueOf(curTime))
	i, e = this.Insert(md)
	beego.Debug(e)
	return
}

