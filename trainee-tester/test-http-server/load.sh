#! /bin/bash

bombardier -n 300000 -c 50 --http1 "http://localhost:8090/ping"
bombardier -n 300000 -c 50 --http1 "http://localhost:8090/cdbsource?name=akton&number=385992026020"
