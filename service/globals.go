package service

import (
	"TGU-MAP/models"
	"TGU-MAP/service/aliasItem"
	"TGU-MAP/service/feedback"
	"TGU-MAP/service/listItem"
	"TGU-MAP/service/noticeItem"
	"TGU-MAP/service/user"
	"github.com/go-redis/redis/v8"
)

var (
	ListItemClient   *listItem.ListItemStub
	UserClient       *user.UserStub
	AliasItemClient  *aliasItem.AliasItemStub
	NoticeItemClient *noticeItem.NoticeItemStub
	FeedbackClient   *feedback.FeedbackStub
	RDB              *redis.Client
	GlobalConfig     models.Config
)
