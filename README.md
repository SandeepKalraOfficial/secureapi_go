Below are the Curl requests for handling users for the project. This is a sample data

curl -X POST http://localhost:8000/people -H "Content-Type: application/json" -d "{\"id\":\"1\", \"fname\":\"Sandeep\", \"lname\":\"Kalra\", \"dob\":\"1995-07-20T00:00:00Z\"}"
curl http://localhost:8000/people
curl http://localhost:8000/people/1
curl -X PUT http://localhost:8000/people/2 -H "Content-Type: application/json" -d "{\"fname\":\"Neha\", \"lname\":\"Sharma\", \"dob\":\"1992-12-11T00:00:00Z\"}"
curl -X DELETE http://localhost:8000/people/1

