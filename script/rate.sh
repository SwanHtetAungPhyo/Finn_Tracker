#!/bin/bash 

LoadTest(){
    wrk -t12 -c400 -d30s --latency http://localhost/user/
}

LoadTest