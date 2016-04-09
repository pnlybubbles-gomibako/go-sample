package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
)

func root(c *gin.Context) {
  session := sessions.Default(c)
  var count int
  v := session.Get("count")
  if v == nil {
    count = 0
  } else {
    count = v.(int)
    count += 1
  }
  // fmt.Printf("%v\n", count)
  session.Set("count", count)
  session.Save()
  c.JSON(200, gin.H{"count": count})
}

func main() {
  router := gin.Default()
  store := sessions.NewCookieStore([]byte("secret"))
  router.Use(sessions.Sessions("session", store))
  router.GET("/", root)
  router.Run(":8080")
}
