package http

import (
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
)

var pageTemplate, _ = template.New("").Parse(`
<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>{{ .Title }}- {{ .Code }}</title>
  <style type="text/css">
    *{margin:0;padding:0;box-sizing:border-box;color:#f1f1f1}body{font-family:Verdana,Geneva,Tahoma,sans-serif;background:#eeaeca;background:radial-gradient(circle,rgba(238,174,202,1) 0,rgba(148,187,233,1) 100%)}.container{height:100vh;width:80vw;margin:0 auto;display:flex;align-items:center;justify-content:center;flex-direction:column;text-align:center}.container .num{font-size:8rem;margin-bottom:40px}.container .stg{font-size:3rem;margin-bottom:40px;display:none;animation:.7s ease-in-out show}@keyframes show{0%{opacity:0}100%{opacity:1}}
  </style>
</head>
<body>
  <!-- partial:index.partial.html -->
  <div class="container">
      <div class="display">
        <h1 class="num"></h1>
        <h1 class="stg"></h1>
    </div>
    </div>
  <!-- partial -->
  <script type="text/javascript">
    // Declare the Elements 
    const dispNum = document.querySelector(".display .num");
    const dispErr = document.querySelector(".container .stg");
    window.onload = function () {
        function showNum () {
            const randomNum = Math.floor(Math.random() * 1000);
            const randomStr = randomNum.toString()
            dispNum.textContent = randomStr
        }
        var interval =  setInterval(showNum , 500);
        setTimeout(()=> {
            clearInterval(interval);
            dispNum.textContent = "{{ .Code }}";
            if (dispNum.textContent == "") {
              dispNum.textContent = "404"
            }
            dispErr.style.display = "block";
            dispErr.textContent = "{{ .Message }}";
            if (dispErr.textContent == "") {
              dispErr = "呀!这个页面走丢了";
            }
        }, 4000);
    }
  </script>
</body>
</html>
`)

func NoRoute(e *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.ContentType() {
		case "application/json":
			c.AbortWithStatusJSON(http.StatusNotFound, errors.NotFound("PageNotFound", "page not found."))
		default:
			c.Abort()
			c.HTML(http.StatusNotFound, "", gin.H{})
		}
	}
}

func NoMethod(data []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.ContentType() {
		case "application/json":
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, errors.NotFound("MethodNotAllowed", "request method not allowed."))
		default:
			c.Abort()
			c.HTML(http.StatusNotFound, "", gin.H{})
		}
	}
}
