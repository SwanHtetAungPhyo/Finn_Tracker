#!/bin/bash 

docker_container(){
    echo "Docker compose running ...."
    docker-compose up --build -d
}

user_service_curl(){
    echo "Curling the user-service ...."
    curl http://localhost/user/ | jq . | tee -a ./json/user_service.json
}

expense_service_curl(){
    echo "Curling the expense-service ...."
    curl http://localhost/expense/ | jq . | tee -a ./json/expense_service.json
}
clean_up(){
    ./clean.sh
}
main(){
    docker_container
    sleep 5
    mkdir json
    user_service_curl & 
    expense_service_curl &
    wait

    sleep 15
    echo "PWD is ${pwd}"    
}


main
