package main

import (
    "flag"
    "net/http"

    "code.google.com/p/uedatakuya-goweb/goweb"
    "labix.org/v2/mgo"
    "./diary"
    "log"
)

var (
    addr string
    dbhost string
    dbname string
)

func init() {
    // Webサーバのアドレスをコマンドライン引数から取得する
    flag.StringVar(&addr, "http", ":8080", "Addres of web server")
    // データベースのホスト名をコマンドライン引数から取得する
    flag.StringVar(&dbhost, "dbhost", "localhost", "DB hosts")
    // データベース名をコマンドライン引数から取得する
    flag.StringVar(&dbname, "dbname", "diary", "DB name")
}

func main() {

    flag.Parse()

    // DBをオープンする
    log.Println("Accessing DB...")
    session, err := mgo.Dial(dbhost)
    if err != nil {
        log.Fatal(err.Error())
    }
    defer session.Close()
    session.SetMode(mgo.Monotonic, true)
    db := session.DB(dbname)
    log.Printf("DB (%s %s) opend.", dbhost, dbname)

    // ハンドラの初期化
    log.Println("Initializing handlers...")
    diary.Init(db)
    goweb.ConfigureDefaultFormatters()
    http.Handle("/", goweb.DefaultHttpHandler)

    // Webサーバの起動
    log.Printf("Starting Diary Server at %s", addr)
    http.ListenAndServe(addr, nil)
}


