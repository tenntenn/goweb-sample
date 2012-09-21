package diary

import (
    "code.google.com/p/uedatakuya-goweb/goweb"
    "labix.org/v2/mgo"
    "log"
)

const COLLECTION = "diary"

func Init(db *mgo.Database) {
    log.Println("Initializing diary package")
    controller := NewController(db)
    goweb.MapRest("/diary", controller)
}
