// Code Written Manually

package v1

func (x *UserInternal) ToUser() *User {
	return &User{
		Uid:         x.GetUid(),
		Name:        x.GetName(),
		Email:       x.Email,
		Phone:       x.Phone,
		Avatar:      x.GetAvatar(),
		ConnectUrl:  x.PushServerIp,
		LoginStatus: x.GetLoginStatus(),
	}
}
