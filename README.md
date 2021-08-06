Pongo4echo
=========

Package pongo2echo is a template renderer that can be used with the minimalist Go web framework
https://github.com/labstack/echo it uses the Pongo2 template library
https://github.com/flosch/pongo2

Compatible with pongo version 4

## Here is Compatible with pongo version 4
 [Pongo2echo](https://github.com/stnc/pongo2echo) - pongo2 echo framework stability renderer / Compatible with pongo version 4

# please don't forget to give stars :)

pongo2 is a Django-syntax like templating-language (official website).

## Installation  

`go get "github.com/stnc/pongo4echo"`

Requirements
------------

Requires Echo 4+ or higher and Pongo2.

Usage
-----

Real Example [echo+pongo+gorm+pagination]

https://github.com/stnc/golang-echo-realworld-example-web-app (pongo v2 Compatible)

![Screen](https://raw.githubusercontent.com/stnc/pongo2echo/master/example/echoScreen.png)
 
Basic Example
-------------

```go

package main

import (
	"net/http"

	"github.com/flosch/pongo2/v4"
	"github.com/labstack/echo/v4"
	"github.com/stnc/pongo4echo"
)



//GetAllData all list
func GetAllData(e echo.Context) error {

	posts := []string{
		"Larry Ellison",
		"Carlos Slim Helu",
		"Mark Zuckerberg",
		"Amancio Ortega ",
		"Jeff Bezos",
		" Warren Buffet ",
		"Bill Gates",
		"selman tunç",
	}

	return e.Render(http.StatusOK, "templates/index.html",
		pongo2.Context{"title": "hello echo fw", "posts": posts})
}

func main() {
	e := echo.New()
	e.Renderer = pongo4echo.Renderer{Debug: true} //pongo2 v4 init
	e.Debug = true
	// http://localhost:8888/
	e.GET("/", GetAllData)
	// Start server
	e.Logger.Fatal(e.Start(":8888"))
}


```

HTML 
----------------


```html

<h1> {{ title }}</h1>

{% for post in posts%}
<ul>
  <li>{{post}}</li>
</ul>
{% endfor %}

```


GoDoc
-----


Specials Thanks
-----

https://github.com/siredwin/pongorenderer Version 3

https://machiel.me/post/pongo2-with-echo-or-net-http/
