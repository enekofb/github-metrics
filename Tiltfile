
docker_build('enekofb/metrics', '.', dockerfile='Dockerfile')

# Install resources I couldn't find elsewhere
k8s_yaml(listdir('tools', recursive=True))

k8s_yaml(helm('./.charts/metrics', name='dev'))
k8s_resource('metrics', port_forwards=8080)


