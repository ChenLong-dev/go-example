#!/bin/bash

images=(

    kube-apiserver:v1.25.4

    kube-controller-manager:v1.25.4

    kube-scheduler:v1.25.4

    kube-proxy:v1.25.4

    #pause:3.8

    #etcd:3.5.5-0

    #coredns:v1.9.3

)

for imageName in ${images[@]};

do

    docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/${imageName}

    docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/${imageName} registry.k8s.io/${imageName}

    docker rmi registry.cn-hangzhou.aliyuncs.com/google_containers/${imageName}

done
