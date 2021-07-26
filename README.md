# FoodLisa

這是一個以 foodpanda 為主題的 Golang side project.

## Skill
* Gin Web Framework
* MySQL
* GORM
* air (live reload for go)


## Usage

本地安裝需要 package
```
cd app
go mod download
```

建 image
```
docker build -t lisa/foodlisa:1.2 ../docker/foodlisa_1.2 
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
| foodlisa | 1.2 | 2021/07/25 | 更新 go 套件初始化方式, 改成依 go.mod 在 container 建起來時自動下載 |
| foodlisa | 1.1 | 2021/06/06 | 加入air (live reload) |
| foodlisa | 1.0 | 2021/06/06 | 基於gin:1.0並加上自動啟動web application |
| gin | 1.0 | 2021/06/06 | go modules + gin框架 |


## 聯絡資訊
Lisa Chen
aa860221g@gmail.com
