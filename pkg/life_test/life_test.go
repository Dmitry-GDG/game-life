package life_test

import (
	"fmt"
	"testing"
	//"github.com/Dmitry-GDG/game-/pkg/life"
)

func TestNewWorld(t *testing.T) {
	// Задаём размеры сетки
	height := 10
	width := 4
	// Вызываем тестируемую функцию
	world := life.NewWorld(height, width)
	// Проверяем, что в объекте указана верная высота сетки
	if world.Height != height {
		t.Errorf("expected height: %d, actual height: %d", height, world.Height)
	}
	// Проверяем, что в объекте указана верная ширина сетки
	if world.Width != width {
		t.Errorf("expected width: %d, actual width: %d", width, world.Width)
	}
	// Проверяем, что у реальной сетки — заданная высота
	if len(world.Cells) != height {
		t.Errorf("expected height: %d, actual number of rows: %d", height, len(world.Cells))
	}
	// Проверяем, что у каждого элемента — заданная длина
	for i, row := range world.Cells {
		if len(row) != width {
			t.Errorf("expected width: %d, actual row's %d len: %d", width, i, world.Width)
		}
	}

	// В реальных тестах стоит проверять вызов функций с проблемными значениями.
	// Например, в нашей программе вызов функции NewWorld с отрицательными размерами
	// вызовет панику, а это недопустимо.
	// В таком случае наша функция должна вернуть ошибку.
	// Попробуйте настроить проверку этой функции самостоятельно.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			t.Errorf("Panic unacceptable!")
		}
	}()
	world = life.NewWorld(-4, -4)
}
