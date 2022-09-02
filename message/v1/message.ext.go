package v1

import (
	"strconv"
)

// This file is for extend proto generated structs methods.

func (x *StorageMessage) HbaseRowKey() string {
	return x.RowKey
}

func (x *StorageMessage) HbaseValues() map[string]map[string][]byte {
	return map[string]map[string][]byte{
		"users": {
			"from":      []byte(strconv.FormatInt(x.Users.FromID, 10)),
			"to":        []byte(strconv.FormatInt(x.Users.ToID, 10)),
			"sessionId": []byte(x.Users.SessionID),
		},
		"content": {
			"type": []byte(strconv.Itoa(int(x.Content.ContentType))),
			"text": []byte(x.Content.Content), // TODO: use content instead of text
		},
		"extra": {
			"timestamp": []byte(strconv.FormatInt(x.Extra.CreateTime.Seconds, 10)),
		},
	}
}

func (x *StorageMessage) TableName() string {
	return "message_history"
}
