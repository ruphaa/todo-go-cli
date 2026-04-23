# 📝 Todo CLI

A tiny command-line app to keep track of your tasks. Stores everything in a simple Markdown file.

---

## What it does

- Add tasks one-by-one or in a friendly interactive mode
- View your full task list
- Everything lives in a Markdown file you can open anywhere

---

## Build it

```bash
go build -o todo
```

---

## How to use

### Add a single task

```bash
./todo add "buy groceries"
```

### Add tasks interactively

```bash
./todo new
```

Type your tasks and hit **Enter**. Type `q` to quit.

### See your tasks

```bash
./todo list
```

---

## Where tasks are saved

By default, tasks are stored in:

```
/Users/ruphaa/Documents/Second-brain/Todo-go.md
```

You can change this with the `-p` flag:

```bash
./todo -p /path/to/my/todos.md add "walk the dog"
./todo -p /path/to/my/todos.md list
```

---

## Tech

- [Go](https://go.dev)
- [Cobra](https://github.com/spf13/cobra) for CLI commands
