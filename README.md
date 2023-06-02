# Тестовое задание на позицию Golang Developer at Softweather

## Краткое описание приложения:
```
Необходимо реализовать CLI утилиту и HTTP-сервер.

CLI утилита должна принимать на вход строку (без пробелов) и url, куда надо слать запрос.
Оба аргумента обязательны. HTTP-сервер обслуживает только один endpoint - /api/substring.
В этот endpoint через CLI надо отправлять строки без пробелов. 
При получении данных, HTTP-сервер должен найти самую длинную подстроку без повторяющихся символов и затем отправить ответ, который выведет CLI утилита. 

Пример: abcabcbb -> abc, bbbb -> b, pwwkew -> wke.
```

## Дополнительные (необязательные) требования функциональности:
```
1. Makefile
2. Unit-тесты
3. Docker
4. Без использования сторонних пакетов
```

## Ендпоинты:
```
POST запросом отправить строку по адресу /api/substring
```

## Запуск проекта:

### Через main.go:
```
1. git clone https://github.com/IKostarev/go-test-softweather.git
2. убедиться что у вас установлен компилятор Golang, в терминале - go version
3. go run .\cmd\substring\main.go
4. запуститься проект на порту :8080, при необходимости можно запустить его на другом порту, для этого нужно прописать go run .\cmd\substring\main.go --addr=":{number}"
5. открыть в браузере localhost:8080/
6. в терминале будет дополнительная информация по обработке запросов
```

### Через Docker:
```
1. git clone https://github.com/IKostarev/go-test-softweather.git
2. в терминале ввести - docker build -t app .
3. затем также в терминале - docker run -p 8080:8080 app
4. при необходимости изменить порт
```

### Через Make:
```
make build - Эта команда собирает исполняемый файл
```
```
make run - Эта команда запускает приложение локально с помощью команды go run
```
```
make docker-build - Эта команда собирает Docker-образ проекта
```
```
make docker-run - Эта команда запускает Docker-контейнер из ранее собранного Docker-образа
```
```
make clean - Эта команда удаляет собранный исполняемый файл, чтобы очистить проект от временных файлов или предыдущих сборок.
```

### Запуск тестов:
```
В корневой папке проекта открыть терминал и ввести go test ./...
```