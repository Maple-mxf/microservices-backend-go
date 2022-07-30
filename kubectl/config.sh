
# 配置集群名称与服务地址
sudo kubectl config --kubeconfig=${HOME}/.kube/config set-cluster cluster-name --server=https://172.20.0.2:46555 --insecure-skip-tls-verify

sudo kubectl config --kubeconfig=${HOME}/.kube/config set-credentials admin --username=voyagerma --password=121101mxf

sudo kubectl config --kubeconfig=${HOME}/.kube/config set-context admin --cluster=cluster-name --namespace=test --user=voyagerma

sudo kubectl config --kubeconfig=${HOME}/.kube/config use-context admin

kubectl get pods

#
kubectl proxy

export KUBE_APISERVER="https://172.20.0.113:6443"
export KUBE_APISERVER="https://127.0.0.1:40297"
kubectl config set-cluster kubernetes \
--certificate-authority=/etc/kubernetes/ssl/ca.pem \
--embed-certs=true \
--server=${KUBE_APISERVER} \
--kubeconfig=devuser.kubeconfig

kubectl config set-credentials devuser \
--client-certificate=/etc/kubernetes/ssl/devuser.pem \
--client-key=/etc/kubernetes/ssl/devuser-key.pem \
--embed-certs=true \
--kubeconfig=devuser.kubeconfig

kubectl config set-context kubernetes \
--cluster=kubernetes \
--user=devuser \
--namespace=dev \
--kubeconfig=devuser.kubeconfig

kubectl config use-context kubernetes --kubeconfig=devuser.kubeconfig

helm --namespace observability install my-release signoz/signoz


NAME: my-release
LAST DEPLOYED: Sat Jul 23 15:18:08 2022
NAMESPACE: observability
STATUS: deployed
REVISION: 1
NOTES:
1. You have just deployed SigNoz cluster:

- frontend version: '0.10.0'
- query-service version: '0.10.0'
- alertmanager version: '0.23.0-0.1'
- otel-collector version: '0.45.1-1.1'
- otel-collector-metrics version: '0.45.1-1.1'

2. Get the application URL by running these commands:

  export POD_NAME=$(kubectl get pods --namespace observability -l "app.kubernetes.io/name=signoz,app.kubernetes.io/instance=my-release,app.kubernetes.io/component=frontend" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:3301 to use your application"
  kubectl --namespace observability port-forward $POD_NAME 3301:3301

export SERVICE_NAME=$(kubectl get svc --namespace observability -l "app.kubernetes.io/component=frontend" -o jsonpath="{.items[0].metadata.name}")

curl -sL https://github.com/SigNoz/signoz/raw/main/sample-apps/hotrod/hotrod-install.sh \
  | HELM_RELEASE=my-release SIGNOZ_NAMESPACE=observability bash


export SERVICE_NAME=$(kubectl get svc --namespace observability -l "app.kubernetes.io/component=frontend" -o jsonpath="{.items[0].metadata.name}")
kind create cluster --config ~/k8s-3nodes.yaml
cat <<EOF | ./kind create cluster --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
EOF



