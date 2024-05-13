# specs
## apis
### auth api
#### admin auth
##### curl
`
curl -X POST -H "Content-Type: application/json" -d '{"user_id":"admin", "password":"admin"}' http://localhost:9000/admin/authentication
`

### user api
#### GET
##### url
`
/user/${userId}
`
##### request
##### response
###### body
| column        | type   | validation | note |
| :------------ | ------ | ---------- | ---- |
| user          | object | not null   |      |
| - user_id     | string | not null   |      |
| - session_key | string | not null   |      |
| - theme_id    | string | not null   |      |

##### curl 

`
curl -X GET -H "Content-Type: application/json" http://localhost:9000/user/user
`


### chapter api
#### GET
##### url
`
/chapter/${chapterId}
`
##### request
###### header

| key           | type   | validation | note              |
| :------------ | ------ | ---------- | ----------------- |
| x-session-key | string | not null   | get with user api |

##### response
###### body

| column               | type   | validation | note |
| :------------------- | ------ | ---------- | ---- |
| chapter              | object | not null   |      |
| - chapter_id         | string | not null   |      |
| - main_code          | string | not null   |      |
| - example_code       | string | not null   |      |
| - expected           | string | not null   |      |
| - best_practice_code | string |            |      |
| - level              | int    | not null   |      |
| - exercise           | string |            |      |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-session-key: xxxxxxxxxx-01" http://localhost:9000/chapter/CHAPTER-AA-01
`

#### POST
##### url
`
/chapter
`
##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

###### body

| column               | type   | validation | note |
| :------------------- | ------ | ---------- | ---- |
| chapter              | object | not null   |      |
| - chapter_id         | string | not null   |      |
| - main_code          | string | not null   |      |
| - example_code       | string | not null   |      |
| - expected           | string | not null   |      |
| - best_practice_code | string |            |      |
| - level              | int    | not null   |      |
| - exercise           | string |            |      |

##### response
######
empty
##### curl
`
curl -X POST -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" -d '{"chapter":{"chapter_id":"CHAPTER0001", "main_code":"sample01", "example_code":"sample02", "expected":"sample03", "best_practice_code":"sample04", "level":1, "exercise":"sample05"}}' http://localhost:9000/chapter
`

#### ALL GET
##### url
`
/chapter/all
`
##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

##### response
###### body

| column               | type   | validation | note |
| :------------------- | ------ | ---------- | ---- |
| chapters             | array  | not null   |      |
| - chapter_id         | string | not null   |      |
| - main_code          | string | not null   |      |
| - example_code       | string | not null   |      |
| - expected           | string | not null   |      |
| - best_practice_code | string |            |      |
| - level              | int    | not null   |      |
| - exercise           | string |            |      |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" http://localhost:9000/chapter/all
`

#### List GET(theme id)
##### url
`
/chapter/list/${themeId}
`
##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

##### response
###### body

| column               | type   | validation | note |
| :------------------- | ------ | ---------- | ---- |
| chapters             | array  | not null   |      |
| - chapter_id         | string | not null   |      |
| - main_code          | string | not null   |      |
| - example_code       | string | not null   |      |
| - expected           | string | not null   |      |
| - best_practice_code | string |            |      |
| - level              | int    | not null   |      |
| - exercise           | string |            |      |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" http://localhost:9000/chapter/list/THEME-01
`


### theme api
#### ALL GET
##### url
`
/theme/all
`

##### request
###### header

| key           | type   | validation | note              |
| :------------ | ------ | ---------- | ----------------- |
| x-session-key | string | not null   | get with user api |

##### response
###### body

| column           | type    | validation | note                              |
| :--------------- | ------- | ---------- | --------------------------------- |
| theme            | string  | not null   |                                   |
| - theme_id       | string  | not null   |                                   |
| - theme          | string  | not null   |                                   |
| - description    | string  | not null   |                                   |
| archivements     | arrays  | not null   |                                   |
| - archivement_id | string  | not null   |                                   |
| - chapter_id     | int     | not null   |                                   |
| - order          | integer | not null   | display order                     |
| - status         | string  | not null   | 0:not start,1:pending,2,completed |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-session-key: xxxxxxxxxx-01" "http://localhost:9000/theme"
`


#### GET
##### url
`
/theme
`

##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

##### response
###### body

| column        | type   | validation | note |
| :------------ | ------ | ---------- | ---- |
| themes        | array  | not null   |      |
| - theme_id    | string | not null   |      |
| - theme       | string | not null   |      |
| - description | string | not null   |      |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" "http://localhost:9000/theme/all"
`

#### POST
##### url
`
/theme
`
##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

###### body

| column        | type   | validation | note |
| :------------ | ------ | ---------- | ---- |
| theme         | object | not null   |      |
| - theme_id    | string | not null   |      |
| - theme       | string | not null   |      |
| - description | string | not null   |

##### response
######
empty
##### curl
`
curl -X POST -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" -d '{"theme":{"theme_id":"THEME-XX", "theme":"sample1", "description":"sample2"}}' http://localhost:9000/theme
`

#### CHAPTER PUT
##### url
`
/theme/chapter
`
##### request
###### header

| key          | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| x-master-key | string | not null   |      |

###### body

| column       | type   | validation | note |
| :----------- | ------ | ---------- | ---- |
| theme_id     | string | not null   |      |
| chapters     | array  | not null   |      |
| - chapter_id | string | not null   |      |
| - order      | int    | not null   |      |

##### response
######
empty
##### curl
`
curl -X PUT -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" -d '{"theme_id":"THEME-XX", "chapters":[{"chapter_id":"CHAPTER-AA-03", "order":1}, {"chapter_id":"CHAPTER-AA-04", "order":2}]}' http://localhost:9000/theme/chapter
`


### archivement api
#### GET
##### url
`
/archivement/${chapterId}
`

##### request
###### header

| key           | type   | validation | note              |
| :------------ | ------ | ---------- | ----------------- |
| x-session-key | string | not null   | get with user api |

##### response
###### body

| column             | type    | validation | note |
| :----------------- | ------- | ---------- | ---- |
| archivement        | string  | not null   |      |
| - archivement_id   | string  | not null   |      |
| - status           | string  | not null   |      |
| - version          | int     | not null   |      |
| - code             | arrays  | not null   |      |
| - comment          | string  | not null   |      |
| - result           | string  | not null   |      |
| - is_compile_error | boolean | not null   |      |

##### curl
`
curl -X GET -H "Content-Type: application/json" -H "x-session-key: xxxxxxxxxx-01" http://localhost:9000/archivement/CHAPTER-AA-01
`

#### POST
##### url
`
/archivement
`

##### request
###### header

| key           | type   | validation | note              |
| :------------ | ------ | ---------- | ----------------- |
| x-session-key | string | not null   | get with user api |

###### body

| column             | type    | validation | note |
| :----------------- | ------- | ---------- | ---- |
| archivement        | string  | not null   |      |
| - archivement_id   | string  | not null   |      |
| - status           | string  | not null   |      |
| - version          | int     | not null   |      |
| - code             | arrays  | not null   |      |
| - comment          | string  | not null   |      |
| - result           | string  | not null   |      |
| - is_compile_error | boolean | not null   |      |

##### response

##### curl
`
curl -X POST -H "Content-Type: application/json" -H "x-session-key: xxxxxxxxxx-01" -d '{"archivement":{"archivement_id":"931d772e-e54f-428c-873a-17be8608dac0","status":"3","version":1,"code":"aaaa","comment":"bbbbb","result":"ccccc","is_compile_error":true}}' http://localhost:9000/archivement
`
