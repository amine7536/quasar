#!/usr/bin/env bash

DOCKERNET="quasarnet"
QUASARIP="172.18.0.2"
GOBGPIP="172.18.0.3"

case "$1" in
	"docker-build")
		echo "Building $2"
		docker build -f $2.Dockerfile --rm -t $2 .
	;;
	
	"quasar") 
    	echo "Start Quasar Dev container with:" 
    	echo "#> docker run --net $DOCKERNET --ip $QUASARIP --hostname quasar-dev -v $(pwd):/go/src/github.com/amine7536/quasar -it $1"
    	docker run --net $DOCKERNET --ip $QUASARIP -v $(pwd):/go/src/github.com/amine7536/quasar -it $1
  	;;
  
	"gobgp") 
		echo "Start Gobgp Dev container with:" 
    	echo "#> docker run --net  $DOCKERNET --ip $QUASARIP --hostname gobgp-dev -it $1"
    	docker run --net  $DOCKERNET --ip $GOBGPIP --hostname gobgp-dev -it $1
  	;;

	"help") 
  		echo "Helper script to bootstrap Dev env"
    	echo "Make sure to have Docker Network 'quasarnet' created with : "
    	echo "#> docker network create --subnet=172.18.0.0/16 quasarnet"
  	;;
  
	"*") 
		echo "Invalid Option"
  	;;

esac