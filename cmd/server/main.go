package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfcheca/Checa_Budai_FinalBack3/cmd/server/handler"
	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/odontologo"
	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/paciente"
	"github.com/jfcheca/Checa_Budai_FinalBack3/pkg/store"

	//	"github.com/jfcheca/Checa_Budai_FinalBack3/internal/domain"
	//	"github.com/joho/godotenv"
	"io/ioutil"
	"strings"
)

func main() {
/*	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env:", err)
	}*/

	// Abrir una conexión temporal a MySQL para ejecutar comandos administrativos
	db, err := sql.Open("mysql", "root:122481624@(localhost:3306)/clinica")
	if err != nil {
		log.Fatal("Error al conectar con MySQL:", err)
	}
	defer db.Close()

	// Eliminar la base de datos 'clinica' si ya existe
	_, err = db.Exec("DROP DATABASE IF EXISTS clinica")
	if err != nil {
		log.Fatal("Error al eliminar la base de datos 'clinica':", err)
	}

	// Crear la base de datos 'clinica'
	_, err = db.Exec("CREATE DATABASE clinica")
	if err != nil {
		log.Fatal("Error al crear la base de datos 'clinica':", err)
	}

	// Conectar a la base de datos 'clinica'
	bd, err := sql.Open("mysql", "root:122481624@(localhost:3306)/clinica")
	if err != nil {
		log.Fatal("Error al conectar con la base de datos 'clinica':", err)
	}
	defer bd.Close()

	// Cargar contenido del archivo schema.sql
	sqlFile, err := ioutil.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error al leer el archivo schema.sql:", err)
	}

	// Dividir el contenido en sentencias SQL individuales
	sqlStatements := strings.Split(string(sqlFile), ";")

	// Ejecutar cada sentencia SQL en el archivo schema.sql
	for _, statement := range sqlStatements {
		// Limpiar la sentencia SQL
		cleanedStatement := strings.TrimSpace(statement)
		if cleanedStatement == "" {
			continue
		}

		_, err := bd.Exec(cleanedStatement)
		if err != nil {
			log.Fatal("Error al ejecutar la sentencia SQL:", err)
		}
	}
	

	// Crear el repositorio, el servicio y el controlador con el almacenamiento configurado


		// Crear el almacenamiento SQL con la base de datos 'clinica'
	storage := store.NewSqlStore(bd)
	repo := odontologo.NewRepository(storage)
	service := odontologo.NewService(repo)
	odontoHandler := handler.NewProductHandler(service)

	// Configurar el enrutador Gin
	r := gin.Default()

	// Rutas para el manejo de odontólogos
	odontologos := r.Group("/odontologos")
	{

		odontologos.POST("/crear", odontoHandler.Post())

	}

	// Ejecutar el servidor en el puerto 8080
	r.Run(":8080")
}