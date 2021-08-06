#!/bin/bash
echo -e "\e[1;32m Setting up traffic over Whonix !! \e[0m"
ifconfig eth0 down
ifconfig eth0 10.152.152.14 netmask 255.255.192.0
route add default gw 10.152.152.10
mv /etc/resolv.conf /etc/resolv.conf.orig
echo "nameserver 10.152.152.10" > /etc/resolv.conf
ifconfig eth0 up
echo -e "\e[1;32m System is ready to talk to Whonix. You are on the Torified path enjoy Anonymity ! \e[0m"