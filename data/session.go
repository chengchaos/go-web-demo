package data

type Session struct {
	Id     int
	Uuid   string
	Email  string
	UserId int
	//CreateAt time.Time
}

func (sess *Session) Check() (ok bool, err error) {
	return true, err

}
