# golang-ci-cd-todo
### Requirements
Install jq tool for bash. https://stedolan.github.io/jq/
On mac you can use brew
``` bash
breq install jq 
```
On windows you can use choco
```
choco install jq
```
And Ubuntu
```
sudo apt-get install jq
```
install kubectl

Install helm. https://helm.sh/

### Running locally with kubernetes
To fully run this project with local kubernetes, you will need to modify your host file (/etc/hosts in linux) and add the following entries:
``` bash
127.0.0.1       docker-registry.default
127.0.0.1       cicdexample.com
```

Then, you can install our local kubernetes solution, like docker for desktop with k8s support.

Go to the infra folder and run install-infra.sh to setup the requirements for the project, like a docker registry and the ingress controller.

Remember to have the correct k8s context setup for kubectl 

##### Helm

# - "User ID=postgres;Password=mypassword;Host=db;Port=5432;Database=postgres;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;"