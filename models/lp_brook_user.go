package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type LpBrookUser struct {
	Id          int       `orm:"column(u_id);auto"`
	Email       string    `orm:"column(u_email);size(255)" description:"邮箱" valid:"Email; MaxSize(50)"`
	Name        string    `orm:"column(u_name);size(255)" description:"用户名" valid:"Range(2, 20)"`
	Passwd      string    `orm:"column(u_passwd);size(255)" description:"密码" valid:"Length(32)"`
	Port        int       `orm:"column(u_port);size(255)" description:"端口"`
	Flow        float64   `orm:"column(u_flow);digits(40);decimals(5)" description:"剩余流量"`
	IsAdmin     int       `orm:"column(u_is_admin)" description:"是否是管理员 0普通用户/1管理员/-1停用"`
	ExpireTime  time.Time `orm:"column(expire_time);type(timestamp);" description:"vip到期时间"`
	FlowTotal   float64   `orm:"column(u_flow_total);digits(40);decimals(5)" description:"总使用流量"`
	Money       int       `orm:"column(u_money)" description:"金币"`
	ProxyPasswd string    `orm:"column(u_proxy_passwd)" description:"代理连接密码"`
	ReginIp     string    `orm:"column(r_ip)" description:"注册ip"`
	// TableTime  time.Time `orm:"column(table_time);type(datetime);auto_now" description:"直接修改表的日期"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);auto_now_add" description:"创建日期"`
	UpdateTime time.Time `orm:"column(update_time);type(timestamp);auto_now" description:"更新日期"`
}

// AddLpBrookUser insert a new LpBrookUser into database and returns
// last inserted Id on success.
func AddLpBrookUser(m *LpBrookUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLpBrookUserById retrieves LpBrookUser by Id. Returns error if
// Id doesn't exist
func GetLpBrookUserById(id int) (v *LpBrookUser, err error) {
	o := orm.NewOrm()
	v = &LpBrookUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateLpBrookUser updates LpBrookUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateLpBrookUserById(m *LpBrookUser) (err error) {
	o := orm.NewOrm()
	v := LpBrookUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLpBrookUser deletes LpBrookUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLpBrookUser(id int) (err error) {
	o := orm.NewOrm()
	v := LpBrookUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LpBrookUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//Get All
func GetLpBrookUserAll() (v []*LpBrookUser, err error) {
	o := orm.NewOrm()
	// 获取 QuerySeter 对象
	lpBrookUserusers := make([]*LpBrookUser, 0)
	qs := o.QueryTable(LpBrookUserTBName())
	_, err = qs.All(&lpBrookUserusers)
	return lpBrookUserusers, err
}

func UpdateUserFlowById(Id int, flow float64) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE lp_brook_user SET u_flow = u_flow - ?  WHERE u_id = ?", flow, Id).Exec()
	return err
}
