##RUN

RUN app

```bash
CONFIG_PATH=~/url-shortener/config/local.yaml  go run main.go
```
You can expose envirenment it differently 

Run test 

```bash
go test
```

## Добавление url от пользоватея
![Добавление url от пользоватея](./url-shortener/img/save.png)

## Перенаправление на сохраненный url 
![Полчение url](./url-shortener/img/get-url.png)
