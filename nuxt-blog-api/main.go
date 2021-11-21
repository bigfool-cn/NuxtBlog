package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "nuxt-blog-api/configs"
  "nuxt-blog-api/routers"
  "strconv"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	r := routers.InitRouter()
	if err := r.Run(configs.Configs.App.Host + ":" + strconv.Itoa(configs.Configs.App.Port)); err != nil {
		log.Fatalf("app run failed: %v",err)
	}

}
