https://gobyexample.com/

https://go-tour-id2.appspot.com/basics/

# Go Keywords

Berikut adalah daftar seluruh **25 keyword** dalam bahasa pemrograman Go beserta penjelasan fungsinya.

| Keyword | Kategori | Fungsi |
|---------|----------|--------|
| `var` | Deklarasi | Mendeklarasikan satu atau lebih variabel dengan tipe eksplisit maupun implisit. |
| `const` | Deklarasi | Mendeklarasikan satu atau lebih konstanta yang nilainya tidak dapat diubah. |
| `type` | Deklarasi | Mendefinisikan tipe baru, termasuk struct, interface, alias tipe, dll. |
| `func` | Deklarasi | Mendeklarasikan sebuah fungsi atau method. |
| `package` | Paket | Menentukan nama paket dari sebuah file Go. |
| `import` | Paket | Mengimpor paket eksternal atau dari standard library untuk digunakan dalam file. |
| `struct` | Tipe Data | Mendefinisikan tipe data komposit yang mengelompokkan beberapa field dengan tipe berbeda. |
| `interface` | Tipe Data | Mendefinisikan sekumpulan method signature yang harus diimplementasikan oleh suatu tipe. |
| `map` | Tipe Data | Mendefinisikan tipe data key-value (hash table / dictionary). |
| `chan` | Tipe Data | Mendeklarasikan channel yang digunakan untuk komunikasi antar goroutine. |
| `if` | Kontrol Alur | Menjalankan blok kode jika kondisi bernilai `true`. |
| `else` | Kontrol Alur | Menjalankan blok kode alternatif jika kondisi `if` bernilai `false`. |
| `for` | Kontrol Alur | Satu-satunya konstruksi perulangan di Go; dapat berfungsi seperti `while` atau `foreach`. |
| `range` | Kontrol Alur | Digunakan bersama `for` untuk mengiterasi elemen slice, array, map, string, atau channel. |
| `switch` | Kontrol Alur | Mengevaluasi ekspresi dan menjalankan blok `case` yang cocok. |
| `case` | Kontrol Alur | Mendefinisikan kondisi dalam pernyataan `switch` atau `select`. |
| `default` | Kontrol Alur | Blok yang dijalankan jika tidak ada `case` yang cocok dalam `switch` atau `select`. |
| `break` | Kontrol Alur | Menghentikan eksekusi loop atau `switch` secara paksa. |
| `continue` | Kontrol Alur | Melompati iterasi saat ini dan langsung menuju iterasi berikutnya dalam loop. |
| `goto` | Kontrol Alur | Melompat ke label tertentu dalam fungsi yang sama. |
| `fallthrough` | Kontrol Alur | Memaksa eksekusi berlanjut ke `case` berikutnya dalam blok `switch`. |
| `return` | Fungsi | Menghentikan eksekusi fungsi dan mengembalikan nilai (jika ada) ke pemanggil. |
| `defer` | Fungsi | Menunda eksekusi sebuah fungsi hingga fungsi yang memanggilnya selesai (dieksekusi secara LIFO). |
| `go` | Konkurensi | Menjalankan sebuah fungsi secara konkuren sebagai goroutine baru. |
| `select` | Konkurensi | Menunggu dan memilih salah satu operasi channel yang siap dieksekusi; mirip `switch` untuk channel. |

## Catatan

- Go hanya memiliki **25 keyword**, jauh lebih sedikit dibanding bahasa lain seperti Java atau C++.
- Keyword bersifat **case-sensitive** dan seluruhnya ditulis dengan huruf kecil.
- Identifier bawaan seperti `make`, `len`, `cap`, `new`, `append`, `copy`, `delete`, `panic`, `recover`, `close`, `print`, dan `println` **bukan keyword**, melainkan *predeclared identifiers*.

# Go Predeclared Identifiers

Berikut adalah daftar **predeclared identifiers** dalam bahasa pemrograman Go. Berbeda dengan keyword, identifier ini dapat di-*shadow* (ditimpa) oleh deklarasi lokal, namun sangat tidak disarankan.

---

## Tipe Data Bawaan (Built-in Types)

