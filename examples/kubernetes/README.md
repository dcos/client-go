# Kubernetes

This example [sets up a Kubernetes cluster on DC/OS enterprise][1].

## Prerequisites

In order to run this program, make sure the [prerequisites][2] are satisfied.

## Running the example

You can then run `go ./...` in order to create your kubernetes cluster.

Once the program finishes, your Kubernetes cluster should be up and running. You can
get its public IP through the following command:

    dcos task exec -it edgelb-pool-0-server curl ifconfig.co

[1]: https://docs.mesosphere.com/services/kubernetes/2.3.0-1.14.1/getting-started/
[2]: https://docs.mesosphere.com/services/kubernetes/2.3.0-1.14.1/getting-started/setting-up/#prerequisite-dcos-enterprise-cluster-enterprise