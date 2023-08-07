# Tarjans_Algorithm
Tarjans_algorithm merupakan algoritma berbasis CLI yang mencari Strongly Connected Component dan Bridge dari sebuah graf menggunakan algoritma tarjans.

# Requirements
1. Golang 20.3
2. Graphviz

# How To Run
1. Buka terminal di folder src
2. Ketikkan ```go build```
3. Jalankan src.exe dengan mengetikkan ```./src``` pada terminal

Setelah itu program akan berjalan dan anda akan diminta untuk memasukkan nama file. File yang ingin dimasukkan harus diletakkan di folder ```tests``` terlebih dahulu. Jika sudah ketikkan nama filenya.
Graf akan tersimpan di folder graph. Buka terminal pada folder ```graph``` lalu ketikkan ```dot -Tsvg -O <namagraf>``` untuk mendapatkan visualisasi setiap graf yang ada pada folder graf. 

![image](https://github.com/MuhLibri/Tarjans_Algorithm/assets/104043362/19461c19-f76f-464f-81ec-94473e14d577)



Visualisasi yang didapatkan akan berupa file dengan ekstensi .svg

![image](https://github.com/MuhLibri/Tarjans_Algorithm/assets/104043362/14f0cd98-4c7e-40b5-b4e4-c21982a4c019)


Bukalah file ekstensi .svg untuk melihat visualisasi

![image](https://github.com/MuhLibri/Tarjans_Algorithm/assets/104043362/10591911-e29a-4344-82d5-bd53bf7c6388)



# About Tarjans Algorithm
Kompleksitas dari algoritma tarjans adalah O(V+E), dimana V adalah jumlah node/simpul dan E adalah jumlah edge/sisi. Algoritma tarjans sendiri sebenarnya dimaksudkan untuk mencari SCC, namun dapat dimodifikasi untuk mendapatkan bridge. Yang dimodifikasi adalah bagian dfsnya dan ada tambahan variabel yaitu parentList. Berbeda dengan mencari scc, saat mencari bridge dfs perlu memperhatikan lowest node yang sedang diproses dengan parentnya.


# Referensi
https://www.w3schools.com/go/go_loops.php
https://youtu.be/ZeDNSeilf-Y
https://youtu.be/Rhxs4k6DyMM

Library yang digunakan disini adalah library untuk memvisualisasikan graphviz
https://github.com/dominikbraun/graph


