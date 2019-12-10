# Сalendar

## Docker-compose preparation
***make build***
***docker-compose up*** 

1. “Заготовка” для микросервиса “Календарь”
Цель: В результате выполнения ДЗ должен получиться базовый скелет микросервиса, который будет развиваться в дальнейших ДЗ. Структура кода должна соответствовать подходу Clean Architecture. В данном задании тренируются навыки:
 - декомпозиции предметной области;
 - построения элементарной архитектуры проекта.
Завести в репозитории отдельную директорию для проекта "Календарь"
Создать внутри структуру директорий, соответствующую Clean Architecture.

Cоздать модели (структуры) календаря
Cоздать методы бизнес логики (методы у структур) для работы с этими структурами:
- добавление событий в хранилище
- удаление событий из хранилища
- изменение событий в хранилище
- листинг событий
- пр. на усмотрение студента
Создать объекты ошибок (error sentinels) соответсвующие бизнес ошибкам, например ErrDateBusy - данное время уже занято другим событием
Реализовать хранение событий в памяти (т.е. просто складывать объекты в слайсы)
Реализовать Unit тесты проверяющие работу бизнес логики (в частности ошибки)

На данном этапе не нужно:
- Делать HTTP, GRPC и пр. интерфейсы к микросервису
- Писать .proto-файлы (это будет позже)
- Использовать СУБД
Критерии оценки: Критерии оценки:
- Созданы структуры, необходимые для реализации календаря
- Соблюдены принципы Clean Architecture
- Работа с хранилищем через интерфейс 
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и проверить пакет с помощью go get / go test

2. Каркас микросервиса
Цель: Реализовать "каркас" микросервиса, считывающий конфиг из файла, создающий логгер/логгеры с указанными уровнями детализации.
Необходимо доработать код сервиса "Календарь" из предыдущего задания, добавив в него:

* Обработку аргументов командной строки
* Чтение файла конфигурации (параметр --config в командной строке)
* Создание логгеров и настройка уровня логирования
* Создание и запуск hello-world web-сервера

Параметры, передаваемые через аргументы командной строки:
* --config - путь к конфигу

Параметры, которые должны быть в конфиге:
* http_listen - ip и port на котором должен слушать web-сервер
* log_file - путь к файлу логов
* log_level - уровень логирования (error / warn / info / debug)

Критерии оценки: Web-сервер на данном этапе может быть не связан с бизнес логикой календаря и должен обрабатывать только URL /hello
Web-сервер должен запускаться на ip:port указанном в конфиге и каждый обработанный запрос должен выводиться в log-файл.
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install
В репозитории должен быть образец конфига
Установленный сервис должен запускаться

3. HTTP интерфейс
Цель: Реализовать HTTP интерфейс для сервиса Календаря.
Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md
Цель данного задания - отработать навыки работы со стандартной HTTP библиотекой,
поэтому технологии JSONRPC, Swagger и т.п. НЕ используются.

В директории с проектом создать отдельный пакет для Web-сервера
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области
Реализовать middleware для логирования запросов

Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09)
В GET методах параметры передаются через queryString, в POST через тело запроса.

В результате каждого запроса должен возвращаться JSON документ содержащий 
либо {"result": "..."} в случае успешного выполнения метода
либо {"error": "..."} в случае ошибки бизнес-логики
Критерии оценки: Все методы должны быть реализованы
Бизнес логика (пакет internal/domain в примере) НЕ должен зависеть от кода HTTP сервера 
В случае ошибки бизнес-логики сервер должен возвращать HTTP 200
В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400
В случае остальных ошибок сервер должен возвращать HTTP 500
Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install

4. GRPC сервис
Цель: Создать GRPC API для сервиса календаря 
Тех. задание: https://github.com/OtusTeam/Go/blob/master/project-calendar.md
Цель данного занятия: отработка навыков работы с GRPC, построение современного API.

Создать отдельную директорию для Protobuf спек.
Создать Protobuf спеки с описанием всех методов API, их объектов запросов и ответов.
Т.к. объект Event будет использоваться во многих ответах разумно выделить его в отдельный message.
Создать отдельный директорию для кода GRPC сервера
Сгенерировать код GRPC сервера на основе Protobuf спек (скрипт генерации сохранить в репозиторий).
Написать код, связывающий GRPC сервер с методами доменной области.

