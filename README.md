
# Тестовое задание Куспанов Азамат

## Приложение принимает 2 фалага: -name и -time.  В качестве БД используется Postgresql.

-name: string задает имя файла.

-time: number временной интервал в секундах между запросами в БД.

Для работы приложения необходим Postgresql сервер со схемой указанной в файле scheme/up.sql. DB_URL можно задать в cmd/main.go

Использование
-------------

Склонируйте репозиторий:

    git clone https://github.com/Kin-dza-dzaa/DNSVladivostok

Смените дерикторию:

    cd DNSVladivostok

Установите зависимости:

    go mod download

Забилдите приложение: 

    go build -o . ./cmd/.

Запустите его указав хотябы 1 флаг -name:

    cmd -name test.txt -time 1
