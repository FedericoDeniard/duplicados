# Duplicate Files Finder in Go

A duplicate file finder for Linux developed in Go. It recursively scans directories using **goroutines** for parallel searches, achieving extremely fast execution times even on large file systems.

## 🚀 Goal

The main goal of this project was to practice using **goroutines** and better understand how to leverage Go's concurrency features to create efficient command-line tools (CLI).

## 🔧 Features

- Duplicate file search
- Goroutine concurrency for improved speed
- Specific path exclusion
- Specific file extension exclusion
- Simple terminal usage

## 🛠️ Installation and Usage

### 📦 Installation from Binary

```bash
curl -L https://github.com/FedericoDeniard/duplicados/releases/latest/download/dupes -o /tmp/dupes
chmod +x /tmp/dupes
sudo mv /tmp/dupes /usr/local/bin/dupes
```

### 📦 Installation from Source

Clone the repository:

```bash
git clone https://github.com/FedericoDeniard/duplicados.git
cd duplicados
```

Build and install the program:

```bash
go build -o dist/duplicados src/main.go
sudo mv dist/duplicados /usr/local/bin/duplicados
```

Then you can use the `duplicados` command from anywhere:

```bash
duplicados
```

To uninstall:

```bash
sudo rm /usr/local/bin/duplicados
```

### 🏷️ Available Flags

| Flag            | Description                                                          |
| --------------- | -------------------------------------------------------------------- |
| `-exclude-dirs` | Comma-separated list of directories to exclude from search           |
| `-include-ext`  | Comma-separated list of file extensions to include (e.g., .jpg,.png) |
| `-exclude-ext`  | Comma-separated list of file extensions to exclude                   |
| `-show-hidden`  | Include hidden files and directories in search                       |
| `-use-sha256`   | Use SHA256 for hashing (slower but more secure than default MD5)     |
| `-help`         | Show help message                                                    |

## 🤝 Contributions

Contributions are welcome! If you find a bug or want to suggest improvements, feel free to open an issue or submit a pull request.

---

Built with Go 🦫 by [Federico Deniard](https://github.com/FedericoDeniard)
