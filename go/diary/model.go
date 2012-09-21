package diary

import (
    "time"
)

type Diary struct {
    Id string `json:"_id" bson:"_id"`
    Date time.Time `json:"-" bson:"-"`
    Title string `json:"title" bson:"title"`
    Content string `json:"content" bson:"content"`
}

func NewModel(title string, selected bool) *Diary {
    return &Diary{"", time.Now(), title, ""}
}
