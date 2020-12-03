package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/home", showIndexPage)
	router.GET("/baas", showBaaSPage)

	router.GET("/signin", showLoginPage)
	router.POST("/signin", performLogin)
	router.GET("/logout", logout)
	router.GET("/signup", showRegistrationPage)
	router.POST("/signup", register)

	router.GET("/news", showNewsPage)
	router.GET("/bots", showBotsPage)
	router.GET("/tasks", showTaskPage)
	router.POST("/tasks", addTask)
	router.GET("/user", showUserPage)
	router.POST("/user", postUser)

	router.GET("/interval/:username", respInterval)
	router.GET("/status/:username", respStatus)
	router.POST("/docking", DOCKING)

	router.GET("/hidden/oh/yeah/omg/nice/godbuntu/:count", createSerial)
}