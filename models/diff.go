package models

type Diff struct {
	Id             int    `gorm:"primaryKey;autoIncrement:true"`
	Content        string `gorm:"unique;index"`
	CommitMessages []CommitMessage
}
