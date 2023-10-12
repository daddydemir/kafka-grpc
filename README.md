# kafka-grpc
----------------------


![image](https://github.com/daddydemir/kafka-grpc/assets/42716195/797d44a0-be0a-4037-b5e4-cc3ac303b946)


> Bu proje Producer ve Consumer olarak iki farkli bolumden olusmaktadir.


### KURULUM
-----

- PostgreSQL Container

```sh
docker run --name database -e POSTGRES_PASSWORD=password -d postgres
```

- Zookeeper Container

```sh
docker run --name zookeeper -p 2181:2181 zookeeper
```

- Kafka Container
```sh
docker run --name kafka -p 9092:9092 -e KAFKA_ZOOKEEPER_CONNECT=localhost:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092/ -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka
```
> localhost yazan kisimlarda bilgisayarinizin local ip adresi girilmesi gerekmektedir.

- Kafka-UI Container (gerekli degil)
```sh
docker run -p 8080:8080 -e DYNAMIC_CONFIG_ENABLED=true provectuslabs/kafka-ui
```

- Proto kodunu derlemek icin

```sh
protoc --go_out=plugins=grpc:proto proto/image_analyze.proto
```

## DOCKERIZE
---------

- Producer

```sh
docker build -t producer .
```
run 

```sh
docker run -p 5001:5001 producer
```

- Consumer

```sh
docker build -t consumer .
```

run 

```sh
docker run consumer
```


> Konfigürasyon ayarlarından kaynaklı olarak projeyi local ortamınızda çalıştırmakta zorlanabilirsiniz. 

