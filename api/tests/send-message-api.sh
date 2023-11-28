curl 127.0.0.1:3000/api/v1/process/run

if [ $? -eq 0 ]; then
    echo "API received message"
else
    echo "API not received message"
    exit 1
fi