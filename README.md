# ctdf-acc-gen - библиотека для генерации данных аккаунтов в ctfd api
Требования для запуска: Docker/Compose  
Использовано: Go 1.25, Docker/Compose, PostgreSQL 17  
Зависимости:  
-github.com/spf13/cobra  
-github.com/joho/godotenv   
-gorm.io/gorm  

# Запуск
## Необходимые переменные окружения:
  CTFD_ADMIN_TOKEN - для авторизации в ctfd API  
  CTFD_URL - ссылка на API  
  DB_NAME  
  DB_USER  
  DB_PASSWORD  
  DB_HOST  
  DB_PORT  
Пример параметров в .env.example  
## Команды
docker compose up db  
docker compose build lib  
docker compose run --rm lib ctfd create_users n  -  создание n пользователей  
docker compose run --rm lib ctfd create_teams teamCount teamSize  -  создание teamCount команд c teamSize пользователями  
docker compose run --rm lib ctfd check_user username1 username2 ...  -  проверка на существование в PostgreDB и CTFd пользователя  
docker compose run --rm lib ctfd check_team team_name1 team_name2 ...  -  проверка на существование в PostgreDB и CTFd команды  

**Если не удается создать объект, то выполнение прекрашается**
## Пример использования
```go
  _ = godotenv.Load()
  # host=db для взаимодействия между контейнерами, либо host=localhost, если обращение на локально запущенную бд
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s sslmode=disable port=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	client := ctfd.NewCTFdClient(os.Getenv("CTFD_URL"), os.Getenv("CTFD_ADMIN_TOKEN"))

	cfg := ctfdgen.Config{CTFDClient: client, DB: db}

	generator, err = ctfdgen.NewGenerator(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	ctx = context.Background()
```
