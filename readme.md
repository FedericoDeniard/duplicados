## 🛠️ Instalación y uso

### 📦 Instalación desde el binario

```bash
curl -L https://github.com/FedericoDeniard/duplicados/releases/tag/v1.0.0/duplicados -o /tmp/duplicados
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
