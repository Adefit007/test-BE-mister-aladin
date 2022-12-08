# Test BE using Golang

## cara menjalankan program

- download aplikasi go language di laptop atau di dekstop

- silahkan clone repository github

```bash
git clone https://github.com/Adefit007/test-BE-mister-aladin.git
```

- jalankan aplikasi XAMPP dan kemudian aktifkan Apache dan Mysql

- buat di dalam phpMyadmin database baru dengan nama article

- jalankan code program

```bash
go run main.go
```

- jalankan aplikasi postman untuk test API

> Post a new articles : `http://localhost:5000/api/v1/articles` with method `POST`
> Get the list of articles : `http://localhost:5000/api/v1/articles` with method `GET`
> Get a specific articles by id : `http://localhost:5000/api/v1/articles/1` with method `GET`
> Update an articles by id : `http://localhost:5000/api/v1/articles/1` with method `PUT`
> Delete an articles by id : `http://localhost:5000/api/v1/articles/1` with method `DELETE`
