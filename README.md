# Engineering 12 Final Project - Teenager

## Collaborators
  1. Irfan Kurniawan - Product Manager - FE2110451 ([Irfan858](https://github.com/Irfan858))
  2. Mohd Ryan Obillah - Frontend Developer - FE2163553 ([obillahh](https://github.com/obillahh))
  3. Umbu Theofilus Dendimara - Backend Developer - BE2144668 ([Rendydinar](https://github.com/Rendydinar))
  4. Widdy Arfiansyah - Backend Developer - BE2108902 ([arvians-id](https://github.com/arvians-id))
  5. Muhammad Abid Fajar - Backend Developer - BE2219863 ([abid313](https://github.com/abid313))
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
  6. Rudiansyah Wijaya Pratama - Backend Developer - BE2242991 ([Reezyx](https://github.com/Reezyx)) 
=======
>>>>>>> Stashed changes
  6. Rudiansyah Wijaya Pratama - Backend Developer - BE2242991 ([Reezyx](https://github.com/Reezyx)) 

# API Specification

All API must use this authentication
Request:
- Header:
    - X-Api-Key: ``` "your secret api key" ```

SUMMARY:
- [Users](#users)
- [User_course](#user-course)
- [Courses](#courses)
- [Modules](#modules)
- [Module_submissions](#module-submissions)
- [Module_articles](#module-articles)
- [Answers](#answers)
- [Questions](#questions)

## users
------------------------------
## Register Users [x]
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users ```
- Header:
    - Content-Type: ``` application/json ```
    - Accept: ``` application/json ```
- Body:
``` json
    {
        "name" : "string",
        "username" : "string", // unique
        "email" : "string", // unique
        "password" : "string",
        "role" : "integer", // enum(1,2)
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(0,1,2)
        "birthdate" : "date"

    }
```
Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "password" : "string",
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(0,1,2)
            "birthdate" : "date",
            "email_verified_at" : "timestamp", // timestamp
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Login Users [x]
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users/login ```
- Header:
    - Content-Type: ``` application/json ```
    - Accept: ``` application/json ```
- Body:
``` json
    {
        "email" : "string",
        "password" : "string",
    }
```
Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "token" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "email" : "string", // unique
            "role" : "integer", // enum(1,2)
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(0,1,2)
        }
    }
```
------------------------------
## Get Users Status [x]
------------------------------
Request:
- Method: ```GET ```
- Endpoint: ```/api/userstatus ```
- Header:
    - Content-Type: ``` application/json ```
    - Accept: ``` application/json ```
    - Authorization: ```{token} ```

Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique,
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(0,1,2)
            "address": "string",
            "birthdate" : "date",
            "image" : "string",
            "description": "string",
        }
    }
```
------------------------------
## Logout Users [x]
------------------------------
Request:
- Method: ```POST ```
- Endpoint: ```/api/users/logout ```
- Header:
    - Content-Type: ``` application/json ```
    - Accept: ``` application/json ```
    - Authorization: ```{token} ```

Response:
``` json
    {   
        "code" : "number",
        "status" : "string",
    }
```
------------------------------
## Get User By ID [x]
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users/{id}```
- Header:
    - Accept: ```application/json```  
    - Authorization: ```{token}```
- Variable: ```id```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(0,1,2)
            "address" : "string",
            "birthdate" : "date",
            "image" : "string",
        }
    }
```
------------------------------
## Update Users [x]
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/users/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
    - Authorization: ```{token}```
- Variable: ```id```
- Body:
``` json
    {
        "name" : "string",
        "username" : "string", // unique
        "role" : "integer", // enum(1,2)
        "phone" : "string",
        "gender" : "integer", // enum(1,2)
        "type_of_disability" : "integer", // enum(0,1,2)
        "address" : "string",
        "birthdate" : "date",
        "image" : "string",
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability" : "integer", // enum(0,1,2)
            "address" : "string",
            "birthdate" : "date",
            "image" : "string",
            "description" : "string"
        }
    }
```
------------------------------
<<<<<<< HEAD
<<<<<<< Updated upstream
=======
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
## Update Users Role [x]
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/users/roleupdate/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Variable:
    - id: ```integer```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "username" : "string", // unique
            "role" : "integer", // enum(1,2)
            "phone" : "string",
            "gender" : "integer", // enum(1,2)
            "type_of_disability": "integer", // enum(0,1,2)
            "address": "string",
            "birthdate": "string",
            "image": "string",
            "description": "string"
        }
    }
```
------------------------------
<<<<<<< HEAD
>>>>>>> Stashed changes
=======
>>>>>>> 6ca9fa7d7d3ad5fb18980dbb0f7d514ea1b3a885
## List Users [x]
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/users```
- Header:
    - Accept: ```application/json```
    - Authorization: ```{token}```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "username" : "string", // unique
                "role" : "integer", // enum(1,2)
                "phone" : "string",
                "gender" : "integer", // enum(1,2)
                "type_of_disability" : "integer", // enum(0,1,2)
                "address" : "string",
                "birthdate" : "date",
                "image" : "string",
                "description" : "string"
            },
        ]
    }
```
------------------------------
## Delete Users [x]
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/users/{id}```
- Header:
    - Accept: ```application/json```
    - Authorization: ```{token}```
    
Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## User course
------------------------------
## Create User_course
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/user_course```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "user_id" : "integer", // foreign key1
        "course_id" : "integer" // foreign key2
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
```
------------------------------
## Get User_course
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
```
------------------------------
## Update User_course
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "user_id" : "integer", // foreign key1
        "course_id" : "integer" // foreign key2
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "user_id" : "integer", // foreign key1
            "course_id" : "integer" // foreign key2
        }
    }
```
------------------------------
## List User_course
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/user_course```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "user_id" : "integer", // foreign key1
                "course_id" : "integer" // foreign key2
            },
        ]
    }
```
------------------------------
## Delete User_course
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/user_course/{user_id}/{course_id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Courses
------------------------------
## Create Courses
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/courses```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "name" : "string",
        "code_course" : "string", // unique
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    } 
```  
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## Get Courses
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## Update Courses
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/courses/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "name" : "string",
        "code_course" : "string", // unique
        "class" : "string",
        "tools" : "string", // longtext
        "about" : "string", // longtext
        "description" : "string" // longtext
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "name" : "string",
            "code_course" : "string", // unique
            "class" : "string",
            "tools" : "string", // longtext
            "about" : "string", // longtext
            "description" : "string", // longtext
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp 
        }
    }
```
------------------------------
## List Courses
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/courses```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "name" : "string",
                "code_course" : "string", // unique
                "class" : "string",
                "tools" : "string", // longtext
                "about" : "string", // longtext
                "description" : "string", // longtext
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp 
            },
        ]
    }
```
------------------------------
## Delete Courses
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/courses/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Modules
------------------------------
## Create Modules
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/modules```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "course_id" : "integer", // foreign key1
        "name" : "string",
        "is_locked" : "boolean",
        "estimate" : "integer",
        "deadline" : "timestamp", // timestamp
        "grade" : "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
```
------------------------------
## Get Modules
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/modules/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
```
------------------------------
## Update Modules
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/modules/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "course_id" : "integer", // foreign key1
        "name" : "string",
        "is_locked" : "boolean",
        "estimate" : "integer",
        "deadline" : "timestamp", // timestamp
        "grade" : "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "course_id" : "integer", // foreign key1
            "name" : "string",
            "is_locked" : "boolean",
            "estimate" : "integer",
            "deadline" : "timestamp", // timestamp
            "grade" : "integer"
        }
    }
```
------------------------------
## List Modules
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/modules```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "course_id" : "integer", // foreign key1
                "name" : "string",
                "is_locked" : "boolean",
                "estimate" : "integer",
                "deadline" : "timestamp", // timestamp
                "grade" : "integer"
            },
        ]
    }
```
------------------------------
## Delete Modules
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/modules/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Module submissions
------------------------------
## Create Module_submissions
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/module_submissions```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "file" : "string",
        "type" : "string",
        "max_size" : "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
```
------------------------------
## Get Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/module_submissions/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
```
------------------------------
## Update Module_submissions
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/module_submissions/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "file" : "string",
        "type" : "string",
        "max_size" : "integer"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "file" : "string",
            "type" : "string",
            "max_size" : "integer"
        }
    }
```
------------------------------
## List Module_submissions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/module_submissions```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "file" : "string",
                "type" : "string",
                "max_size" : "integer"
            },
        ]
    }
```
------------------------------
## Delete Module_submissions
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/module_submissions/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Module articles
------------------------------
## Create Module_articles
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/module_articles```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "content" : "string" // longtext
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
```
------------------------------
## Get Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/module_articles/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
```
------------------------------
## Update Module_articles
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/module_articles/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "content" : "string" // longtext
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "content" : "string" // longtext
        }
    }
