# ho

features
- ✅ a simple mc server use [library](https://github.com/rpcxio/gomemcached)
- write data into  kafka which use  mc protocol. 



## a simple mc server 

run server like this
- ./ho mc-server -p /you config file path/golang/ho/configs -f config-mc-server.toml

connect server
- telnet 127.0.0.1 9190

set data 
- set tom 0 1200 3

get data 
- get tom

<img width="363" alt="image" src="https://user-images.githubusercontent.com/7270440/189366315-3353ae47-4799-49d2-882a-f6fe14df5b98.png">

