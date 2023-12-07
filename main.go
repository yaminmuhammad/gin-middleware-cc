package main

import "gin/delivery"

func main() {

	delivery.NewServer().Run()
}

/**
TODO:
1. Kita akan pisahkan smeua yang melakukan konfigurasi hardcode, seperti koneksi, host
2. Kita akan pisahkan koneksi nya agar lebih clean lagi
3. Di dalam Server kita akan coba untuk lebih clean lagi
4. Custom response, baik itu create, get single dan get paging
5. Middleware log

Tugas:
1. Buatlah sebuah service untuk menampilkan author berikut dengan task nya
2. Buatlah sebuah service untuk menampilkan task berdasarkan AuthorId

Waktu pengerjaan sampai pukul 13:30


*/

// Branch:
// 1-native-http
// 2-with-gin
// 3-middleware-basic-auth
// 4-middleware-logger
// 5-clean-code-part-1
