# specs
## apis
### auth api
##### admin auth
`
curl -X POST -H "Content-Type: application/json" -d '{"user_id":"admin", "password":"admin"}' http://localhost:9000/admin/authentication
`

#### chapter api
##### GET
###### url
`
/chapter/${chapterId}
`
`
curl -X GET -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" http://localhost:9000/chapter/CHAPTER0001
`

##### POST
###### url
`
/chapter
`
###### header

| key                 | type   | validation | note | 
| :------------------ | ------ | ---------- | ---- | 
| x-master-key        | string | not null   |      | 

###### body

| column              | type   | validation | note | 
| :------------------ | ------ | ---------- | ---- | 
| chapter             | string | not null   |      | 
| 　chapter_id        | string | not null   |      | 
| 　main_execute_code | string | not null   |      | 
| 　init_code         | string | not null   |      | 
| 　expected          | string | not null   |      | 
| 　answer_code       | string |            |      | 
| 　level             | int    | not null   |      | 

##### curl
`
curl -X POST -H "Content-Type: application/json" -H "x-master-key: 5d71bca7-17f6-4646-82df-7dc9397e9422" -d '{"chapter":{"chapter_id":"CHAPTER0001", "main_execute_code":"sample01", "init_code":"sample02", "expected":"sample03", "answer_code":"sample04", "level":1}}' http://localhost:9000/chapter
`