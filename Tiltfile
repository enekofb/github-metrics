
docker_build('enekofb/metrics', '.', dockerfile='Dockerfile')

k8s_yaml(helm('./.charts/metrics', name='dev'))
k8s_resource('metrics', port_forwards=8080)
