# REST-API-CRUD-GIN

Menggunakan gol ver 1.13.15 dan framework gin dan gorm

## API dapat menyimpan Task
API yang Anda buat harus dapat menyimpan task melalui route :

- Method : POST

- URL : /tasks

- Body Request: 
```
{
    "assignedTo": string,
    "task": string,
    "deadline" : dateTime
}
```

Response to be returned :

- Status Code : 200
- Response Body : 
``` json
{
    "data": {
        "id": 1,
        "assignedTo": "Golang",
        "task": "RestAPI",
        "deadline": "2021-11-10T07:00:00+07:00",
        "created_at": "2021-11-01T17:10:37+07:00",
        "updated_at": "2021-11-01T17:10:37+07:00"
    }
} 
```

## API dapat menampilkan seluruh Task

API yang Anda buat harus dapat menampilkan semua Task yang disimpan melalui route :

- Method : GET

- URL : /tasks

## API dapat menampilkan satu Task

API yang Anda buat harus dapat menampilkan Task yang disimpan melalui route :

- Method : GET

- URL : /tasks/:id

## API dapat mengubah data Task
API yang Anda buat harus dapat mengubah data task berdasarkan id melalui rute:

- Method : PATCH

- URL : /tasks/:id

- Body Request: 
```
{
    "assignedTo": string,
    "task": string,
    "deadline" : dateTime
}
```

## API dapat menghapus data Task

API yang Anda buat harus dapat mengubah data task berdasarkan id melalui route :

- Method : DELETE

- URL : /tasks/:id

