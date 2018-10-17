echo "Build latest Go server"
env GOOS=linux GOARCH=amd64 go build main.go
echo "SUCCESS"

echo "Scp main executable to server"
scp main ec2-34-217-206-186.us-west-2.compute.amazonaws.com:
echo "SUCCESS"

echo "Move to ansible directory and start setup playbook"
cd ansible/
ansible-playbook -i dev setup.yml
echo "SUCCESS"