
# BACKEND CAKESTORE

Privy Test Cake Store RESTFul API


## Tech Stack

**Golang:** programming language


**Gorilla/mux:** handle http request

**Mysql:** DBMS

**Docker:** Deploy Applications




## Installation Steps

1. Clone Repository
```bash
  git clone https://github.com/HendraLatumeten/CakeStore.git
```
2. Install Dependencies
```bash
  go get -u ./...
```
3. Create Database mysql
```bash
 CREATE DATABASE db_name;
```
4. Run Migration
```bash
 migrate -path src/database/migration -database "mysql://root:@tcp(localhost:3306)/db_name" up
```
5. Run App
```bash
 go run .
```





## API Reference

#### 1. Get All Data Cake

```
  GET /menucake/list
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-` | `string` | list data cake|

#### 2. Get Detail Data Cake

```
  GET /menucake/detail/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### 3. Get All Data Cake

```
  GET /menucake/save
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `-` | `string` | save data cake |

#### 4. Delete Data Cake

```
  GET /menucake/delete/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id as params remove

 
 #### 5. Update Data Cake

```
  GET /menucake/Update/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id as params update

 
