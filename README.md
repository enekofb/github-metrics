# Github metrics poc 

A quick poc to understand some github and zenhub metrics capabilities 

In particular the following user stories

```
Given a github repo
For a given period of time ~ less say a month
I want to get a metric of the #nummer of issues of a type by different categories
Where 
- categories are labels
- issues are labels
```

## Requirements
- kind
- tilt 
- flux

## Getting Started

For development just execute
- make kind
- make dev

Get prometheus stack installed via [helm](./tools/prometheus.yaml)

Once everything is ready get grafana to visualice available metrics

```yaml
kubectl port-forward -n prometheus svc/prometheus-prometheus-grafana 8000:80
```


