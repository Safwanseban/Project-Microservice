package controllers

import (
	"fmt"
	"net/http"

	"github.com/Safwanseban/Project-Microservices/broker-service/helpers"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

var client = resty.New()

type RequestPayload struct {
	Action      string      `json:"action"`
	AuthPayload AuthPayload `json:"auth,omitempty"`
	LogPayload  LogPayload  `json:"log,omitempty"`
}
type AuthPayload struct {
	Email    string
	Password string
}
type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func HomeBroker(c *gin.Context) {
	data := helpers.JsonREsponse{
		Error:   false,
		Messege: "hit the broker",
	}

	c.JSON(http.StatusAccepted, gin.H{
		"msg": data,
	})

}
func HandleSubmission(c *gin.Context) {

	var payload RequestPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(404, gin.H{
			"err": err.Error(),
		})
	}

	switch payload.Action {
	case "auth":
		Authenticate(c, payload.AuthPayload)
	case "log":
		Logitem(c, payload.LogPayload)
	default:
		c.JSON(404, gin.H{
			"msg": "no actopn specifiedd",
		})
	}

}
func Authenticate(c *gin.Context, data AuthPayload) {

	// req, err := http.NewRequest("POST", "http://localhost:8086/authenticate", bytes.NewBuffer(jsonDat))
	// if err != nil {
	// 	c.JSON(404, gin.H{
	// 		"err": err.Error(),
	// 	})
	// 	return
	// }

	// client := resty.New()
	resp, err := client.R().SetHeader("Accept", "application/json").
		SetBody(data).SetResult(&data).Post("http://localhost:8086/authenticate")
	if err != nil {
		c.JSON(404, gin.H{
			"err": "error senting req",
		})
		return
	}
	fmt.Println(resp.StatusCode())
	if resp.StatusCode() == http.StatusUnauthorized {
		c.JSON(404, gin.H{
			"err": "invalid credentials",
		})
		return
	} else if resp.StatusCode() != http.StatusAccepted {
		c.JSON(404, gin.H{
			"err": "error calling service",
		})

		return
	}
	c.JSON(200, gin.H{
		"data": resp.Body(),
	})

}

func Logitem(c *gin.Context, payload LogPayload) {

	logService_URL := "http://localhost:8082/log"
	resp, err := client.R().SetHeader("Accept", "application/json").SetBody(payload).
		Post(logService_URL)
		fmt.Println(resp.StatusCode())
	if err != nil {

		c.JSON(404, gin.H{
			"err": err.Error(),
		})
		return
	}
	if resp.StatusCode() != http.StatusOK {
		c.JSON(404, gin.H{
			"err": "error calling service",
		})

		return
	}
	c.JSON(200, gin.H{
		"msg": http.StatusAccepted,
	})

}

// client := &http.Client{}

// response, err := client.Do(req)
// if err != nil {
// 	c.JSON(404, gin.H{
// 		"err": err.Error(),
// 	})
// 	return
// }
// defer response.Body.Close()

// if response.StatusCode == http.StatusUnauthorized {
// 	c.JSON(404, gin.H{
// 		"err": "invalid credentials",
// 	})
// 	return
// } else if response.StatusCode != http.StatusAccepted {
// 	c.JSON(404, gin.H{
// 		"err": "error calling service",
// 	})

// 	return
// }
// if err := c.ShouldBindJSON(&response.Body); err != nil {
// 	c.JSON(500, gin.H{
// 		"err": err.Error(),
// 	})
// 	return
// }
// c.JSON(200, gin.H{
// 	"data": response.Body,
// })
// }
