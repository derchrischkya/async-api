curl "http://127.0.0.1:4151/stats?topic=async-api&channel=backendDemoProcessor" 

if [ $? -eq 0 ]; then
    echo "NSQ topic available"
else
    echo "NSQ topic non available"
    exit 1
fi


curl -d "Hello World" "http://127.0.0.1:4151/pub?topic=async-api&channel=backendDemoProcessor"

if [ $? -eq 0 ]; then
    echo "Message sent to NSQ"
else
    echo "Message not sent to NSQ"
    exit 1
fi
