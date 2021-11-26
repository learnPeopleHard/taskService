package userinfo

type UserInfoPO struct {
	Id int64 `db:"id"`
	UserId int64 `db:"user_id"`
	UserName string `db:"username"`
	Nickname string  `db:"nickname"`
	ProfilePicture string `db:"profile_picture"`
}
