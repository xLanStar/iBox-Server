package middlewares

// func ErrorHandler(c *gin.Context) {
// 	c.Next()
// 	if code, hasErr := c.Get("code"); hasErr {
// 		if _, skip := c.Get("skip"); !skip {
// 			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 				"code": code,
// 			})
// 		}
// 	}
// }
