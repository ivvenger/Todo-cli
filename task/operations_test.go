package task

import (
	"errors"
	"path/filepath"
	"testing"
)

func newTestStorage(t *testing.T) *Storage {
	t.Helper()
	path := filepath.Join(t.TempDir(), "tasks.json")
	return NewStorage(path)
}

func mustAdd(t *testing.T, s *Storage, title string) Task {
	t.Helper()
	added, err := s.Add(title)
	if err != nil {
		t.Fatalf("Не удалось добавить задачу %q: %v", title, err)
	}
	return added
}

func TestNextID(t *testing.T) {
	tests := []struct {
		name  string
		tasks []Task
		want  int
	}{
		{"Пустой список", nil, 1},
		{"Одна задача", []Task{{ID: 1}}, 2},
		{"Несколько подряд", []Task{{ID: 1}, {ID: 2}, {ID: 3}}, 4},
		{"С пропусками — берём максимум", []Task{{ID: 1}, {ID: 5}, {ID: 3}}, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextID(tt.tasks); got != tt.want {
				t.Errorf("nextID() = %d, ожидалось %d", got, tt.want)
			}
		})
	}
}

func TestFindIndex(t *testing.T) {
	tasks := []Task{{ID: 1}, {ID: 2}, {ID: 3}}

	tests := []struct {
		name string
		id   int
		want int
	}{
		{"Первый элемент", 1, 0},
		{"Средний элемент", 2, 1},
		{"Последний элемент", 3, 2},
		{"Не найдено", 99, -1},
		{"Пустой id", 0, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndex(tasks, tt.id); got != tt.want {
				t.Errorf("findIndex(%d) = %d, ожидалось %d", tt.id, got, tt.want)
			}
		})
	}
}

func TestFilterByStatus(t *testing.T) {
	tasks := []Task{
		{ID: 1, Done: false},
		{ID: 2, Done: true},
		{ID: 3, Done: false},
		{ID: 4, Done: true},
	}

	tests := []struct {
		name    string
		done    bool
		wantIDs []int
	}{
		{"Только выполненные", true, []int{2, 4}},
		{"Только невыполненные", false, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterByStatus(tasks, tt.done)
			if len(got) != len(tt.wantIDs) {
				t.Fatalf("Получено %d задач, ожидалось %d", len(got), len(tt.wantIDs))
			}
			for i, id := range tt.wantIDs {
				if got[i].ID != id {
					t.Errorf("got[%d].ID = %d, ожидалось %d", i, got[i].ID, id)
				}
			}
		})
	}
}

func TestAdd(t *testing.T) {
	s := newTestStorage(t)

	first, err := s.Add("Первая задача")
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
	if first.ID != 1 {
		t.Errorf("ID первой задачи = %d, ожидалось 1", first.ID)
	}
	if first.Done {
		t.Error("Новая задача не должна быть выполненной")
	}

	second, err := s.Add("вторая задача")
	if err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
	if second.ID != 2 {
		t.Errorf("ID второй задачи = %d, ожидалось 2", second.ID)
	}
}

func TestAddTrimsAndRejectsEmpty(t *testing.T) {
	s := newTestStorage(t)

	if _, err := s.Add("  дела  "); err != nil {
		t.Fatalf("Ошибка: %v", err)
	}
	tasks, _ := s.All()
	if tasks[0].Title != "дела" {
		t.Errorf("Title = %q, ожидалось %q ", tasks[0].Title, "дела")
	}

	for _, empty := range []string{"", "   ", "\t\n"} {
		if _, err := s.Add(empty); !errors.Is(err, ErrEmptyTitle) {
			t.Errorf("Add(%q): ожидалась ErrEmptyTitle, получено %v", empty, err)
		}
	}
}

func TestComplete(t *testing.T) {
	s := newTestStorage(t)
	mustAdd(t, s, "задача")

	if err := s.Complete(1); err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	tasks, _ := s.All()
	if !tasks[0].Done {
		t.Error("После Complete задача должна быть выполненной!")
	}
}

func TestCompleteNotFound(t *testing.T) {
	s := newTestStorage(t)

	err := s.Complete(42)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("Ожидалась ErrTaskNotFound, получено %v", err)
	}
}

func TestDelete(t *testing.T) {
	s := newTestStorage(t)
	mustAdd(t, s, "первая")
	mustAdd(t, s, "вторая")

	if err := s.Delete(1); err != nil {
		t.Fatalf("Ошибка: %v", err)
	}

	tasks, _ := s.All()
	if len(tasks) != 1 {
		t.Fatalf("Осталось %d задач, ожидалась 1", len(tasks))
	}
	if tasks[0].ID != 2 {
		t.Errorf("Осталась задача с ID %d, ожидалось 2", tasks[0].ID)
	}
}

func TestDeleteNotFound(t *testing.T) {
	s := newTestStorage(t)

	err := s.Delete(42)
	if !errors.Is(err, ErrTaskNotFound) {
		t.Errorf("Ожидалась ErrTaskNotFound, получено %v", err)
	}
}
