# ho

features
- ✅ a simple mc server use this [library](https://github.com/rpcxio/gomemcached)
- ✅ write data into  kafka which use  mc protocol. 
- ✅ write data into  kafka which use  redis protocol. 
- ✅ trans data from kafka to kafka
- trans kafka to elasticsearch


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

run server
./ho mc kafka -p /xxx/golang/ho/configs -f config-dev.toml

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



kafka consume result:

<img width="492" alt="image" src="https://user-images.githubusercontent.com/7270440/189933224-eb648850-d3c8-451c-a249-c45179a2aa6c.png">



##  redis send to kafka

run server

./ho redis kafka -p /xxxx/golang/ho/configs -f config-dev-redis.toml


connect server
- telnet 127.0.0.1 9192

set data 
- set test_mc 123456

<img width="317" alt="image" src="https://user-images.githubusercontent.com/7270440/189932884-dce6de03-f686-4f50-8ed9-a640156e6c2e.png">



kafka config

```toml
title="redis-to-kafka"
[main]
serverIp="127.0.0.1"
serverPort=9192
mainLogPath="/tmp/ho_redis_2_kafka"
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


kafka consume result:

<img width="511" alt="image" src="https://user-images.githubusercontent.com/7270440/189933129-2792d1f2-0894-4a4c-b0f2-92ccfb23f235.png">



## kafka to kafka

1. simple config 

```toml
title="kafka-to-kafka"
[main]
serverIp="127.0.0.1"
serverPort=9193
mainLogPath="/tmp/ho_kafka_2_kafka"
mainLogModel=3
showSaramaDebug=true
[kafkas]
    [kafkas.test_mc]
    brokers=["tt.com:18888"]
    topic="test_mc"
    group="test"
    sslEnable=false
    user=""
    password=""
    producerOn=false
    [kafkas.test_mc2]
    brokers=["tt.com:18888"]
    topic="test_mc2"
    group=""
    sslEnable=false
    user=""
    password=""
    producerOn=true
```

2. run server

```
./ho kafka -p /you config file path/golang/ho/configs -f config-dev-kafka-2kafka.toml -r test_mc -t test_mc2
```

3. push data into kafka




## dependent on

a memmcache server: [go-memcache](https://github.com/rpcxio/gomemcached)

a common framework: [ngo](https://github.com/NetEase-Media/ngo)

a redis server:[redcon](https://github.com/tidwall/redcon)









