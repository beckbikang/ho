# ho

features
- ✅ a simple mc server use [library](https://github.com/rpcxio/gomemcached)
- ✅ write data into  kafka which use  mc protocol. 
- trans data from kafka to kafka
- write data into  kafka which use  redis protocol. 



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

##  mc send to kafka

connect server
- telnet 127.0.0.1 9191

set data 
- set topic_name 0 1200 3

kafka config

```
title="mc-to-kafka"
[main]
serverIp="127.0.0.1"
serverPort=9191
mainLogPath="/tmp/ho_mc_2_kafka"
mainLogModel=3 

[kafkas]
    [kafkas.test_mc]
    brokers=["tt.com:18888"]
    topic="test_mc"
    group=""
    sslEnable=false
    user=""
    password=""
    producerOn=true
    [kafkas.test_mc2]
    brokers=["tt.com:18888"]
    topic="test_mc2"
    group=""
    sslEnable=false
    user=""
    password=""
    producerOn=true

```



## dependent on

a memmcache server: [go-memcache](https://github.com/rpcxio/gomemcached)

a common framework: [ngo](https://github.com/NetEase-Media/ngo)

a redis server:[redcon](https://github.com/tidwall/redcon)









