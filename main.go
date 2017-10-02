package main

import (
	goz "github.com/TransAssist/goz"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	_ "github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/oauth2"
	"github.com/google/go-github/github"
	"github.com/ktrysmt/go-bitbucket"

)
import (
	"fmt"
	"runtime"
	"os"
	"io/ioutil"
	"log"
	"encoding/json"
	"database/sql"
	"bufio"
	_ "strings"
	_ "time"
	"context"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/oauth2"
)

const windowWidth = 800
const windowHeight = 600

func init() {
	runtime.LockOSThread()
}

func main() {
	fmt.Println("HelloWorld")
	goz.Hello("goz")
	// os.Exit(0)
	fmt.Printf("%s\n", os.Args[0])
	//init glfw
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// make an application window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Hello", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// init gl
	if err := gl.Init(); err != nil {
		panic(err)
	}
	fmt.Println("OpenGL version", gl.GoStr(gl.GetString(gl.VERSION)))

	gl.ClearColor(1.0, 1.0, 1.0, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
func argfunc(){
 if len(os.Args) == 2 {
  fmt.Printf("引数1: %s\n", os.Args[1])
  if os.Args[1]=="ok" {
  }
 }
}

func envfunc(){
 // 既存の環境変数を取得
 fmt.Println("1. $HOME:",os.Getenv("HOME"))
 // プログラム中で設定してやることもできる
 fmt.Println("2. $TEST:",os.Getenv("TEST"))
 os.Setenv("TEST", "abcdedg")
 fmt.Println("3. $TEST:",os.Getenv("TEST"))
 // 環境変数を配列で取得する
 fmt.Println(os.Environ())
}

/** JSONデコード用に構造体定義 */
type Person struct {
 Id       int    `json:"id"`
 Name     string `json:"name"`
 Birthday string `json:"birthday"`
}
func jsonfunc(){
 // JSONファイル読み込み
 bytes, err := ioutil.ReadFile("vro.json")
 if err != nil {
  log.Fatal(err)
 }
 // JSONデコード
 var persons []Person
 if err := json.Unmarshal(bytes, &persons); err != nil {
  log.Fatal(err)
 }
 // デコードしたデータを表示
 for _, p := range persons {
  fmt.Printf("%d : %s\n", p.Id, p.Name)
 }
}
func sqlitefunc(){
 var dbfile string = "./test.db"
 os.Remove( dbfile )
 //	db, err := sql.Open("sqlite3", ":memory:")
 db, err := sql.Open("sqlite3", dbfile)
 if err != nil { panic(err) }
 _, err = db.Exec( `CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))` )
 if err != nil { panic(err) }
 _, err = db.Exec(
  `INSERT INTO "world" ("country", "capital") VALUES (?, ?) `,
  "日本",
  "東京",
 )
 if err != nil { panic(err) }
 stmt, err := db.Prepare( `INSERT INTO "world" ("country", "capital") VALUES (?, ?) ` )
 if err != nil { panic(err) }
 if _, err = stmt.Exec("アメリカ", "ワシントンD.C."); err != nil { panic(err) }
 if _, err = stmt.Exec("ロシア", "モスクワ"); err != nil { panic(err) }
 if _, err = stmt.Exec("イギリス", "ロンドン"); err != nil { panic(err) }
 if _, err = stmt.Exec("オーストラリア", "シドニー"); err != nil { panic(err) }
 stmt.Close()
 db.Close()
 //db, err := sql.Open("sqlite3", "./mydatabase.db")
 //checkErr(err)
 //
 ////データの挿入
 //stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
 //checkErr(err)
 //
 //res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
 //checkErr(err)
 //
 //id, err := res.LastInsertId()
 //checkErr(err)
 //
 //fmt.Println(id)
 ////データの更新
 //stmt, err = db.Prepare("update userinfo set username=? where uid=?")
 //checkErr(err)
 //
 //res, err = stmt.Exec("astaxieupdate", id)
 //checkErr(err)
 //
 //affect, err := res.RowsAffected()
 //checkErr(err)
 //
 //fmt.Println(affect)
 //
 ////データの検索
 //rows, err := db.Query("SELECT * FROM userinfo")
 //checkErr(err)
 //
 //for rows.Next() {
 // var uid int
 // var username string
 // var department string
 // var created time.Time
 // err = rows.Scan(&uid, &username, &department, &created)
 // checkErr(err)
 // fmt.Println(uid)
 // fmt.Println(username)
 // fmt.Println(department)
 // fmt.Println(created)
 //}
 //
 ////データの削除
 //stmt, err = db.Prepare("delete from userinfo where uid=?")
 //checkErr(err)
 //
 //res, err = stmt.Exec(id)
 //checkErr(err)
 //
 //affect, err = res.RowsAffected()
 //checkErr(err)
 //
 //fmt.Println(affect)
 //
 //db.Close()
}
func checkErr(err error) {
 if err != nil {
  panic(err)
 }
}
func interactfunc(){
 if Question("今日は元気ですか？[y/n] ") {
  fmt.Println("そうですか、頑張ってください！")
 } else {
  fmt.Println("元気出してくださいね！")
 }
}
func Question(q string) bool {
 result := true
 fmt.Print(q)

 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {
  i := scanner.Text()

  if i == "Y" || i == "y" {
   break
  } else if i == "N" || i == "n" {
   result = false
   break
  } else {
   fmt.Println("yかnで答えてください。")
   fmt.Print(q)
  }
 }

 if err := scanner.Err(); err != nil {
  panic(err)
 }
 return result
}

func bitbucketaip(name string,pass string){
 //c := bitbucket.NewBasicAuth("","")
 //opt := &bitbucket.PullRequestsOptions{
 // Owner:      "your-team",
 // Repo_slug:  "awesome-project",
 // Source_branch: "develop",
 // Destination_branch: "master",
 // Title: "fix bug. #9999",
 // Close_source_branch: true,
 //}
 //res, err := c.Repositories.PullRequests.Create(opt)
 c := bitbucket.NewBasicAuth(name,pass)
 res, err := c.Repositories.ListPublic()
 if err != nil {
  panic(err)
 }
 spew.Dump(res)
}

func githubapi(token string){
 ctx := context.Background()

 ts := oauth2.StaticTokenSource(
  &oauth2.Token{AccessToken: token},
 )
 tc := oauth2.NewClient(oauth2.NoContext, ts)

 client := github.NewClient(tc)

 // 「github」というオーガニゼイションのpublicリポジトリを取得します
 //opt := &github.RepositoryListByOrgOptions{Type: "public"}
 //repos, _, err := client.Repositories.ListByOrg(ctx,"github", opt)

 // 自分の全てのリポジトリ取得
 repos, _, _ := client.Repositories.List(ctx, "",nil)
 spew.Dump(repos)
}

