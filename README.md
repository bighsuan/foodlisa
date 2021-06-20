# FoodLisa

這是一個以 foodpanda 為主題的 Golang side project.

## Skill
* gin web framework
* air (live reload for go)
* my-sql

## Usage

安裝需要package
```
cd app
go mod download
```

建 image
```
docker build -t lisa/foodlisa:1.1 ../docker/foodlisa_1.1 
```

建container
```
docker compose -f ../docker/docker-compose.yml up -d
```



打開localhost:8080即可看到
```
{"message":"Welcome to FoodLisa"}
```


## Docker Image

| Image Name | Version | Update Date | Description 
| :----:| :---: | :----: | :---- |
| foodlisa | 1.0 | 2021/06/06 | 基於gin:1.0並加上自動啟動web application |
| foodlisa | 1.1 | 2021/06/06 | 加入air (live reload) |
| gin | 1.0 | 2021/06/06 | go modules + gin框架 |

## 聯絡資訊
Lisa Chen
aa860221g@gmail.com