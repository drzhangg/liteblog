package models

type Note struct {
	Model
	Key string `gorm:"unique;not null"`
	UserID int
	User User
	Title string	`gorm:"type:text"`
	Summary string	`gorm:"type:text"`
	Content string	`gorm:"type:text"`
	Visit int		`gorm:"default:0"`
	Praise int		`gorm:"default:0"`
}

func QueryNoteByKey(key string) (note Note, err error) {
	return note,db.Where("key = ?",key).Take(&note).Error
}

func QueryNoteByKeyAndUserId(key string, userid int) (note Note, err error) {
	return note,db.Where("key = ? and user_id = ?",key,userid).Find(&note).Error
}

func QueryNoteByPage(page, limit int) (note []*Note, err error) {
	return note,db.Offset((page - 1)*limit).Order("updated_at desc").Limit(limit).Find(&note).Error
}

func QueryNotesCount() (count int, err error) {
	return count,db.Model(&Note{}).Count(&count).Error
}

func DeleteUserByKeyAndUserId(key string, userid int) error {
	return db.Delete(&Note{},"key = ? and user_id = ?",key,userid).Error
}

func SaveNote(note *Note) error {
	return db.Save(note).Error
}