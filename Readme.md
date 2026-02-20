# dbkeeper

Developer-first database backup guardrail for PostgreSQL and MySQL.

---

## Overview

dbkeeper adalah CLI kecil yang fokus membantu developer dan tim kecil menjaga backup database tetap prediktif dan terobservasi.

Tool ini **tidak** mengeksekusi backup sendiri; alih-alih, tool ini:

- Mengelola konfigurasi YAML sederhana untuk koneksi database dan pengaturan backup.
- Menginisialisasi workspace lokal di direktori home user.
- Menyediakan wizard konfigurasi interaktif.
- Memvalidasi dan menampilkan konfigurasi.
- Mencatat log terpusat ke file dan menyediakan log viewer berwarna.
- Memverifikasi ketersediaan tools client database yang dibutuhkan (`pg_dump`, `mysqldump`) lewat perintah `doctor`.

dbkeeper didesain untuk berjalan dekat dengan database Anda, menggunakan client resmi dan file biasa sehingga mudah diintegrasikan ke workflow backup/restore Anda tanpa lock-in.

---

## Features

### Commands

Seluruh command diimplementasikan dengan [Cobra](https://github.com/spf13/cobra).

- `dbkeeper`  
  - Root command.
  - Menginisialisasi logger dan mendelegasikan ke subcommands.

- `dbkeeper init`  
  - Menginisialisasi workspace dbkeeper di `~/.dbkeeper/`.
  - Membuat:
    - Base directory: `~/.dbkeeper/`
    - Log directory: `~/.dbkeeper/logs/`
    - Backup directory: `~/.dbkeeper/backups/`
    - Config file: `~/.dbkeeper/config.yaml` (jika belum ada).
  - Menulis konfigurasi default jika `config.yaml` belum ada.
  - Mencatat event inisialisasi ke log file.

- `dbkeeper setup`  
  - Menjalankan wizard interaktif untuk (re)generate `~/.dbkeeper/config.yaml`.
  - Menanyakan:
    - Database type (`postgres` atau `mysql`).
    - Host.
    - Port.
    - Database name.
    - Username.
    - Password.
    - Backup retention (hari).
    - Backup directory.
  - Meng-overwrite config yang sudah ada setelah memberi peringatan.
  - Mencatat seluruh nilai konfigurasi ke log (password tidak dicetak eksplisit di log).

- `dbkeeper config`  
  - Parent command untuk subcommand terkait konfigurasi.
  - Saat ini digunakan sebagai namespace untuk:
    - `dbkeeper config show`
    - `dbkeeper config validate`

- `dbkeeper config show`  
  - Membaca `~/.dbkeeper/config.yaml`.
  - Mendukung:
    - `--show-json` untuk menampilkan konfigurasi dalam format JSON terformat.
    - `--show-password` untuk menampilkan password secara plain text.
  - Jika `--show-password` **tidak** diberikan, password diganti menjadi `(secret_key)` sebelum ditampilkan.
  - Output non-JSON menggunakan warna dan dibagi per section (“Database Config”, “Backup Config”).

- `dbkeeper config validate`  
  - Mem-validasi file konfigurasi.
  - Mengecek field-field berikut harus ada dan tidak kosong:
    - `database.host`
    - `database.port`
    - `database.user`
    - `database.password`
    - `database.name`
  - Mencetak pesan error untuk field yang hilang.
  - Mencetak `Configuration is valid` jika semua aturan terpenuhi.
  - Mengembalikan error non-nil jika konfigurasi tidak valid.

- `dbkeeper logs`  
  - Menampilkan log dbkeeper.
  - Membaca log dari `~/.dbkeeper/logs/dbkeeper.log`.
  - Mendukung:
    - `--level, -l` (default: `INFO`)  
      - Filter pesan berdasarkan severity dengan prioritas:
        - `DEBUG` (1), `INFO` (2), `WARN` (3), `SUCCESS` (4), `ERROR` (5).
      - `--level=ERROR` hanya menampilkan log `ERROR`.
      - `--level=INFO` menampilkan `INFO`, `WARN`, `SUCCESS`, dan `ERROR`, dst.
    - `--tail, -t` (default: `100`)  
      - Hanya menampilkan N baris terakhir setelah filter.
    - `--no-color, -c`  
      - Menonaktifkan warna pada output.
  - Mapping warna (saat warna aktif):
    - `DEBUG`: bright black.
    - `INFO`: bright blue.
    - `WARN`: bright yellow.
    - `ERROR`: bright red.
    - `SUCCESS`: bright green.

- `dbkeeper doctor`  
  - Mengecek kehadiran dan versi tools eksternal yang digunakan untuk backup:
    - `pg_dump` (PostgreSQL).
    - `mysqldump` (MySQL).
  - Untuk setiap tool:
    - Mencari executable di `PATH`.
    - Menjalankan `<tool> --version` dengan timeout.
    - Mem-parse string versi yang ramah dibaca.
    - Mencetak pesan success jika ditemukan.
    - Mencetak blok failure dengan hint instalasi per OS jika tidak ditemukan.
  - Mencetak header dan footer di awal/akhir pemeriksaan.

- `dbkeeper version`  
  - Mencetak versi CLI:
    - `dbkeeper v0.1.0`.

### Flags

Berikut flags yang diimplementasikan:

- `dbkeeper logs`:
  - `--level, -l string`  
    Filter level log (`DEBUG`, `INFO`, `WARN`, `ERROR`, `SUCCESS`); default `INFO`.
  - `--tail, -t int`  
    Jumlah baris log yang ditampilkan setelah filter; default `100`.
  - `--no-color, -c`  
    Menonaktifkan warna.

- `dbkeeper config show`:
  - `--show-json`  
    Menampilkan konfigurasi dalam format JSON.
  - `--show-password`  
    Menampilkan password secara plain text (tanpa masking).

- Root command:
  - `--toggle, -t`  
    Placeholder flag yang saat ini belum punya efek fungsional.

---

## Installation

### Prerequisites

- Go (disarankan Go 1.20+; module mendeklarasikan `go 1.25.0`).
- Git (untuk clone repository).

### Build dari Source

```bash
git clone https://github.com/jonalphabert/db-guard.git
cd db-guard

go build -o dbkeeper
```

Perintah di atas menghasilkan binary `dbkeeper` (atau `dbkeeper.exe` di Windows) di direktori saat ini.

### Menjadikan Binary Global

Unix-like (Linux/macOS):

```bash
go build -o dbkeeper
sudo mv dbkeeper /usr/local/bin/
```

Windows (PowerShell):

```powershell
go build -o dbkeeper.exe
# Pindahkan dbkeeper.exe ke folder yang ada di PATH Anda, misalnya C:\Tools
```

Verifikasi instalasi:

```bash
dbkeeper version
```

---

## Usage

Contoh penggunaan CLI berdasarkan command yang ada.

### Inisialisasi Workspace

```bash
dbkeeper init
```

- Membuat `~/.dbkeeper/` dengan:
  - `logs/dbkeeper.log`
  - `backups/`
  - `config.yaml` (jika belum ada; berisi nilai default).

### Setup Interaktif

```bash
dbkeeper setup
```

- Menjalankan wizard interaktif di terminal untuk menentukan:
  - Tipe database (`postgres` / `mysql`).
  - Parameter koneksi.
  - Backup retention (hari).
  - Direktori backup.
- Menyimpan YAML ke `~/.dbkeeper/config.yaml`, meng-overwrite file lama (dengan warning).

### Pemeriksaan Sistem (Doctor)

```bash
dbkeeper doctor
```

- Menjalankan pengecekan sistem untuk:
  - `pg_dump`
  - `mysqldump`
- Melaporkan:
  - Apakah tool tersedia di `PATH`.
  - String versi yang ter-parse (jika ada).
  - Hint instalasi per platform jika tidak ditemukan.

### Melihat Log

```bash
dbkeeper logs --level=ERROR
```

- Membaca `~/.dbkeeper/logs/dbkeeper.log`.
- Memfilter hanya baris dengan severity `ERROR`.
- Menampilkan maksimal 100 baris terbaru (default).

Contoh lain:

```bash
dbkeeper logs                 # default INFO+ (INFO, WARN, SUCCESS, ERROR), 100 baris terakhir, berwarna
dbkeeper logs --tail=200      # INFO+ untuk 200 baris terakhir
dbkeeper logs --level=DEBUG   # semua level
dbkeeper logs --no-color      # tanpa warna
```

### Menampilkan Konfigurasi

```bash
dbkeeper config show
```

- Menampilkan tampilan konfigurasi yang sudah di-beautify dan berwarna.
- Password akan dimasking menjadi `(secret_key)` secara default.

Menampilkan sebagai JSON (password tetap dimasking kecuali diminta sebaliknya):

```bash
dbkeeper config show --show-json
```

Menampilkan termasuk password (gunakan dengan hati‑hati):

```bash
dbkeeper config show --show-password
dbkeeper config show --show-json --show-password
```

### Validasi Konfigurasi

```bash
dbkeeper config validate
```

- Membaca `~/.dbkeeper/config.yaml`.
- Memvalidasi field yang wajib.
- Mencetak:
  - Pesan error jika ada field yang belum diisi.
  - `Configuration is valid` jika semua aturan terpenuhi.

---

## Doctor Command Details

Perintah `doctor` diimplementasikan di `cmd/doctor.go` dan package `internal/doctor`.

### Apa yang Dicek

Untuk setiap `pg_dump` dan `mysqldump`:

1. **Kehadiran executable**
   - Menggunakan `exec.LookPath(<tool>)` untuk menentukan apakah tool tersedia di `PATH`.

2. **Deteksi versi**
   - Menjalankan `<tool> --version` dengan timeout 2 detik.
   - Mem-parse output:
     - `pg_dump`: mengambil PostgreSQL version (mis. `16.3`).
     - `mysqldump`: mengambil versi dari pola `Ver <x.y.z>`.
     - Fallback: generic version parsing atau `unknown`.

3. **Feedback ke user**
   - Jika ditemukan:
     - Mencetak baris success berwarna hijau, mis. `✓ pg_dump found (v16.3)`.
   - Jika tidak ditemukan:
     - Mencetak baris failure berwarna merah +:
       - Penjelasan tujuan tool (mis. “Required for PostgreSQL backups.”).
       - Hint instalasi per platform.

### Tools yang Dibutuhkan

- `pg_dump`
  - Dibutuhkan untuk backup PostgreSQL.
- `mysqldump`
  - Dibutuhkan untuk backup MySQL.

### Hint Instalasi per OS

Ditampilkan otomatis saat tool tidak ditemukan:

- `pg_dump`
  - Windows:  
    `Download PostgreSQL installer and select "Command Line Tools"`  
    `https://www.postgresql.org/download/windows/`
  - macOS:  
    `brew install postgresql`
  - Ubuntu:  
    `sudo apt install postgresql-client`

- `mysqldump`
  - Windows:  
    `Install MySQL Community Server and select "Client Tools"`  
    `https://dev.mysql.com/downloads/installer/`
  - macOS:  
    `brew install mysql`
  - Ubuntu:  
    `sudo apt install mysql-client`

---

## Configuration

Konfigurasi dimodelkan di package `internal/models` dan dibaca menggunakan YAML.

### Lokasi

Semua helper konfigurasi menggunakan base directory yang sama:

- Base directory: `~/.dbkeeper/`
- Config file: `~/.dbkeeper/config.yaml`

Ini diresolve lewat:

- `BaseDir()` → `~/.dbkeeper`
- `ConfigPath()` → `~/.dbkeeper/config.yaml`

### Struktur

YAML dipetakan ke tipe Go berikut:

```yaml
database:
  type: postgres
  host: localhost
  port: 5432
  name: example_db
  user: example_user
  password: example_password

backup:
  retention: 3
  dir: ./backups
```

- `database` → `DatabaseConfig`
  - `type` (string)
  - `host` (string)
  - `port` (int)
  - `name` (string)
  - `user` (string)
  - `password` (string)
- `backup` → `BackupConfig`
  - `dir` (string)
  - `retention` (int)

`dbkeeper init` menulis konfigurasi default dengan `backup.path` ke `~/.dbkeeper/backups`, sedangkan `setup` interaktif menghasilkan field `backup.dir` berdasarkan input user. YAML tags di struct mengarahkan ke key yang sesuai saat parsing.

### Default Configuration

Saat `dbkeeper init` dijalankan dan tidak ada config, tool menulis:

```yaml
database:
  type: postgres
  host: localhost
  port: 5432
  name: example_db
  user: example_user
  password: example_password

backup:
  path: ~/.dbkeeper/backups
  retention: 3
```

Pemanggilan `dbkeeper setup` berikutnya dapat meng-overwrite file ini dengan konfigurasi baru yang berisi field `backup.dir` dan nilai sesuai input user.

### Password Masking

Pada `dbkeeper config show`:

- Secara default, password diubah menjadi `(secret_key)` sebelum ditampilkan.
- Flag `--show-password` menonaktifkan masking dan menampilkan nilai asli.

Berlaku baik untuk output berwarna maupun JSON.

### JSON Output

`dbkeeper config show --show-json`:

- Melakukan marshal struct konfigurasi ke JSON menggunakan `encoding/json`.
- Menampilkan dengan indentasi.
- Menghormati masking password kecuali `--show-password` digunakan.

---

## Logging

Logging diimplementasikan oleh package `internal/logger` dan `internal/logviewer`.

### Lokasi Log

- File log: `~/.dbkeeper/logs/dbkeeper.log`
- Dibuat dan dibuka oleh root command saat startup:
  - Resolve home directory user.
  - Pastikan `~/.dbkeeper/logs/` ada.
  - Membuka/membuat `dbkeeper.log` dalam mode append.

### Format Log

Logger menggunakan package standar `log` dengan prefix:

- `[INFO] 2006/01/02 15:04:05 message`
- `[WARN] 2006/01/02 15:04:05 message`
- `[ERROR] 2006/01/02 15:04:05 message`
- `[SUCCESS] 2006/01/02 15:04:05 message`

Package `internal/logger` menyediakan:

- `Info(format, ...interface{})`
- `Warn(format, ...interface{})`
- `Error(format, ...interface{})`
- `Success(format, ...interface{})`

### Melihat Log

Gunakan command `logs`:

```bash
dbkeeper logs
dbkeeper logs --level=ERROR
dbkeeper logs --tail=200
dbkeeper logs --no-color
```

Log viewer akan:

- Membaca seluruh isi file log.
- Memfilter berdasarkan severity dengan priority map:
  - `DEBUG` ≤ `INFO` ≤ `WARN` ≤ `SUCCESS` ≤ `ERROR`.
- Mengambil N baris terakhir (`Tail`) setelah filter.
- Mencetak tiap baris dengan label severity berwarna (kecuali `--no-color`).

### Colored Output dan `--no-color`

Warna diatur menggunakan library `fatih/color`:

- Setiap level memiliki warna berbeda (lihat mapping di atas).
- Saat warna aktif:
  - Format: `LEVEL: <message_after_prefix>`.
- Saat `--no-color`:
  - Format: `LEVEL: <message_after_prefix>` tanpa ANSI codes.

Flag `--no-color` hanya memengaruhi output `logs`, bukan format penyimpanan file log itu sendiri.

---

## Architecture Overview

### Struktur CLI Berbasis Cobra

- Entry point: `main.go` memanggil `cmd.Execute()`.
- Root command: didefinisikan di `cmd/root.go` dengan:
  - `Use: "db-guard"` (nama internal; binary biasanya diberi nama `dbkeeper` saat build).
  - Deskripsi panjang yang menjelaskan filosofi (backup yang benar‑benar bisa direstore).
  - Hook inisialisasi (`cobra.OnInitialize(initLogger)`) untuk setup logging sekali di awal.

Subcommand diregistrasi lewat `init()` masing‑masing:

- `cmd/init.go` → `init`.
- `cmd/setup.go` → `setup`.
- `cmd/logs.go` → `logs`.
- `cmd/doctor.go` → `doctor`.
- `cmd/version.go` → `version`.
- `cmd/config/config.go` → `config` (parent).
  - `cmd/config/show.go` → `config show`.
  - `cmd/config/validate.go` → `config validate`.

### Internal Packages

- `internal/logger`
  - Logging ke file dengan logger per level.

- `internal/logviewer`
  - Pembacaan file, tailing, filter severity, dan printing dengan/ tanpa warna.

- `internal/config`
  - Resolusi path config.
  - Pembacaan dan parsing YAML.
  - Validasi konfigurasi (`ValidateConfigRule`, `ValidateConfigFile`, `ValidateConfig`).
  - Penampilan konfigurasi ke user (`Show`, `ShowAsJson`).

- `internal/setup`
  - Path filesystem (`BaseDir`, `ConfigPath`).
  - Inisialisasi workspace (`Init`).
  - Wizard konfigurasi interaktif (`RunWizard` dengan `survey/v2`).
  - Rendering config (`RenderConfig`).
  - Orkestrasi setup tingkat tinggi (`Run`).

- `internal/doctor`
  - Pengecekan executable, parsing versi, dan output `doctor`.

- `internal/models`
  - Struktur data `Config`, `DatabaseConfig`, `BackupConfig`.

### Format Data & Konfigurasi

- File konfigurasi YAML (`config.yaml`):
  - Diparse dengan `gopkg.in/yaml.v3`.
  - Terdiri atas section `database` dan `backup`.

- Output konfigurasi JSON:
  - Dihasilkan dari struct `Config` di memori menggunakan `encoding/json`.

### Penggunaan `fatih/color`

- `internal/config/show.go`
  - Mewarnai section dan key saat menampilkan konfigurasi.

- `internal/logviewer/color.go`
  - Mendefinisikan warna per level log.

- `internal/doctor/output.go`
  - Menggunakan warna untuk menonjolkan success (hijau), error/warning (merah/kuning) dan format hint instalasi.

Struktur kode memisahkan colorization dari logika inti sehingga mudah dites dan memungkinkan opsi seperti `--no-color` di log viewer.

---

## Roadmap (Konseptual)

Bagian ini belum diimplementasikan di kode; ini hanyalah arah pengembangan yang natural berdasarkan struktur yang ada.

- **Backup execution**  
  Menambahkan command `backup` yang:
  - Membaca konfigurasi.
  - Memanggil `pg_dump` atau `mysqldump` dengan flags yang tepat.
  - Menyimpan file dump ke `~/.dbkeeper/backups/` atau `backup.dir`.
  - Mencatat metadata dan status success/failure ke log.

- **Restore verification**  
  Menambahkan command `verify` atau `restore-test` yang:
  - Menyalakan instance database sementara.
  - Me-restore backup terbaru.
  - Menjalankan health check sederhana atau query yang didefinisikan user.
  - Mencatat status pass/fail.

- **Scheduler integration**  
  Integrasi dengan scheduler eksternal (cron, systemd timer, Windows Task Scheduler) dengan:
  - Contoh definisi job.
  - Mungkin helper `schedule` untuk mencetak contoh entri cron.

- **Multi-database support**  
  Memperluas konfigurasi dan doctor check untuk engine lain, dengan tetap mempertahankan dukungan utama untuk `postgres` dan `mysql`.
