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


CREATE TABLE `relations` (
  `relation_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户关系ID',
  `relation_type` int(11) NOT NULL COMMENT '关系类型',
  `relation_group` varchar(20) NOT NULL COMMENT '关系组名',
  `from_uid` varchar(30) NOT NULL COMMENT '用户ID',
  `to_uid` varchar(30) NOT NULL COMMENT '被关注用户ID',
  `created_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系建立时间',
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '关系更新时间',
  PRIMARY KEY (`relation_id`),
  UNIQUE KEY `unique_user_user` (`from_uid`,`to_uid`) USING BTREE,
  KEY `fk_user_id2` (`to_uid`),
  CONSTRAINT `fk_user_id1` FOREIGN KEY (`from_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_user_id2` FOREIGN KEY (`to_uid`) REFERENCES `users` (`user_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8

*/

// Relation struct is a row record of the relations table in the employees database
type Relation struct {
	RelationID    int       `gorm:"AUTO_INCREMENT;column:relation_id;type:INT;primary_key" json:"relation_id"` //[ 0] relation_id                                    INT                  null: false  primary: true   auto: true
	RelationType  int       `gorm:"column:relation_type;type:INT;" json:"relation_type"`                       //[ 1] relation_type                                  INT                  null: false  primary: false  auto: false
	RelationGroup string    `gorm:"column:relation_group;type:VARCHAR;size:20;" json:"relation_group"`         //[ 2] relation_group                                 VARCHAR[20]          null: false  primary: false  auto: false
	FromUID       string    `gorm:"column:from_uid;type:VARCHAR;size:30;" json:"from_uid"`                     //[ 3] from_uid                                       VARCHAR[30]          null: false  primary: false  auto: false
	ToUID         string    `gorm:"column:to_uid;type:VARCHAR;size:30;" json:"to_uid"`                         //[ 4] to_uid                                         VARCHAR[30]          null: false  primary: false  auto: false
	CreatedAt     time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                        //[ 5] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt     null.Time `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                        //[ 6] updated_at                                     DATETIME             null: true   primary: false  auto: false

}

// TableName sets the insert table name for this struct type
func (r *Relation) TableName() string {
	return "relations"
}

func (r *Relation) BeforeSave() error {
	return nil
}

func (r *Relation) Prepare() {
}

func (r *Relation) Validate(action Action) error {

	return nil
}
