package userinfo

type UserPO struct {
	Id int64 `db:"id"`
	Name string  `db:"name"`
	Pwd string `db:"pwd"`
}
