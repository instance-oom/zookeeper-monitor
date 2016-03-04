package models

import "github.com/astaxie/beego/orm"

//Mail is mail information
type Mail struct {
	ID      int    `orm:"column(ID)"`
	Address string `orm:"column(Address)"`
}

//AddMail : Add alert mail info
func AddMail(mail *Mail) (int64, error) {
	return orm.NewOrm().Insert(mail)
}

//Update : Update alert mail info
func (mail *Mail) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(mail, "Address"); err != nil {
		return err
	}
	return nil
}

//GetMail : Get alert mail info
func GetMail() (*Mail, error) {
	var mails []*Mail
	_, err := orm.NewOrm().QueryTable("mail").Limit(10).All(&mails)
	if len(mails) > 0 {
		return mails[0], err
	}
	return new(Mail), err
}
