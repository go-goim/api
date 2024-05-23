package v1

import (
	"github.com/go-goim/api/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func (x *StorageMessage) TableName() string {
	return "message_history" // TODO: use config
}

type StorageMessageDBContent struct {
	Content     string             `json:"content"`
	ContentType MessageContentType `json:"content_type"`
}

type StorageMessageDB struct {
	MsgID       types.ID                 `gorm:"column:id;primaryKey"`
	SessionID   string                   `gorm:"column:session_id"`
	SessionType SessionType              `gorm:"column:session_type,type:tinyint"`
	FromUID     types.ID                 `gorm:"column:from_uid,type:bigint"`
	ToUID       types.ID                 `gorm:"column:to_uid,type:bigint"`
	Content     *StorageMessageDBContent `gorm:"column:content, type:bytes, serializer:json"`
	CreateTime  int64                    `gorm:"column:create_time,type:int"`
}

func (s *StorageMessageDB) ToProto() *StorageMessage {
	return &StorageMessage{
		MsgID:       s.MsgID.Int64(),
		SessionID:   s.SessionID,
		SessionType: s.SessionType,
		FromID:      s.FromUID.Int64(),
		ToID:        s.ToUID.Int64(),
		Content:     &StorageMessage_Content{Content: s.Content.Content, ContentType: s.Content.ContentType},
		CreateTime:  timestamppb.New(time.UnixMilli(s.CreateTime)),
	}
}

func (s *StorageMessageDB) FromProto(m *StorageMessage) *StorageMessageDB {
	s.MsgID = types.ID(m.MsgID)
	s.SessionID = m.SessionID
	s.SessionType = m.SessionType
	s.FromUID = types.ID(m.FromID)
	s.ToUID = types.ID(m.ToID)
	s.Content = &StorageMessageDBContent{Content: m.Content.Content, ContentType: m.Content.ContentType}
	s.CreateTime = m.CreateTime.AsTime().UnixMilli()
	return s
}
