# project book list app

book list app is an application to create and manage task, This app has : 
* User Friendly
* Easy to Use
* RESTful endpoint for product CRUD operation
* JSON formatted response

&nbsp;

## List of available endpoints
### users
- `GET /users`
- `POST /users/register`
- `POST /users/login`
- `GET /users/:user_id`
- `PUT /users/:user_id`
- `DELETE /users/:user_id`

### todos
- `GET /books`
- `POST /books`
- `GET /books/:book_id`
- `PUT /books/:book_id`
- `DELETE /books/:book_id`


## RESTful endpoints users
### GET /users

> Get All users

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
[
  {
      "id" : 1,
      "name" : "afista pratama",
      "email" : "pratama@mail.com",
      "date_birth": "1997-08-12",
      "address" : "jember", 
  }, {
      "id" : 2,
      "name" : "eddy permana",
      "email" : "eddy_p@mail.com",
      "date_birth" : "YYYY-MM-DD",
      "address" : "palembang",
  }
]
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : ""
}
```
---

### POST /users/register

> Create new user 

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "name" : "<name to get insert into>",
  "address" : "<address to get insert into>",
  "date_birth" : "<date_birth to get insert into>",
  "email": "<email to get insert into>",
  "password": "<password to get insert into>"
}
```

_Response (201)_
```json
{
  "id" : "<given id by system>",
  "name" : "<posted name>",
  "address" : "<posted address>",
  "date_birth" : "<posted date_birth>",
  "email" : "<posted email>"    
}
```

_Response (400 - Bad Request)_
```json
{
  "errors" : []
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : ""
}
```
---

### POST /users/login

> Compare data login on database with data inputed

_Request Header_
```
not needed
```

_Request Body_
```json
{
  "email": "<email to get compare>",
  "password": "<password to get compare>"
}
```

_Response (200)_
```json
{
  "id" : "<posted id>",
  "name": "<posted name>",
  "address" : "<posted address>",
  "date_birth" : "<posted date_birth>",
  "email": "<posted email>",
  "authorization" : "<your access token/ authorization>"
}
```

_Response (400 - Bad Request)_
```json
{
  "errors" : ["<posted error>","<posted error>"]
}
```

_Response (401 - Unauthorized)_
```json
{
  "error" : "<posted error>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### GET /users/:user_id

> Get user by user ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "id": "<posted id>",
  "name": "<posted name>",
  "address": "<posted address>",
  "date_birth" : "<posted date_birth>",
  "email": "<posted email>",
}
```

_Response (400 - Bad Request)_
```json
{
  "error" : "<posted error>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : ""
}
```
---

### PUT /users/:user_id

> Update user by User iD

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "name" : "<name to update>",
    "address" : "<address to update>",
    "date_birth" : "<date_birth to update format:YYYY-MM-DD>",
}
```

_Response (200)_
```json
{
  "id" : "<posted id>",
  "name" : "<posted name update>",
  "address" : "<posted address update>",
  "date_birth" : "<posted date_birth update>",
  "email" : "<posted email>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### DELETE /users/:user_id

> Delete user by ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{ 
  "data" : "user id <posted id by params> success deleted",
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : ""
}
```
---

## RESTful endpoints books
### GET /books

> Get All books by user login

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
[
  {
    "id": 1,
    "title" : "harry potter",
    "author" : "J.K Rowling",
    "year" : 1999
  },{
    "id" : 2,
    "title" : "Harry Potter and the Cursed Child",
    "author" : "J.K Rowling",
    "year" : 2001
  }
]
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### POST /books

> Create new book 

_Request Header_
```json
{
    "Authorization" : "<your Authorization>"
}
```

_Request Body_
```json
{
  "title" : "<title to get insert into>",
  "author" : "<author to get insert into>",
  "year": "<year to get insert into >",
}
```

_Response (201)_
```json
{
  "id": <given id by system>,
  "title": "<posted title>",
  "author": "<posted author>",
  "year": "<posted year>",
  "user_id": "<posted user_id>",
  "created_at": "<posted created_at>",
  "updated_at": "<posted updated_at>"
}
```

_Response (400 - Bad Request)_
```json
{
  "errors" : ["<posted error>", "<posted error>"]
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### GET /books/:book_id

> Get books detail by book_id

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "id": "<posted id>",
  "title": "<posted title>",
  "author": "<posted author>",
  "year": "<posted year>",
  "user_id": "<posted user_id>",
  "created_at": "<posted created_at>",
  "updated_at": "<posted updated_at>"
}
```

_Response (401 - Unauthorized)_
```json
{
  "error" : "<posted error>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### PUT /books/:book_id

> Update book by book ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```json
{
    "title" : "harry potter go to indonesia",
    "author" : "JK Rowling",
    "year" : 2021,
}
```

_Response (200)_
```json
{
  "id": 1,
  "title": "harry potter go to indonesia",
  "author": "JK Rowling",
  "year": 2021,
  "user_id": 1,
  "created_at": "2021-05-26T15:24:05.36100953Z",
  "updated_at": "2021-05-26T15:24:05.361009826Z"
}
```

_Response (400 - Bad Request)_
```json
{
  "errors" : ["<posted error>", "<posted error>"]
}
```


_Response (401 - Unauthorized)_
```json
{
  "error" : "<posted error>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---

### DELETE /books/:book_id

> Delete book by book ID

_Request Header_
```json
{
   "Authorization": "<your Authorization>"
}
```

_Request Body_
```
not needed
```

_Response (200)_
```json
{
  "data" : "book id <posted id> success deleted",
}
```

_Response (400 - Bad Request)_
```json
{
  "error" :  "<posted error>"
}
```

_Response (500 - Internal Server Error)_
```json
{
  "error" : "<posted error>"
}
```
---