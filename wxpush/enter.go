package wxpush

type Markdown struct {
	userData
	DoTask []string
	sign
	DoGift []string
	Time   string
}

type sign struct {
	Sign  string
	SignB bool
	Bad   string
	Good  string
	Lunar string
}

type userData struct {
	UserId   string
	UserName string
	RoleName string
	RoleJob  string
}
