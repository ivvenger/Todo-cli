package task

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadMissingFileReturnsEmpty(t *testing.T) {
	// Файл, которого точно нет.
	path := filepath.Join(t.TempDir(), "no-such-file.json")
	s := NewStorage(path)

	tasks, err := s.Load()
	if err != nil {
		t.Fatalf("Load несуществующего файла не должен возвращать ошибку, получено: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("ожидался пустой список, получено %d задач", len(tasks))
	}
}

func TestSaveThenLoad(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	s := NewStorage(path)

	want := []Task{
		{ID: 1, Title: "первая", Done: false},
		{ID: 2, Title: "вторая", Done: true},
	}

	if err := s.Save(want); err != nil {
		t.Fatalf("Save вернул ошибку: %v", err)
	}

	got, err := s.Load()
	if err != nil {
		t.Fatalf("Load вернул ошибку: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("загружено %d задач, ожидалось %d", len(got), len(want))
	}
	for i := range want {
		if got[i].ID != want[i].ID || got[i].Title != want[i].Title || got[i].Done != want[i].Done {
			t.Errorf("задача %d = %+v, ожидалось %+v", i, got[i], want[i])
		}
	}
}

func TestLoadInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "broken.json")
	if err := os.WriteFile(path, []byte("это не json"), 0644); err != nil {
		t.Fatal(err)
	}
	s := NewStorage(path)

	if _, err := s.Load(); err == nil {
		t.Error("Load битого JSON должен возвращать ошибку")
	}
}
