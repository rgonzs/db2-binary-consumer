#!/usr/bin/env bash

folders=$(ls)
MAX_DAYS=90

echo "Escaneando archivos con ${MAX_DAYS} de antiguedad"

read -p "Desea continuar? " -n 1 -r
if [[ ! $REPLY =~ ^[Yy]$ ]]
echo ""
then
    for folder in $folders; do
        if [[ $folder = 'modasa' ]]; then
          continue 
        fi
        echo "Borrando los archivos de: ${folder}"
        find /home/abr15rec/mail/recepcionfacturas.pe/${folder}/cur -type f -mtime +$DAYS -delete
        echo $folder
    done
fi