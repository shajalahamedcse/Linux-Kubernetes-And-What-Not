## Etcd & Kubernetes, A Closer Look


###  What is Etcd ?
Etcd is nothing but a distributed, reliable key-value store. Kubernetes uses it to store the entire state of the cluster: it's configuration and specifications.

###  kubernetes and  Ectd :
When we deploy any kind of service or apps kubernetes write about it into Etcd. Let's do some hands on. It will be clear then.

### Minikube :
We need minikube installed in our computer to this hands on. Lets fire up a kubernetes cluster.

	$ minikube start 
  
  
  ### The Etcd Pod :
  
  First, let’s list all the Pods running in the cluster:
  
    $ kubectl get pods --all-namespaces
    
      NAMESPACE              NAME                                        READY   STATUS    RESTARTS   AGE
      kube-system            coredns-5644d7b6d9-htmvs                     1/1     Running   1          107m
      kube-system            coredns-5644d7b6d9-q8bx7                     1/1     Running   1          107m
      kube-system            etcd-minikube                                1/1     Running   1          106m
      kube-system            kube-apiserver-minikube                      1/1     Running   1          106m
      kube-system            kube-controller-manager-minikube             1/1     Running   1          106m
      kube-system            kube-proxy-cf46m                             1/1     Running   1          107m
      kube-system            kube-scheduler-minikube                      1/1     Running   1          105m

   In kube-system namespace we can find **etcd-minikube** POD which in we are interested.

### etcdctl
As we want to query into etcd , we will use **etcdctl** for that. To work with etcdctl we need CA file to accept Etcd server certificate. We will copy certificates from **kube-apiserver** container and inject them into **Etcd container**. 

	$ kubectl cp --namespace kube-system kube-apiserver-minikube:var/lib/minikube/certs/apiserver-etcd-client.crt apiserver-etcd-client.crt
	$ kubectl cp --namespace kube-system apiserver-etcd-client.crt etcd-minikube:var/lib/minikube/certs/
	$ kubectl cp --namespace kube-system kube-apiserver-minikube:var/lib/minikube/certs/apiserver-etcd-client.key apiserver-etcd-client.key
	$ kubectl cp --namespace kube-system apiserver-etcd-client.key etcd-minikube:var/lib/minikube/certs/

Now, let's exec into Etcd container:

    $ kubectl exec -it --namespace kube-system etcd-minikube sh
 
 We need to set etcdctl tool to **v3 API** version using following environment variable:
 
	$ export ETCDCTL_API=3
	$ cd /var/lib/minikube/certs
	
Now, moment of truth:
		
	$ etcdctl --cacert="etcd/ca.crt" --key=apiserver-etcd-client.key --cert=apiserver-etcd-client.crt get /registry/deployments --prefix --keys-only

	/registry/deployments/kube-system/coredns

	/registry/deployments/kube-system/nginx-ingress-controller

	/registry/deployments/kubernetes-dashboard/dashboard-metrics-scraper

	/registry/deployments/kubernetes-dashboard/kubernetes-dashboard 

Yes !!!!!! We hacked the most secure key-value store in the world. Cheap joke, right?
 
  I will create an alias to save some keystroke.
		     
    $ alias ktcdl="etcdctl --cacert="etcd/ca.crt" --key=apiserver-etcd-client.key --cert=apiserver-etcd-client.crt"

Now let's use that alias.

    $ ktcdl get /registry/deployments --prefix --keys-only

	/registry/deployments/kube-system/coredns
	/registry/deployments/kube-system/nginx-ingress-controller
	/registry/deployments/kubernetes-dashboard/dashboard-metrics-scraper
	/registry/deployments/kubernetes-dashboard/kubernetes-dashboard

These are the deployments available in the cluster.

Let's deploy nignx in this cluster.

	$ kubectl delete deployment nginx

Now if we check **etcd** again.

	$ ktcdl get /registry/deployments --prefix --keys-only


### Read data value

Since the data in the default etcd of k8s is stored in protobuf format, **etcdctl** result is unreadable.

	$ ktcdl get /registry/deployments/default/nginx -w=json
	
	{"header":{"cluster_id":15344640958435093190,"member_id":5063703067157974452,"revision":18477,"raft_term":3},"kvs":[{"key":"L3JlZ2lzdHJ5L2RlcGxveW1lbnRzL2RlZmF1bHQvbmdpbng=","create_revision":18245,"mod_revision":18280,"version":5,"value":"azhzAAoVCgdhcHBzL3YxEgpEZXBsb3ltZW50EuoECoABCgVuZ2lueBIAGgdkZWZhdWx0IgAqJDEzZWQzOGRkLTE1OWQtNDcxYS04NTUzLTlmNThjYTg0ZGRlNDIAOAFCCAidlfjzBRAAWgwKA2FwcBIFbmdpbnhiJgohZGVwbG95bWVudC5rdWJlcm5ldGVzLmlvL3JldmlzaW9uEgExegAS8QEIARIOCgwKA2FwcBIFbmdpbngaqgEKIAoAEgAaACIAKgAyADgAQgBaDAoDYXBwEgVuZ2lueHoAEoUBEkAKBW5naW54EgVuZ2lueCoAQgBqFC9kZXYvdGVybWluYXRpb24tbG9ncgZBbHdheXOAAQCIAQCQAQCiAQRGaWxlGgZBbHdheXMgHjIMQ2x1c3RlckZpcnN0QgBKAFIAWABgAGgAcgCCAQCKAQCaARFkZWZhdWx0LXNjaGVkdWxlcsIBACInCg1Sb2xsaW5nVXBkYXRlEhYKCQgBEAAaAzI1JRIJCAEQABoDMjUlKAAwCjgASNgEGvABCAEQARgBIAEoADJlCglBdmFpbGFibGUSBFRydWUiGE1pbmltdW1SZXBsaWNhc0F2YWlsYWJsZSokRGVwbG95bWVudCBoYXMgbWluaW11bSBhdmFpbGFiaWxpdHkuMggIqpX48wUQADoICKqV+PMFEAAyewoLUHJvZ3Jlc3NpbmcSBFRydWUiFk5ld1JlcGxpY2FTZXRBdmFpbGFibGUqOlJlcGxpY2FTZXQgIm5naW54LTg2YzU3ZGI2ODUiIGhhcyBzdWNjZXNzZnVsbHkgcHJvZ3Jlc3NlZC4yCAiqlfjzBRAAOggInZX48wUQADgBGgAiAA=="}],"count":1}

That's it for today.....
