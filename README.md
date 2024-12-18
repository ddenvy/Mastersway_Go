# Mastersway_Go
Learning Golang with Mastersway

### **Основы языка Go**

- Изучите базовые концепции языка:
    - Переменные, типы данных, функции.
    - Управляющие конструкции (if, for, switch).
    - pointer, structs, slices, and maps
    - Пакеты и импорт.
    - defer, panic, recover.
    - Функции и замыкания
    - Методы и интерфейсы
    - Общие типы (Generics)
    - Конкурентность (Concurrency)

---

### **2. Работа с базами данных**

- **Стандартные пакеты**:
    - `database/sql`: Основной пакет для работы с SQL-базами данных.
- **Драйверы для баз данных**:
    - **PostgreSQL**: `pgx`, `sqlx`.
    - **MySQL**: `mysql`.
    - **SQLite**: `go-sqlite3`.
    - **MSSQL**: `go-mssqldb`.
- **NoSQL базы данных**:
    - **MongoDB**: `mongo-go-driver`.
    - **Redis**: `redis`, `redigo`.
- **ORM и библиотеки**:
    - **Gorm**: Популярная ORM для Go.
    - **Ent**: ORM от Facebook для работы с графовыми базами данных.

---

### **3. Веб-разработка**

- **Стандартные пакеты**:
    - `net/http`: Основной пакет для создания HTTP-серверов и клиентов.
    - `html/template`: Работа с HTML-шаблонами.
- **Фреймворки**:
    - **Gin**: Быстрый и минималистичный веб-фреймворк.
    - **Echo**: Ещё один популярный фреймворк.
    - **Fiber**: Фреймворк, вдохновлённый Express.js.
    - **Beego**: Полнофункциональный фреймворк.

---

### **4. Микросервисы и обмен сообщениями**

- **Микросервисные фреймворки**:
    - **GoMicro**: Фреймворк для создания микросервисов.
    - **GoKit**: Набор инструментов для построения микросервисов.
- **gRPC**:
    - `grpc-go`: Библиотека для работы с gRPC.
- **Очереди сообщений**:
    - **Kafka**: `kafka-go`.
    - **RabbitMQ**: `amqp`.
    - **NATS**: `nats.go`.
    - **Amazon SQS**: `aws-sdk-go/sqs`.
    - **Google Cloud Pub/Sub**: `google-cloud-go/pubsub`.

---

### **5. Тестирование**

- **Стандартные пакеты**:
    - `testing`: Основной пакет для написания тестов.
    - `go tool cover`: Инструмент для проверки покрытия кода тестами.
    - `go tool pprof`: Инструмент для профилирования.
- **Библиотеки для тестирования**:
    - **Testify**: Упрощает написание тестов.
    - **GoMock**: Инструмент для создания mock-объектов.
    - **GoDog**: BDD-фреймворк для Go.

---

### **6. Логирование и конфигурация**

- **Логирование**:
    - **log**: Стандартный пакет для логирования.
    - **Zap**: Быстрый и мощный логгер.
    - **ZeroLog**: Альтернатива Zap.
    - **Apex**: Ещё один логгер с поддержкой JSON.
- **Конфигурация**:
    - **flag**: Стандартный пакет для работы с аргументами командной строки.
    - **Viper**: Библиотека для работы с конфигурациями.
    - **Cobra**: Библиотека для создания CLI-приложений.

---

### **7. Дополнительные инструменты**

- **Go Prompt**: Инструмент для создания интерактивных CLI-приложений.
- **Squirrel**: Библиотека для построения SQL-запросов.
- **Jaeger**: Инструмент для трассировки (tracing).

---

### **8. Практика и проекты**

- Создайте проекты, чтобы закрепить знания:
    - **REST API**: Используйте Gin или Echo.
    - **Микросервис**: Напишите простой микросервис с использованием GoMicro или GoKit.
    - **CLI-приложение**: Используйте Cobra для создания утилиты.
    - **Работа с базой данных**: Реализуйте CRUD-операции с использованием Gorm или Ent.
