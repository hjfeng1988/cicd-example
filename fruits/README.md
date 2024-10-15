目录结构
```
[root@hjfeng example]# tree
.
├── app0
│   ├── base
│   │   ├── deployment.yaml
│   │   └── kustomization.yaml
│   ├── dev
│   │   ├── kustomization.yaml
│   │   └── patch.yaml
│   ├── prod
│   │   ├── kustomization.yaml
│   │   └── patch.yaml
│   └── staging
│       ├── kustomization.yaml
│       └── patch.yaml
```

示例
```
[root@hjfeng example]# kustomize build app0/base
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app0
spec:
  selector:
    matchLabels:
      app: app0
  template:
    metadata:
      labels:
        app: app0
    spec:
      containers:
      - image: app0:0.0.0
        name: app0
[root@hjfeng example]# kustomize build app0/dev
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app0
spec:
  replicas: 1
  selector:
    matchLabels:
      app: app0
  template:
    metadata:
      labels:
        app: app0
    spec:
      containers:
      - image: registery.your.com/app0:0.0.1
        name: app0
[root@hjfeng example]# kustomize build app0/staging
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app0
spec:
  replicas: 2
  selector:
    matchLabels:
      app: app0
  template:
    metadata:
      labels:
        app: app0
    spec:
      containers:
      - image: registery.your.com/app0:0.0.2
        name: app0
[root@hjfeng example]# kustomize build app0/prod
apiVersion: apps/v1
kind: Deployment
metadata:
  name: app0
spec:
  replicas: 3
  selector:
    matchLabels:
      app: app0
  template:
    metadata:
      labels:
        app: app0
    spec:
      containers:
      - image: registery.your.com/app0:0.0.3
        name: app0
```