Критерии оценки: Все методы должны быть реализованы
Бизнес логика (пакет internal/domain в примере) НЕ должен зависеть от кода GRPC сервера 
GRPC-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
   
Код должен проходить проверки go vet и golint
У преподавателя должна быть возможность заново сгенерировать код по Protobuf спекам
У преподавателя должна быть возможность скачать и установить пакет с помощью go get / go install

5. Работа с базами данных
Изменить код сервиса-календаря, так что бы события хранились в базе данных.

6. Работа с очередями
Реализовать "напоминания" о событиях с помощью RabbitMQ.
Создать процесс, который периодически сканирует основную базу данных,
выбирая события о которых нужно напомнить (ЕК: и кладет в очередь).
Создать процесс, который читает сообщения из очереди и шлет уведомления.

7. Доработка сервиса
Цель: Данное ДЗ посвящено доработки кода. Новые навыки не отрабатываются.
Важно довести проект до рабочего состояния, разделив его на отдельные сервисы.
В результате компиляции проекта должно получаться 3 отдельных исполнимых (по одному на микросерви).
Каждый из сервисов должен принимать путь файлу конфигурации:
./calendar_api --config=/path/to/config.yaml
./calendar_scheduler --config=/path/to/config.yaml
./calendar_sender --config=/path/to/config.yaml
Настройки СУБД и очереди собщений должны браться из файла конфигурации.
Критерии оценки: Преподаватель может собрать все файлы одной одной командой
Например make build или go install ./...

8. Докеризация сервиса
Цель: Цель домашнего задания: запустить все компоненты проекта в Docker
Отрабатываются навыки работы с Docker и docker-compose
Создать Dockerfile для каждого из микросервисов (api, scheduler, sender)
Собрать образы и проверить их локальный запуск
Создать docker-compose файл, который запускает PostgreSQL, RabbitMQ и все микросервисы вместе.
Для PostgreSQL и RabbitMQ использовать официальные образы из dockerhub. 
Так же в docker-compose должен запускаться one-shot скрипт который применяет SQL миграции, создавая структуру СУБД.
Для контейнера с API необходимо пробросить (expose) порт 8888 на хост-машину.
Критерии оценки: У преподавателя должна быть возможность запустить весь проект с помощью команды docker-compose up
После этого API должно быть доступно по URL http://localhost:8888/

9. Интеграционное тестирование
Цель: Цель данного домашнего задания: научиться писать интеграционные тесты к web-сервисам
В данном ДЗ изучается BDD, язык Gherkin, отрабатываются навыки работы с BDD библиотекой github.com/DATA-DOG/godog
И еще раз с docker-compose =)
Создать отдельный пакет для интеграционных тестов
Описать все бизнес-сценарии на языке Gherkin в *.feature файлах.
Реализовать все шаги сценариев с использованием библиотеки Godog
При этом шаги могут рассчитывать на то что запущены в docker-compose и знают hostname:port все сервисов.
Создать docker-compose файл, поднимающий все сервисы проекта + контейнер с интеграционными тестами
В Makefile добавить команду test, которая будет запускать интеграционные тесты (см -https://docs.docker.com/compose/reference/up/ -exit-code-from)
Критерии оценки: Тесты покрывают все основные бизнес сценарии:
 - добавление события и обработку бизнес ошибок
 - получение листинга событий на день / неделю / месяц
 - отправку уведомлений
Преподаватель может запустить интеграционные тесты в docker-compose с помощью команды make test
В случае успешного выполнения команда должна возвращать 0, иначе 1

10. Мониторинг сервиса
Обеспечить простейший мониторинг проекта с помощью prometheus
Prometheus запустить в docker контейнере рядом с остальными сервисами.
Для API сервиса необходимо измерять:
 * Requests per second
 * Latency
 * Коды ошибок
 * Все это в разделении по методам (использовать отдельный тэг prometheus для каждого метода API)
Для баз данных:
 * Количество записей в таблице events (данные брать из pg_stat_user_tables)
 * Стандартные метрики базы: Transactions per second, количество подключений (использовать готовый exporter)
Для расслыльщика:
 * RPS (кол-во отправленных сообщений в сек)