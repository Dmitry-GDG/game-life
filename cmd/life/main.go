package main

import (
	"context"
	"os"

	"github.com/Dmitry-GDG/game-life/pkg/application"
)

// import (
// 	"context"
// 	"errors"
// 	"log"
// 	"os"

// 	"github.com/Dmitry-GDG/game/pkg/application"
// )

// func main() {
// 	ctx := context.Background()
// 	// Exit завершает программу с заданным кодом
// 	os.Exit(mainWithExitCode(ctx))
// }

// func mainWithExitCode(ctx context.Context) int {
// 	cfg := application.Config{
// 		Width:  10,
// 		Height: 10,
// 	}
// 	app := application.New(cfg)
// 	// Запускаем приложение
// 	if err := app.Run(ctx); err != nil {
// 		switch {
// 		case errors.Is(err, context.Canceled):
// 			log.Println("Processing cancelled.")
// 		default:
// 			log.Println("Application run error", err)
// 		}
// 		// Возвращаем значение, не равное нулю, чтобы обозначить ошибку
// 		return 1
// 	}
// 	// Выход без ошибок
// 	return 0
// }

func main() {
	ctx := context.Background()
	// Exit приводит к завершению программы с заданным кодом.
	os.Exit(mainWithExitCode(ctx))
}

func mainWithExitCode(ctx context.Context) int {
	cfg := application.Config{
		Width:  100,
		Height: 100,
	}
	app := application.New(cfg)

	return app.Run(ctx)
}
