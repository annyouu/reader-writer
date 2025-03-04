package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
 _  "github.com/lib/pq"
  	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"echo/config"

)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	var cfg config.EnvConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/hello-world", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world..")
	})

	e.Start("127.0.0.1:8080")
}




// type Response struct {
// 	Message string `json:"message"`
// }

// func main() {
// 	e := echo.New()

// 	// ミドルウェアの設定
// 	e.Use(middleware.Logger())  // 各リクエストをログに記録
// 	e.Use(middleware.Recover()) // パニック発生時に自動回復

// 	//　ルートハンドラ
// 	e.GET("/", func(c echo.Context) error {
// 		// return c.String(http.StatusOK, "Hello, World!")

// 		res := Response{
// 			Message: "Hello world...",
// 		}
// 		return c.JSON(http.StatusOK, res)
// 	}) 

// 	// 意図的にエラーを発生させるエンドポイント
// 	e.GET("/error", func(c echo.Context) error {
// 		panic("意図的なエラーが発生")
// 	})

// 	// サーバー起動
// 	e.Start(":8080")
// }