## Step to Start Cluster

* Run `export CLUSTER_SECRET=$(od  -vN 32 -An -tx1 /dev/urandom | tr -d ' \n')` followed by `docker-compose up -d`
* Upload api-endpoint: http://localhost:8090/upload

* Download api-endpoint: `http://localhost:8090/file/hashID`