```
------------------------------
## List Module_articles
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/module_articles```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "content" : "string" // longtext
            },
        ]
    }
```
------------------------------
## Delete Module_articles
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/module_articles/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Answers
------------------------------
## Create Answers
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/answers```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Get Answers
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/answers/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Update Answers
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/answers/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "question_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "question_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## List Answers
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/answers```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "question_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
```
------------------------------
## Delete Answers
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/answers/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
```
## Questions
------------------------------
## Create Questions
------------------------------
Request:
- Method: ```POST```
- Endpoint: ```/api/questions```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Get Questions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/questions/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## Update Questions
------------------------------
Request:
- Method: ```PUT```
- Endpoint: ```/api/questions/{id}```
- Header:
    - Content-Type: ```application/json```
    - Accept: ```application/json```
- Body:
``` json
    {
        "module_id" : "integer", // foreign key1
        "user_id" : "integer", // foreign key2
        "title" : "string",
        "tags" : "string",
        "description" : "string"
    }
```
Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : {
            "id" : "integer", // primary key
            "module_id" : "integer", // foreign key1
            "user_id" : "integer", // foreign key2
            "title" : "string",
            "tags" : "string",
            "description" : "string",
            "created_at" : "timestamp", // timestamp
            "updated_at" : "timestamp" // timestamp
        }
    }
```
------------------------------
## List Questions
------------------------------
Request:
- Method: ```GET```
- Endpoint: ```/api/questions```
- Header:
    - Accept: ```application/json```
- Query Param:
    - size : ```number```
    - page : ```number```

Response:
``` json
    {
        "code" : "number",
        "status" : "string",
        "data" : [
            {
                "id" : "integer", // primary key
                "module_id" : "integer", // foreign key1
                "user_id" : "integer", // foreign key2
                "title" : "string",
                "tags" : "string",
                "description" : "string",
                "created_at" : "timestamp", // timestamp
                "updated_at" : "timestamp" // timestamp
            },
        ]
    }
```
------------------------------
## Delete Questions
------------------------------
Request:
- Method: ```DELETE```
- Endpoint: ```/api/questions/{id}```
- Header:
    - Accept: ```application/json```

Response:
``` json
    {
        "code" : "number",
        "status" : "string"
    }
<<<<<<< Updated upstream
```
=======
```
>>>>>>> Stashed changes
>>>>>>> Stashed changes
