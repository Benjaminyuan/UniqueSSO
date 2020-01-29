package common

import (
	"unique/jedi/entity"
	"unique/jedi/session"
)

func InjectUserInfoToSession(userSession session.Session,user *entity.User){
	userSession.Set(SessionUserNameKey,user.Name)
	userSession.Set(SessionUserUIDKey,user.UID)
	userSession.Set(SessionUserCollegeKey,user.College)
	userSession.Set(SessionUserEmailKey,user.EMail)
}