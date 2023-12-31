# Go and gRPC introduction

## Что такое gRPC

**RPC (CORBA, Sun RPC, DCOM, ...):**

* сетевые вызовы абстрагированы от кода;
* инструменты для кодогенерации;
* использование различных протоколов.

[За что отвечает буква g в gRPC](https://github.com/grpc/grpc/blob/master/doc/g_stands_for.md)

ЯП на которых можно создавать системы с использованием gRPC:

* C#/.NET
* C++
* Dart
* Go
* Java
* Kotlin
* PHP
* Python
* Ruby

## PROTOCOL BUFFERS

**Protobuf (Protocol buffers)** - наиболее используемый в gRPC механизм сериализации структурированных данных.

**Что не так с JSON:**

* Динамическая схема - набор полей строго не регламентирован;
* Медленный парсиг данных;
* Отсутствие типизации - JSON это просто key-value строк;
* Большой размер данных.

**Плюсы PROTOBUF:**

* Меньший размер относительно JSON;
* Типизация;
* Схема данных.

[Типы данных в PROTOBUF](https://protobuf.dev/programming-guides/proto3/#scalar)  

Пример генерации кода из .proto файла и сравнение размеров сериализованных данных: [JSON and PB compare](https://github.com/zexy-swami/go_and_grpc/tree/main/code/json_and_pb_compare)

## gRPC

**Что не так с HTTP:**

* Создает новое соединение на запрос;
* Head of blocking - нет возможности ассинхронной передачи контента;
* Нет server push.

**Приемущества HTTP/2.0:**

* Одно TCP соединение;
* Бинарный протокол - повыщение гибкости передачи данных;
* Сжатие заголовков.

[Сравнение скорости передачи контента в HTTP версиях 1.1 и 2](http://www.http2demo.io)  

Пример запуска gRPC сервера и клиента на языке go: [gRPC example](https://github.com/zexy-swami/go_and_grpc/tree/main/code/grpc_example)
