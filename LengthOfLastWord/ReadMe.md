# LeetCode 58 - Length of Last Word (Go)

Bu LeetCode 58: Length of Last Word problemi için  
Go (Golang) ile yazılmış temiz ve verimli bir çözüm içermektedir.

## Problem Özeti
Kelimeler ve boşluklardan oluşan bir `s` string’i verilmektedir.  
Bu string içindeki **son kelimenin uzunluğu** istenmektedir.

Bir kelime, boşluk karakteri içermeyen en uzun karakter dizisidir.

## Yaklaşım
- String’in sonundan itibaren dolaşılır
- Sondaki boşluklar atlanır
- Boşluk görülene kadar karakterler sayılır

Bu yöntem ek bellek kullanmadan çalışır.

## Karmaşıklık
- Zaman Karmaşıklığı: O(n)
- Alan Karmaşıklığı: O(1)

## Çalıştırma
```bash
go run main.go
