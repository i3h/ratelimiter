#!/bin/bash
ab -n 100000 -c 1000 http://127.0.0.1:5000/ > result.dat
echo

err=$(grep Non-2xx result.dat)
IFS=':' read -ra ERR <<< "$err"
err=$(echo ${ERR[1]} | xargs)

total=$(grep Complete result.dat)
IFS=':' read -ra TOTAL <<< "$total"
total=$(echo ${TOTAL[1]} | xargs)
 
time=$(grep 'Time taken for tests' result.dat)
IFS=' ' read -ra TIME <<< "$time"
time=$(echo ${TIME[4]} | xargs)

qps=$(echo "scale=2; ($total - $err) / $time" | bc -l)

echo "Err:   $err"
echo "Total: $total"
echo "Time:  $time"
echo "QPS:   $qps"
