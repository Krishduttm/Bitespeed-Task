# Bitespeed-Task


This is the repository for the bitespeed task -> https://bitespeed.notion.site/Bitespeed-Backend-Task-Identity-Reconciliation-53392ab01fe149fab989422300423199

Step 1 -> download the code
Step 2 -> Create a database called "bitespeed"
Step 3 -> run command "docker-compose up --build"
Exposed PORT -> 8000
APIs ->
1) localhost:8000/bitespeed/getAll -> Returns all the contacts and their details for reference purposes
2) localhost:8000/bitespeed/getAll -> The API mentioned in the task.
cURL ->
```
curl --location 'localhost:8000/bitespeed/identify' \
--header 'Content-Type: application/json' \
--data-raw '{
	"email": "mcfly@hillvalley.edu",
	"phone_number": "123456"
}'
```
