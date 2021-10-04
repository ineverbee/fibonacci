# fibonacci

## Микросервис для создания среза чисел Фибоначчи.

**Сервис:**

Реализован микросервис для создания среза чисел Фибоначчи с порядковыми номерами от `x` до `y`. В данном сервисе реализованы два протокола: HTTP REST и GRPC. 

## Запуск приложения:

```
docker-compose up
```
Приложение будет доступно на порте `8081` для `curl` запросов

## gRPC

Чтобы создать запрос к серверу запустите клиент в другом терминале
```
docker-compose exec fibonacci ./client X Y
```
`X`,`Y` - аргументы, порядковые номера начала и конца среза

## REST

#### /fibo
* `POST` : Создать срез с числами Фибоначчи

Запрос:
```
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"x":7,"y":9}' \
  http://localhost:8081/fibo
```
`x,y` - порядковые номера начала и конца среза, `result` - срез с числами Фибоначчи
```
 {
  "result": [
    13,
    21,
    34
  ]
}
```

