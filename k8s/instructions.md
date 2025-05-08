# Kubernetes Deployment Guide

This guide provides step-by-step instructions for deploying the Pinned Guest application to a Kubernetes cluster.

## Prerequisites

- Kubernetes cluster up and running
- `kubectl` configured and connected to your cluster
- Docker registry access (if using private images)
- Basic understanding of Kubernetes concepts

## 1. Environment Configuration

### 1.1 Application Secrets

1. Navigate to the `k8s` directory:
   ```bash
   cd k8s
   ```

2. Create your environment file:
   ```bash
   cp .env.example .env
   ```

3. Edit the `.env` file with your specific configuration values

4. Create the Kubernetes secret:
   ```bash
   kubectl create secret generic guest-secret \
     --from-env-file k8s/.env \
     --namespace=default
   ```

### 1.2 Docker Registry Authentication

If you're using a private Docker registry, you'll need to create a registry secret:

```bash
kubectl create secret docker-registry regsecret \
  --docker-server=$DOCKER_REGISTRY_SERVER \
  --docker-username=$DOCKER_USER \
  --docker-password=$DOCKER_PASSWORD \
  --docker-email=$DOCKER_EMAIL \
  --namespace=default
```

Required variables:
- `$DOCKER_REGISTRY_SERVER`: Your registry URL (e.g., `docker.io`, `ghcr.io`)
- `$DOCKER_USER`: Registry username
- `$DOCKER_PASSWORD`: Registry password or access token
- `$DOCKER_EMAIL`: (Optional) Your email address

## 2. Deployment

### 2.1 Deploy All Resources

To deploy all Kubernetes resources:

```bash
kubectl apply -f ./k8s
```

This will create:
- Deployments
- Services
- ConfigMaps
- Ingress rules (if configured)
- Other Kubernetes resources

### 2.2 Verify Deployment

Check the status of your deployment:

```bash
# Check pods
kubectl get pods

# Check services
kubectl get services

# Check deployments
kubectl get deployments
```

## 3. Troubleshooting

### 3.1 Common Issues

1. **Secret Creation Fails**
   - Ensure all required environment variables are set
   - Verify the `.env` file format is correct
   - Check for special characters in values

2. **Registry Authentication Fails**
   - Verify registry credentials
   - Check if the registry is accessible from your cluster
   - Ensure the secret is created in the correct namespace

3. **Pods Not Starting**
   - Check pod logs: `kubectl logs <pod-name>`
   - Describe pod: `kubectl describe pod <pod-name>`
   - Verify resource limits and requests

### 3.2 Useful Commands

```bash
# View pod logs
kubectl logs <pod-name>

# Describe resource
kubectl describe <resource-type> <resource-name>

# Get pod details
kubectl get pod <pod-name> -o yaml

# Check events
kubectl get events
```

## 4. Maintenance

### 4.1 Updating Secrets

To update existing secrets:

```bash
# Delete old secret
kubectl delete secret guest-secret

# Create new secret
kubectl create secret generic guest-secret --from-env-file k8s/.env
```

### 4.2 Scaling

Scale your deployment:

```bash
kubectl scale deployment <deployment-name> --replicas=<number>
```

## 5. Cleanup

To remove all deployed resources:

```bash
kubectl delete -f ./k8s
```

To remove secrets:

```bash
kubectl delete secret guest-secret
kubectl delete secret regsecret
```

## Support

If you encounter any issues or need assistance:
1. Check the [GitHub Issues](https://github.com/patricksferraz/pinned-guest/issues)
2. Review the application logs
3. Consult the Kubernetes documentation
