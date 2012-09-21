package diary

import (
    "code.google.com/p/uedatakuya-goweb/goweb"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
    "net/http"
    "log"
)

type Controller struct {
    db *mgo.Database
}
func NewController(db *mgo.Database) *Controller {
    if db == nil {
        panic("database cannot be nil")
    }

    return &Controller{db}
}

func (cr *Controller) Create(cx *goweb.Context) {
    log.Println("Creating a diary...")
    c := cr.db.C(COLLECTION)

    var diary Diary
    decoder := new(goweb.JsonRequestDecoder)
    decoder.Unmarshal(cx, &diary)

    diary.Id = bson.NewObjectId().Hex()

    if err := c.Insert(&diary); err != nil {
       log.Println("Error: %s", err.Error())
       cx.RespondWithError(http.StatusForbidden)
       return
    }

    log.Printf("Created diary id=%s", diary.Id)
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.RespondWithData(diary.Id)
}

func (cr *Controller) Delete(id string, cx *goweb.Context) {
    log.Printf("Deleting a diary id=%s...", id)
    c := cr.db.C(COLLECTION)
    if err := c.RemoveId(id); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Printf("Deleted diary id=%s", id)
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "DELETE")
    cx.RespondWithOK()
}

func (cr *Controller) DeleteMany(cx *goweb.Context) {
    log.Println("Deleting all diaries...")
    c := cr.db.C(COLLECTION)
    if _, err := c.RemoveAll(nil); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Println("Deleted all diaries")
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.RespondWithOK()
}

func (cr *Controller) Read(id string, cx *goweb.Context) {
    log.Printf("Read a diary id=%s", id)
    c := cr.db.C(COLLECTION)
    var diary Diary
    if err := c.FindId(id).One(&diary); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Printf("Read diary id=%s", id)
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.RespondWithData(diary)
}

func (cr *Controller) ReadMany(cx *goweb.Context) {
    log.Println("Read all diaries...")
    c := cr.db.C(COLLECTION)
    count, err := c.Count()
    if err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    diaries := make([]*Diary, count)
    if err := c.Find(nil).All(&diaries); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Printf("Read all %d diaries", count)
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.RespondWithData(diaries)
}

func (cr *Controller) Update(id string, cx *goweb.Context) {
    log.Printf("Update a diary id=%s...", id)
    c := cr.db.C(COLLECTION)

    var diary *Diary
    decoder := new(goweb.JsonRequestDecoder)
    decoder.Unmarshal(cx, &diary)

    if err := c.UpdateId(id, diary); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Printf("Updated a diary id=%s", id)
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Headers","*");
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Methods","PUT,DELETE,POST,GET,OPTIONS");
    cx.RespondWithOK()
}

func (cr *Controller) UpdateMany(cx *goweb.Context) {
    log.Println("Update all diaries...")
    c := cr.db.C(COLLECTION)

    var diaries []*Diary
    decoder := new(goweb.JsonRequestDecoder)
    decoder.Unmarshal(cx, &diaries)

    if _, err := c.UpdateAll(nil, diaries); err != nil {
        log.Println("Error: %s", err.Error())
        cx.RespondWithError(http.StatusForbidden)
        return
    }

    log.Println("Updated all diaries")
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.RespondWithOK()
}

func (cr *Controller) Options(cx *goweb.Context) {
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Headers","*");
    cx.ResponseWriter.Header().Set("Access-Control-Allow-Methods","PUT,DELETE,POST,GET,OPTIONS");

    cx.RespondWithOK()
}
