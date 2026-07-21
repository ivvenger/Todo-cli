# todo-cli

[![CI](https://github.com/ivvenger/todo-cli/actions/workflows/ci.yml/badge.svg)](https://github.com/ivvenger/todo-cli/actions/workflows/ci.yml)

Простой менеджер задач для командной строки на Go. 

## Возможности

- Добавление, просмотр, отметка и удаление задач
- Хранение списка в JSON-файле (по умолчанию `tasks.json`)
- Понятные сообщения об ошибках и корректные коды выхода
- Путь к файлу настраивается флагом `--file`

## Установка

```bash
git clone https://github.com/ivvenger/todo-cli.git
cd todo-cli
go build -o todo .
```

## Использование

```bash
# добавить задачу
./todo add "купить хлеб"

# показать все задачи
./todo list
# [ ] 1: купить хлеб

# отметить выполненной
./todo done 1
./todo list
# [x] 1: купить хлеб

# удалить задачу
./todo rm 1

# использовать другой файл
./todo --file ~/tasks.json add "погулять с собакой"
```

## Разработка

```bash
make test    # прогнать тесты
make cover   # тесты с покрытием
make lint    # golangci-lint
make build   # собрать бинарник
```

## Структура проекта

```
todo-cli/
├── cmd/            # команды CLI (add, list, done, rm) на cobra
├── task/           # модель Task, хранилище и операции + тесты
├── main.go         # точка входа
├── Makefile
└── .golangci.yml
```

## Лицензия

MIT