package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type RestController struct {

}

func NewRestController() * RestController {
	ret := &RestContoller {}
	go ret.Run()
	return ret
}

func (*RestContoller) handle(command string, dto *DTO) {
  return;
}

func PostCreateInput(c *gin.Context) {

  c.JSON(http.StatusOk, gin.H {
    "message": "create"
  })
}

func PostCreateOutput(c *gin.Context) {

  c.JSON(http.StatusOk, gin.H {
    "message": "create"
  })
}

func (controller *RestController) Run() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("/createInput", PostCreateInput)
		v1.POST("/createOutput", PostCreateOutput)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

