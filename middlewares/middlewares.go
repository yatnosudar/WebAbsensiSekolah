package middlewares

import "github.com/labstack/echo/v4/middleware"

var LoginGuru = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("guru"),
})

var LoginAdmin = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("admin"),
})
var LoginSiswa = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("siswa"),
})
