### to compile and run locally

```bash
$ go build -o rest main.go
$ ## Terminal 1
$ ./rest
$ ## Terminal 2
$ curl -s "localhost:3000/mul?num1=4&num2=2"
$ curl -s "localhost:3000/mul?num1=4&num2=0"
```

### to deploy to the Kubernetes

```bash
$chenglim@chenglim-GL503VM:~/work/src/github.com/teochenglim/gocal$ pwd
/home/chenglim/work/src/github.com/teochenglim/gocal
$ dep init
$ # build docker container locally
$ docker build . -t teochenglim/gocal:latest
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
$ helm install stable/prometheus-operator --name prometheus-operator --namespace monitoring
$ # Grafana password
$ kubectl get secret --namespace monitoring prometheus-operator-grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
$ # Grafana
$ kubectl port-forward $(kubectl get pods --selector=app=grafana -n monitoring --output=jsonpath="{.items..metadata.name}") -n monitoring 3000
$ # prometheus
$ kubectl port-forward -n monitoring prometheus-prometheus-operator-prometheus-0 9090
$ # alertmanager
$ kubectl port-forward -n monitoring alertmanager-prometheus-operator-alertmanager-0 9093
```

### Load generator
```bash
$ for i in {1..3000}; do curl -s "localhost:32160/mul?num1=$i&num2=$i"; done
```
