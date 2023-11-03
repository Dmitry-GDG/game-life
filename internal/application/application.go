package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

type Config struct {
	Width  int
	Height int
}

type Application struct {
	Cfg Config
}

func New(config Config) *Application {
	return &Application{
		Cfg: config,
	}
}

// func (a *Application) Run(ctx context.Context) error {
// 	// Объект для хранения текущего состояния сетки
// 	currentWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
// 	// Объект для хранения очередного состояния сетки
// 	nextWorld := life.NewWorld(a.Cfg.Height, a.Cfg.Width)
// 	// Заполняем сетку на 30%
// 	currentWorld.RandInit(30)
// 	for {
// 		// Здесь мы можем записывать текущее состояние  — например, в очередь сообщений. Для нашего примера просто выводим на экран
// 		fmt.Println(currentWorld)
// 		life.NextState(currentWorld, nextWorld)
// 		currentWorld = nextWorld
// 		// Проверяем контекст
// 		select {
// 		case <-ctx.Done():
// 			return ctx.Err() // Возвращаем причину завершения
// 		default: // По умолчанию делаем паузу
// 			time.Sleep(100 * time.Millisecond)
// 			break
// 		}
// 		// Очищаем экран
// 		fmt.Print("\033[H\033[2J")
// 	}
// }

func (a *Application) Run(ctx context.Context) int {
	// Создание логгера с настройками для production
	logger := setupLogger()

	shutDownFunc, err := server.Run(ctx, logger, a.Cfg.Height, a.Cfg.Width)
	if err != nil {
		logger.Error(err.Error())

		return 1 // вернем код для регистрация ошибки системой
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-c
	cancel()
	//  завершим работу сервера
	shutDownFunc(ctx)

	return 0

}

// настройки логгера
func setupLogger() *zap.Logger {
	// Настройка конфигурации логгера
	config := zap.NewProductionConfig()

	// Уровень логирования
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// Настройка логгера с конфигом
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Ошибка настройки логгера: %v\n", err)
	}

	return logger
}
