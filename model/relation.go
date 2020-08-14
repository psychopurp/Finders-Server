package model

import (
	"errors"
	"finders-server/global"
	"finders-server/pkg/e"
	"time"

	uuid "github.com/satori/go.uuid"
)

/*
DB Table Details
用户关系表

CREATE TABLE `relations` (
  `relation_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户关系ID',
  `relation_type` int(11) NOT NULL COMMENT '关系类型',
  `relation_group` varchar(20) NOT NULL COMMENT '关系组名',
  `from_uid` varchar(50) NOT NULL COMMENT '用户ID',
  `to_uid` varchar(50) NOT NULL COMMENT '被关注用户ID',
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
	RelationGroup string    `gorm:"column:relation_group;type:varchar(20);" json:"relation_group"`             //[ 2] relation_group                                 VARCHAR[20]          null: false  primary: false  auto: false
	FromUID       uuid.UUID `gorm:"column:from_uid;type:varchar(50);" json:"from_uid"`                         //[ 3] from_uid                                       VARCHAR[30]          null: false  primary: false  auto: false
	ToUID         uuid.UUID `gorm:"column:to_uid;type:varchar(50);" json:"to_uid"`                             //[ 4] to_uid                                         VARCHAR[30]          null: false  primary: false  auto: false
	CreatedAt     time.Time `gorm:"column:created_at;type:DATETIME;" json:"created_at"`                        //[ 5] created_at                                     DATETIME             null: false  primary: false  auto: false
	UpdatedAt     time.Time `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`                        //[ 6] updated_at                                     DATETIME             null: true   primary: false  auto: false
	//DeletedAt     *time.Time `gorm:"column:deleted_at;type:DATETIME;" json:"deleted_at"`                        //[ 6] deleted_at                                     DATETIME             null: true   primary: false  auto: false
}

const baseIndex = 1

// RelationType
const (
	FOLLOW = baseIndex + iota
	DENY
)

// RelationGroup
const (
	FOLLOW_RELATION_GROUP = "followRelation"
	DENY_RELATION_GROUP   = "denyRelation"
)

var relationGroupByIndex = map[int]string{
	FOLLOW: FOLLOW_RELATION_GROUP,
	DENY:   DENY_RELATION_GROUP,
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

func ExistRelation(fromUID, toUID string, relationType int) (isExist bool, err error) {
	db := global.DB
	if relationGroupByIndex[relationType] == "" {
		return false, errors.New(e.TYPE_ERROR)
	}
	data := make(map[string]interface{})
	data["from_uid"] = fromUID
	if toUID != "" {
		data["to_uid"] = toUID
	}

	data["relation_type"] = relationType
	var relation Relation
	isExist = !db.Where(data).First(&relation).RecordNotFound()
	return
}

func ExistRelationByData(data map[string]interface{}) (isExist bool, err error) {
	db := global.DB
	if data["relation_type"] != "" && relationGroupByIndex[data["relation_type"].(int)] == "" {
		return false, errors.New(e.TYPE_ERROR)
	}
	if data["relation_group"] != "" {
		delete(data, "relation_group")
	}
	var relation Relation
	isExist = !db.Where(data).First(&relation).RecordNotFound()
	return
}

func AddRelation(fromUID, toUID string, relationType int) (relation Relation, err error) {
	db := global.DB
	if relationGroupByIndex[relationType] == "" {
		return relation, errors.New(e.TYPE_ERROR)
	}
	relation.FromUID = uuid.FromStringOrNil(fromUID)
	relation.ToUID = uuid.FromStringOrNil(toUID)
	relation.RelationType = relationType
	relation.RelationGroup = relationGroupByIndex[relationType]
	err = db.Create(&relation).Error
	return
}

func DeleteRelation(fromUID, toUID string, relationType int) (relation Relation, err error) {
	db := global.DB
	if relationGroupByIndex[relationType] == "" {
		return relation, errors.New(e.TYPE_ERROR)
	}
	data := make(map[string]interface{})
	data["from_uid"] = fromUID
	data["to_uid"] = toUID
	data["relation_type"] = relationType
	err = db.Unscoped().Where(data).Delete(&relation).Error
	return
}

func GetRelations(fromUID uuid.UUID, relationType int) (relations []Relation, err error) {
	db := global.DB
	if relationGroupByIndex[relationType] == "" {
		return relations, errors.New(e.TYPE_ERROR)
	}
	data := make(map[string]interface{})
	data["from_uid"] = fromUID.String()
	data["relation_type"] = relationType
	err = db.Where(data).Find(&relations).Error
	return
}

func GetRelationsByData(data map[string]interface{}) (relations []Relation, err error) {
	db := global.DB
	if data["relation_type"] != "" && relationGroupByIndex[data["relation_type"].(int)] == "" {
		return relations, errors.New(e.TYPE_ERROR)
	}
	if data["relation_group"] != "" {
		delete(data, "relation_group")
	}
	err = db.Model(&Relation{}).Where(data).Find(&relations).Error
	return
}

func UpdateRelationType(data map[string]interface{}, relationType int) (relation Relation, err error) {
	db := global.DB
	if relationGroupByIndex[relationType] == "" {
		return relation, errors.New(e.TYPE_ERROR)
	}
	relation.RelationType = relationType
	relation.RelationGroup = relationGroupByIndex[relationType]
	err = db.Model(&Relation{}).Where(data).
		Updates(&relation).
		Error
	return
}
