curl 127.0.0.1:3000/api/v1/ping

if [ $? -eq 0 ]; then
    echo "API available"
else
    echo "API non available"
    exit 1
fi