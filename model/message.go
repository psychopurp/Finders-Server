package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details


CREATE TABLE `messages` (
  `message_id` int(11) NOT NULL COMMENT '消息ID',
  `message_info` text COMMENT '消息内容',
  `message_status` int(11) NOT NULL COMMENT '消息状态状态',
  `send_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '发送时间',
  `from_uid` varchar(30) NOT NULL COMMENT '发送用户ID',
  `to_uid` varchar(30) NOT NULL COMMENT '接受用户ID',
  `message_type` int(11) DEFAULT NULL COMMENT '消息类型',
  `link` varchar(100) DEFAULT NULL COMMENT '消息链接',
  PRIMARY KEY (`message_id`),
  KEY `fk_letter_user_id1` (`from_uid`),
  KEY `fk_letter_user_id2` (`to_uid`),
  CONSTRAINT `fk_letter_user_id1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_letter_user_id2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Message struct is a row record of the messages table in the employees database
type Message struct {
	MessageID     int         `gorm:"column:message_id;type:INT;primary_key" json:"message_id"`      //[ 0] message_id                                     INT                  null: false  primary: true   auto: false
	MessageInfo   null.String `gorm:"column:message_info;type:TEXT;size:65535;" json:"message_info"` //[ 1] message_info                                   TEXT[65535]          null: true   primary: false  auto: false
	MessageStatus int         `gorm:"column:message_status;type:INT;" json:"message_status"`         //[ 2] message_status                                 INT                  null: false  primary: false  auto: false
	SendTime      time.Time   `gorm:"column:send_time;type:DATETIME;" json:"send_time"`              //[ 3] send_time                                      DATETIME             null: false  primary: false  auto: false
	FromUID       string      `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"`         //[ 4] from_uid                                       VARCHAR[30]          null: false  primary: false  auto: false
	ToUID         string      `gorm:"column:to_uid;type:VARCHAR;size:30;" json:"to_uid"`             //[ 5] to_uid                                         VARCHAR[30]          null: false  primary: false  auto: false
	MessageType   null.Int    `gorm:"column:message_type;type:INT;" json:"message_type"`             //[ 6] message_type                                   INT                  null: true   primary: false  auto: false
	Link          null.String `gorm:"column:link;type:VARCHAR;size:100;" json:"link"`                //[ 7] link                                           VARCHAR[100]         null: true   primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) BeforeSave() error {
	return nil
}

func (m *Message) Prepare() {
}

func (m *Message) Validate(action Action) error {

	return nil
}
