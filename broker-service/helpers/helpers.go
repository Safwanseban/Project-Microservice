package helpers

type JsonREsponse struct {
	Error   bool   `json:"error"`
	Messege string `json:"messege"`
	Data    any    `json:"data,omitempty"`
}

// func Readjson(c *gin.Context, data any) error {
// 	maxbytes := 1048576 //one megabyte

// 	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxbytes))
// 	if err := c.BindJSON(&c.Request.Body); err != nil {
// 		c.JSON(400, gin.H{
// 			"err": err.Error(),
// 		})
// 		c.JSON(200, gin.H{
// 			"data": c.Request.Body,
// 		})
// 	}

// }
