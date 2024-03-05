#!/bin/bash

minikube start --driver=docker
minikube addons enable metrics-server
minikube dashboard&