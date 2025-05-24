# Buscador de Archivos Repetidos en Go

Un buscador de archivos repetidos para Linux desarrollado en Go. Recorre directorios de forma recursiva utilizando **goroutines** para realizar búsquedas paralelas, logrando tiempos de ejecución rapidísimos incluso en sistemas de archivos grandes.

## 🚀 Objetivo

El objetivo principal del proyecto fue practicar el uso de **goroutines** y comprender mejor cómo aprovechar el paralelismo en Go para crear herramientas eficientes desde la línea de comandos (CLI).

## 🔧 Características

- Búsqueda de archivos repetidos
- Concurrencia con goroutines para mayor velocidad
- Exclusión de rutas específicas
- Exclusión de extensiones específicas
- Uso simple desde la terminal

## 🛠️ Instalación y uso

### 📦 Instalación desde el binario

```bash
curl -L https://github.com/FedericoDeniard/duplicados/releases/latest/download/duplicados  -o /tmp/duplicados
chmod +x /tmp/duplicados
sudo mv /tmp/duplicados /usr/local/bin/duplicados
```

### 📦 Instalación desde el código fuente

Cloná el repositorio:

```bash
git clone https://github.com/FedericoDeniard/duplicados.git
cd duplicados
```

Compilá e instalá el programa:

```bash
go build -o dist/duplicados src/main.go
sudo mv dist/duplicados /usr/local/bin/duplicados
```

Luego podés usar el comando `duplicados` directamente desde cualquier lugar:

```bash
duplicados
```

Para desinstalarlo:

```bash
sudo rm /usr/local/bin/duplicados
```

### 🏷️ Flags disponibles

| Flag               | Descripción                                                      |
| ------------------ | ---------------------------------------------------------------- |
| `-exclude`         | Lista de directorios a excluir, separados por comas sin espacios |
| `-file-extensions` | Lista de extensiones a buscar, separadas por comas sin espacios  |
| `-show-hidden`     | Muestra archivos ocultos                                         |
| `-use-sha256`      | Usa SHA256 en lugar de MD5 (más seguro, pero más lento)          |
| `-help`            | Muestra el mensaje de ayuda                                      |

## 🤝 Contribuciones

¡Las contribuciones son bienvenidas! Si encontrás un bug o querés proponer mejoras, sentite libre de abrir un issue o un pull request.

---

Desarrollado con Go 🦫 por [Federico Deniard](https://github.com/FedericoDeniard)
