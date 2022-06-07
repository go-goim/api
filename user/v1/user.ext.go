// Code Written Manually

package v1

func (x *UserInternal) ToUser() *User {
	return &User{
		Uid:         x.GetUid(),
		Name:        x.GetName(),
		Email:       x.GetEmail(),
		Phone:       x.GetPhone(),
		Avatar:      x.GetAvatar(),
		ConnectUrl:  x.PushServerIp,
		LoginStatus: x.GetLoginStatus(),
	}
}
