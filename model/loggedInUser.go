package model

type LoggedInUser struct {
	userId int64
	email  string
}

func (l *LoggedInUser) SaveUser(userId int64, email string) {
	l.userId = userId
	l.email = email

}
func (l *LoggedInUser) GetLoggedInUserId() int64 {
	return l.userId

}
