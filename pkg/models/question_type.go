package models

type QuestionType struct {
	ID   uint   `json:"id" bson:"_id" gorm:"column:id"`
	Name string `json:"name" bson:"name" gorm:"column:name"`
}

func (QuestionType) TableName() string {
	return "question_types"
}

func (QuestionType) CollectionName() string {
	return "question_types"
}
