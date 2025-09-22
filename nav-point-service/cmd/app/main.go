package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/handler"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/repository"
	"github.com/ZakSlinin/licey-maps-backend/nav-point-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func main() {
	// 1. Конфиг из переменных окружения
	dbURL := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	if dbURL == "" {
		log.Fatal("DB_HOST is not set")
	}

	// 2. Подключение к БД
	db, err := sqlx.Connect("pgx", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	// 3. Настройка миграций
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("ошибка сборки миграций: %v", err)
	}

	// 4. Применяем миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("ошибка миграций: %v", err)
	}
	fmt.Println("Успешная сборка миграций.")

	// 5. Настройка Gin Framework
	r := gin.Default()

	navPointRepository := repository.NewNavPointRepository(db)
	navPointService := service.NewNavPointService(navPointRepository)
	navPointHandler := handler.NewNavPointHandler(navPointService)

	// 6. Запуск сервиса
	r.POST("/create-nav-point", navPointHandler.CreateNavPoint)
	r.GET("/get-nav-point", navPointHandler.GetNavPointByNavPointId)

	log.Printf("Server started on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
