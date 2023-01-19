Golang Test Task
- You will need docker and docker-compose for this task
- Fork the repository https://github.com/catbyte-GmbH/golang-test-task and run
  docker-compose up -d to start the docker containers
- Goal is to create 3 applications within the project as described below

Application 1: API (please use gin)
POST Endpoint: /message
POST body: { sender: String, receiver: String, message: String }
-> pushes received information to a RabbitMQ Queue
Return OK Status if everything is there, otherwise Bad Request

Application 2: MessageProcessor
Subscribes to queue from RabbitMQ and processes the message
Processing of message means saving the message to Redis in a way that application 3
works.

Application 3: Reporting API
GET Endpoint: /message/list
Parameters: sender: String, receiver String
Returns an array of objects with sender, receiver and message content that were
exchanged between sender and receiver in chronological descending order
Please commit your results to your own repository.