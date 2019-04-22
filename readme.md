### to compile and run locally

```bash
$chenglim@chenglim-GL503VM:~/work/src/github.com/teochenglim/gocal$ pwd
/home/chenglim/work/src/github.com/teochenglim/gocal
$ dep init
$ go build -o rest main.go
$ ./rest
$ ## Terminal 2
$ curl -s "localhost:3000/mul?num1=4&num2=2"
$ curl -s "localhost:3000/mul?num1=4&num2=0"
```

### to deploy to the Kubernetes

```bash
$ # build docker container locally
$ docker run -it --name gocal teochenlgim/gocal:latest
$ # docker push to docker.io
$ docker push teoochenglim/gocal:latest
$ # apply the deployment to kubernetes cluster
$ kubectl apply -f deployment.yml
$ # apply the service to kubernetes cluster
$ kubectl apply -f service.yml
$ # test the kubernetes service
$ curl -s "localhost:32160/add?num1=2&num2=4"
{ "num1": 2.000000, "num2": 4.000000, "answers": 6.000000 }
```

### to deploy Grafana and Prometheus
```bash
$ helm install stable/prometheus --namespace monitoring  --name prometheus
$ helm install stable/grafana --namespace monitoring --name grafana
```
