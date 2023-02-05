go build -o splunk
env --debug $(cat .env | grep -v '^#') ./splunk
