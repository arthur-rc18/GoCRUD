# CRUDZEUS

Users API to create, read, delete and update users.

## Endpoints

```sh
GET /users
GET /users/:id
POST /users
PUT /users/:id
DELETE /users/:id
PATCH /users/:id
```

## Dependencies

- Postgresql
- PgAdmin 4

## Cloning procedure

```sh
git clone 'something'
```

## How to run this application

In the main folder

```sh
go run .
```

## API Documentation

This API uses `REST` to communicate and HTTP [response codes](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status) to indenticate status and errors. All responses come in standard JSON. All requests must include a content-type of application/json and the body must be a valid JSON.

## Response Codes

``` sh
200: Success
404: Not found
500: Internal Server Error
```

## Domains

**Development:**

```sh
localhost:8800
```

## Get Users (present data of users in the postgres database)

present data of users in the postgres database

**Request:**

``` sh
GET http://'domain'/users
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
[
    {
        "id": 63,
        "nome": "arthur",
        "email": "test@gmail.com",
        "senha": "123456",
        "telefone": "3495754098"
    },
    {
        "id": 64,
        "nome": "alberto",
        "email": "testing@gmail.com",
        "senha": "123456",
        "telefone": "3495724098"
    }
]
```

## Get Users

Gets the data of a specific user by the ID.

**Request:**

``` sh
GET http://'domain'/users/'id'
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
[
    {
        "id": 63,
        "nome": "arthur",
        "email": "test@2gmail.com",
        "senha": "123456",
        "telefone": "3495754098"
    }
]
```

**Failed Response:**

Request with a non-existent ID

http code: 404 Not found

```JSON
{
    "error": "ID not found"
}
```

## Post Users

Create a new user with all values null(still have to fix this).

**Request:**

``` sh
Post http://'domain'/users/
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
{"user":{"id":0,"nome":"something","email":"something","senha":"*****","telefone":"******"}}
```

## Put Users

Update an already existed user by the id.

**Request:**

``` sh
Put http://'domain'/users/'id'
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
{"message":"user updated with success","user":{"id":0,"nome":"","email":"","senha":"","telefone":""}}
```

## Delete Users

Delete an user by the ID.

**Request:**

``` sh
Delete http://'domain'/users/'id'
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
{"message":"User deleted","userId":0}
```

**Failed Response:**

http code: 404 Not found

```JSON
```

## Patch Users

Update partially the user's field.

**Request:**

``` sh
Patch http://'domain'/users/'id'
```

**Successful Response:**

http code: 200 OK

Example:  

```JSON
{"message":"user updated with success","user":{"id":0,"nome":"","email":"","senha":"","telefone":""}}
```
