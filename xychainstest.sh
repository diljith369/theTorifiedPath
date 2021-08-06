#!/bin/bash
IP=$(curl -s 'https://suip.biz/ip/')
echo -e "\e[1;32m Current Ip :  \e[0m" "$IP"
xyIP=$(proxychains curl -s 'https://suip.biz/ip/')
echo -e "\e[1;32m Ip over Torified Proxychains : \e[0m" "$xyIP"
