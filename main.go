package main

import (
	// "net/http"
	"github.com/aadeliafitri/project-karyawan/controllers/KaryawanController"
	"github.com/aadeliafitri/project-karyawan/controllers/JabatanController"
	"github.com/aadeliafitri/project-karyawan/controllers/GajiTunjanganController"
	"github.com/aadeliafitri/project-karyawan/models"
	"github.com/gin-gonic/gin"
	
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	// r.SetTrustedProxies([]string{"192.168.1.2"})

	// karyawan
	k := r.Group("/api/karyawan")
	{
		k.GET("", KaryawanController.Index)
		k.GET("/page/:page", KaryawanController.Index)
		// k.GET("/search", KaryawanController.TotalGaji)
		k.GET("/:id", KaryawanController.Show)
		k.POST("", KaryawanController.Create)
		k.PUT("/:id", KaryawanController.Update)
		k.DELETE("/:id", KaryawanController.Delete)
	}

	// jabatan
	j := r.Group("/api/jabatan")
	{
		j.GET("", JabatanController.Index)
		j.GET("/page/:page", JabatanController.Index)
		j.GET("/:id", JabatanController.Show)
		j.POST("", JabatanController.Create)
		j.PUT("/:id", JabatanController.Update)
		j.DELETE("/:id", JabatanController.Delete)
	}

	// Gaji Tunjangan
	g := r.Group("/api/gajitunjangan")
	{
		g.GET("", GajiTunjanganController.Index)
		g.GET("/page/:page", GajiTunjanganController.Index)
		g.GET("/:id", GajiTunjanganController.Show)
		g.POST("", GajiTunjanganController.Create)
		g.PUT("/:id", GajiTunjanganController.Update)
		g.DELETE("/:id", GajiTunjanganController.Delete)
	}

	r.Run(":8080")
}