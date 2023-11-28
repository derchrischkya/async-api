sleep 10
curl --silent "127.0.0.1:3000/api/v1/ping" > /dev/null
if [ $? -eq 0 ]; then
    echo "API available"
else
    echo "API non available"
    exit 1
fi

curl --silent "127.0.0.1:3000/api/v1/process/run"
if [ $? -eq 0 ]; then
    echo "API received message"
else
    echo "API not received message"
    exit 1
fi
exit 0