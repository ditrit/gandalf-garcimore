//CL
./garcimore cluster init cluster 127.0.0.1:9000 
./garcimore cluster join cluster 127.0.0.1:9001 127.0.0.1:9000
./garcimore cluster join cluster 127.0.0.1:9002 127.0.0.1:9000

//A
./garcimore aggregator init agg1 titi 127.0.0.1:8000 127.0.0.1:9000
./garcimore aggregator init agg2 titi 127.0.0.1:8001 127.0.0.1:9000

//C
./garcimore connector init con1 titi 127.0.0.1:7000 127.0.0.1:7100 127.0.0.1:8000
./garcimore connector init con2 titi 127.0.0.1:7001 127.0.0.1:7101 127.0.0.1:8001

//W
./garcimore test send
./garcimore test receive