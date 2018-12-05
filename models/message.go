package models

type Message struct {
	Model
	Key string `gorm:"unique_index;not null;" json:"key"`
	Content string 	`json:"content"`
	UserId int `json:"user_id"`
	User User `json:"user"`
	NoteKey string `json:"note_key"`
	Praise int `gorm:"default:0" json:"praise"`
}

func SaveMessage(message *Message) error {
	return db.Save(message).Error
}

func QueryMessageByNoteKey(noteKey string) (ms []*Message,err error)  {
	return ms,db.Where("note_key = ?",noteKey).Order("updated_at desc").Find(&ms).Error
}


