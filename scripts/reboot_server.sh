#!/bin/bash
db_host=$(aws ssm get-parameters --names /Prod/NorthernLife/DB/Host --with-decryption --query Parameters[0].Value)
db_host=`echo $db_host | sed -e 's/^"//' -e 's/"$//'`
db_port=$(aws ssm get-parameters --names /Prod/NorthernLife/DB/Port --with-decryption --query Parameters[0].Value)
db_port=`echo $db_port | sed -e 's/^"//' -e 's/"$//'`
db_name=$(aws ssm get-parameters --names /Prod/NorthernLife/DB/Name --with-decryption --query Parameters[0].Value)
db_name=`echo $db_name | sed -e 's/^"//' -e 's/"$//'`
db_user=$(aws ssm get-parameters --names /Prod/NorthernLife/DB/User --with-decryption --query Parameters[0].Value)
db_user=`echo $db_user | sed -e 's/^"//' -e 's/"$//'`
db_password=$(aws ssm get-parameters --names /Prod/NorthernLife/DB/Password --with-decryption --query Parameters[0].Value)
db_password=`echo $db_password | sed -e 's/^"//' -e 's/"$//'`

export DB_HOST=$db_host
export DB_PORT=$db_port
export DB_NAME=$db_name
export DB_USER=$db_user
export DB_PASSWORD=$db_password

./main > /dev/null 2>&1 &
