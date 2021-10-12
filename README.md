# Kubernetes Inventory Tools

# Description:
The tool allows you to list the resources available on a Kubernetes cluster and to make a static report in the form of an HTML file.

The account when running the script is taken from the .kube / config file.

# How works?

The account used must have read access (list) on the following resources:
- nodes
- namespaces
- secrets
- configmaps
- deployment
- pods
- pvcs

