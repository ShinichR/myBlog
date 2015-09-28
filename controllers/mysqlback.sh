#!/bin/bash
USER="root"
PASS="linger"
DBNAME="blog"

Now=$(date +"%d-%m-%Y--%H:%M:%S")
File=backup-$Now.sql
BACK_DIR="/data/blog"
LogFile="$BACK_DIR"/dbbakup.log


cd $BACK_DIR

mysqldump -u"root" -p"linger" "blog" > $File

echo -----------------------"$(date +"%y-%m-%d %H:%M:%S")"----------------------- >> $LogFile 
echo  createFile:$File >> $LogFile

#find "/data/backdata/" -cmin +1 -type f -name "*.sql" -print > deleted.txt

find $BACK_DIR -ctime +7 -type f -name "*.sql" -print > deleted.txt

echo -e "delete files:\n" >> $LogFile 
cat deleted.txt | while read LINE
do
    rm -rf $LINE
    echo $LINE>> $LogFile
done


echo "---------------------------------------------------------------" >> $LogFile
