echo "Build latest Go server"
env GOOS=linux GOARCH=amd64 go build main.go

echo "Dockerize the binary"
docker build -t bitdecaygames/fireport:latest .

echo "Tarball the docker image"
docker save -o ./fireport-container.tar bitdecaygames/fireport:latest

echo "Scp main executable to server"
scp fireport-container.tar fireport01_ubuntu:

echo "Move to ansible directory and start setup playbook"
cd ansible/
ansible-playbook -i dev setup.yml

echo "Run deploy playbook"
ansible-playbook -i dev deploy.yml