| Identifier | Fungsi |
|------------|--------|
| `bool` | Tipe boolean dengan nilai `true` atau `false`. |
| `byte` | Alias dari `uint8`, merepresentasikan satu byte data. |
| `rune` | Alias dari `int32`, merepresentasikan satu karakter Unicode code point. |
| `string` | Tipe untuk menyimpan rangkaian karakter UTF-8 yang bersifat immutable. |
| `int` | Tipe bilangan bulat bertanda, ukurannya bergantung platform (32-bit atau 64-bit). |
| `int8` | Bilangan bulat bertanda 8-bit (rentang: -128 hingga 127). |
| `int16` | Bilangan bulat bertanda 16-bit (rentang: -32.768 hingga 32.767). |
| `int32` | Bilangan bulat bertanda 32-bit (rentang: -2.147.483.648 hingga 2.147.483.647). |
| `int64` | Bilangan bulat bertanda 64-bit. |
| `uint` | Bilangan bulat tak bertanda, ukurannya bergantung platform (32-bit atau 64-bit). |
| `uint8` | Bilangan bulat tak bertanda 8-bit (rentang: 0 hingga 255). |
| `uint16` | Bilangan bulat tak bertanda 16-bit (rentang: 0 hingga 65.535). |
| `uint32` | Bilangan bulat tak bertanda 32-bit (rentang: 0 hingga 4.294.967.295). |
| `uint64` | Bilangan bulat tak bertanda 64-bit. |
| `uintptr` | Bilangan bulat tak bertanda yang cukup besar untuk menyimpan nilai pointer. |
| `float32` | Bilangan desimal presisi tunggal (IEEE 754 32-bit). |
| `float64` | Bilangan desimal presisi ganda (IEEE 754 64-bit). |
| `complex64` | Bilangan kompleks dengan bagian real dan imajiner bertipe `float32`. |
| `complex128` | Bilangan kompleks dengan bagian real dan imajiner bertipe `float64`. |
| `error` | Interface bawaan untuk merepresentasikan kondisi error; memiliki method `Error() string`. |
| `any` | Alias dari `interface{}`, merepresentasikan tipe apapun (diperkenalkan di Go 1.18). |
| `comparable` | Interface constraint yang menandai tipe yang dapat dibandingkan dengan `==` dan `!=`, digunakan dalam generics. |

---

## Konstanta Bawaan (Built-in Constants)

| Identifier | Fungsi |
|------------|--------|
| `true` | Nilai boolean `true`. |
| `false` | Nilai boolean `false`. |
| `iota` | Konstanta integer yang nilainya di-*reset* ke `0` di setiap blok `const` dan bertambah 1 setiap baris; digunakan untuk membuat enumerasi. |

---

## Nilai Zero Khusus (Zero Value)

| Identifier | Fungsi |
|------------|--------|
| `nil` | Nilai zero untuk tipe pointer, slice, map, channel, function, dan interface. |

---

## Fungsi Bawaan (Built-in Functions)

| Identifier | Fungsi |
|------------|--------|
| `make` | Mengalokasikan dan menginisialisasi objek bertipe `slice`, `map`, atau `chan`; mengembalikan nilai (bukan pointer). |
| `new` | Mengalokasikan memori untuk tipe apapun dan mengembalikan pointer ke nilai zero dari tipe tersebut. |
| `len` | Mengembalikan panjang (jumlah elemen) dari `string`, `array`, `slice`, `map`, atau `chan`. |
| `cap` | Mengembalikan kapasitas dari `slice` atau `chan`. |
| `append` | Menambahkan satu atau lebih elemen ke akhir sebuah `slice`; mengembalikan slice baru. |
| `copy` | Menyalin elemen dari satu `slice` ke `slice` lain; mengembalikan jumlah elemen yang berhasil disalin. |
| `delete` | Menghapus elemen dari `map` berdasarkan key yang diberikan. |
| `close` | Menutup sebuah `channel` untuk menandakan tidak ada nilai lagi yang akan dikirim. |
| `panic` | Menghentikan eksekusi fungsi secara tiba-tiba dan memulai proses *panicking* (mirip throw exception). |
| `recover` | Menangkap nilai panic dan mengembalikan kontrol eksekusi; hanya efektif jika dipanggil di dalam fungsi `defer`. |
| `complex` | Membentuk bilangan kompleks dari dua nilai `float` (bagian real dan imajiner). |
| `real` | Mengekstrak bagian real dari sebuah bilangan kompleks. |
| `imag` | Mengekstrak bagian imajiner dari sebuah bilangan kompleks. |
| `clear` | Menghapus semua elemen dari `map` atau mereset semua elemen `slice` ke nilai zero-nya (diperkenalkan di Go 1.21). |
| `min` | Mengembalikan nilai terkecil dari dua atau lebih argumen bertipe `ordered` (diperkenalkan di Go 1.21). |
| `max` | Mengembalikan nilai terbesar dari dua atau lebih argumen bertipe `ordered` (diperkenalkan di Go 1.21). |
| `print` | Mencetak nilai ke standard error; digunakan untuk keperluan debugging (tidak dijamin ada di semua implementasi). |
| `println` | Mencetak nilai ke standard error dengan newline; digunakan untuk keperluan debugging (tidak dijamin ada di semua implementasi). |

---

## Catatan

- Predeclared identifiers berada di **universe block**, yaitu scope terluar dalam Go.
- Tidak seperti keyword, identifier ini **bisa di-shadow** oleh deklarasi lokal, contoh: `var len = 10` — tapi ini praktik yang buruk.
- Untuk output produksi, gunakan paket `fmt` (seperti `fmt.Println`) sebagai pengganti `print` dan `println`.