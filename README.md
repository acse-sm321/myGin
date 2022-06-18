# myGin
## About
A light web framework inspired by Gin.

```
|── go.mod
├── main.go
└── myGin
    ├── context.go
    ├── go.mod
    ├── myGin.go
    └── router.go
```
## Installation and use
```
$ git clone https://github.com/acse-sm321/myGin.git
```

In the main program of the application:
```
         // Initilization
	r := myGin.New()
	
	// register the methods and then run
	r.GET("/", func(c *myGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello myGin</h1>")
	})

	r.GET("/hello", func(c *myGin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	// now implement context for Query + PostForm
	r.POST("/login", func(c *myGin.Context) {
		c.JSON(http.StatusOK, myGin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
```

## Compile and run
```
# run directly
$ go run main.go
# Alternative: run the compiled binary 
$ go install main.go
$ cd bin/
$ ./main
$ curl http://localhost:9999/
```



