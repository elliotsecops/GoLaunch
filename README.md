# GoLaunch: Comprehensive and Secure CI/CD Solution for Go Web Applications on Kubernetes

## Overview

**GoLaunch** is a comprehensive and secure CI/CD solution designed for seamlessly deploying Go web applications to Kubernetes. Leveraging the power of GitHub Actions and AWS, GoLaunch automates the entire software delivery lifecycle, from initial code commit to production deployment. This project emphasizes security and robustness, employing AWS Secrets Manager for secure secret storage and retrieval, and implementing strict network policies for enhanced pod security within the Kubernetes cluster.

GoLaunch goes beyond basic CI/CD by incorporating best practices for resource management, including carefully tuned resource requests and limits, resource quotas, and limit ranges, ensuring efficient and reliable application performance. Health checks, including startup, liveness, and readiness probes, further enhance the resilience of deployed applications. The pipeline is triggered automatically on every push to the main branch, providing continuous integration and delivery. This project serves as a practical example of building a production-ready CI/CD pipeline for Go applications on Kubernetes, offering a clear and adaptable template for your own deployments.

## Technologies Used

- **GitHub Actions**: For CI/CD automation.
- **Git**: Version control.
- **Kubernetes**: Container orchestration.
- **Docker**: Containerization.
- **AWS**: For cloud services, including EKS (Elastic Kubernetes Service) and AWS Secrets Manager.

## Project Structure

- `main.go`: The Go web application.
- `Dockerfile`: Dockerfile for containerizing the application.
- `.github/workflows/ci-cd.yml`: GitHub Actions workflow for CI/CD.
- `config/`: Kubernetes configuration files.
  - `deployment.yaml`: Kubernetes deployment manifest.
  - `network-policies.yaml`: Kubernetes network policies.
  - `resource-quotas.yaml`: Kubernetes resource quotas.
  - `limit-range.yaml`: Kubernetes limit ranges.
- `README.md`: Project documentation.

## Getting Started

### Prerequisites

- Go 1.21
- Docker
- Kubernetes cluster (AWS EKS recommended)
- GitHub account
- AWS account with IAM roles and policies configured

### Steps

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/elliotsecops/golaunch.git
   cd golaunch
   ```

2. **Build and Run Locally:**
   ```bash
   go run main.go
   ```

3. **Build Docker Image:**
   ```bash
   docker build -t go-web-app .
   ```

4. **Run Docker Container:**
   ```bash
   docker run -p 8080:8080 go-web-app
   ```

5. **Push to GitHub:**
   ```bash
   git add .
   git commit -m "Initial commit"
   git push origin main
   ```

6. **Deploy to Kubernetes:**
   - Ensure your Kubernetes cluster is configured.
   - Apply the Kubernetes manifests:
     ```bash
     kubectl apply -f config/
     ```

## CI/CD Pipeline

### GitHub Actions Workflow

The workflow is defined in `.github/workflows/ci-cd.yml`. It performs the following steps:

1. **Build and Test**: Builds the Go application, runs tests, and builds the Docker image.
2. **Push Image**: Pushes the Docker image to GitHub Container Registry.
3. **Deploy**: Deploys the application to the Kubernetes cluster.

### Environment Variables

- `REGISTRY`: GitHub Container Registry.
- `IMAGE_NAME`: Docker image name.
- `PORT`: Application port.
- `AWS_REGION`: AWS region for EKS.

### Secrets

- `AWS_ACCOUNT_ID`: AWS account ID.
- `GITHUB_TOKEN`: GitHub token for authentication.

## Security and Robustness

### AWS Secrets Manager

GoLaunch uses AWS Secrets Manager to securely store and retrieve secrets. This ensures that sensitive information is managed securely and can be rotated easily.

### Network Policies

Strict network policies are implemented to enhance pod security within the Kubernetes cluster. These policies follow the principle of least privilege, allowing only necessary traffic and restricting access to sensitive resources.

### Resource Management

#### Container Resources

Resource limits and requests are carefully tuned based on typical Go application profiles:

```yaml
resources:
  limits:
    cpu: "1"    # Maximum CPU usage
    memory: "1Gi"  # Maximum memory usage
  requests:
    cpu: "500m"    # Guaranteed CPU allocation
    memory: "512Mi"  # Guaranteed memory allocation
```

#### Resource Quotas

Resource quotas prevent resource exhaustion at the namespace level:

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: go-web-app-quota
spec:
  hard:
    requests.cpu: "2"
    requests.memory: 2Gi
    limits.cpu: "4"
    limits.memory: 4Gi
    pods: "10"
```

#### Limit Ranges

Limit ranges set default limits for containers, ensuring consistent resource allocation:

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: go-web-app-limits
spec:
  limits:
  - default:
      cpu: "1"
      memory: 1Gi
    defaultRequest:
      cpu: "500m"
      memory: 512Mi
    type: Container
```

### Health Checks

Health checks, including startup, liveness, and readiness probes, enhance the resilience of deployed applications:

```yaml
startupProbe:
  httpGet:
    path: /health
    port: 8080
  failureThreshold: 30
  periodSeconds: 10
livenessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
readinessProbe:
  httpGet:
    path: /health
    port: 8080
  initialDelaySeconds: 5
  periodSeconds: 10
  timeoutSeconds: 5
  failureThreshold: 3
```

## Monitoring and Maintenance

### Resource Usage

Monitor resource usage patterns and adjust limits based on metrics:

```bash
kubectl top pods -l app=go-web-app
```

### Network Policy Validation

Verify network policy logs and DNS resolution:

```bash
kubectl describe networkpolicy go-web-app-policy
```

### Regular Review

- **Monitor resource usage patterns**
- **Adjust limits based on metrics**
- **Review network policy effectiveness**

## Troubleshooting

### Resource Issues

- **Check for OOMKilled status**
- **Monitor CPU throttling**
- **Review resource quotas**

### Network Issues

- **Verify network policy logs**
- **Check for blocked connections**
- **Validate DNS resolution**

## Best Practices

### Resource Management

- **Always set both requests and limits**
- **Monitor actual usage**
- **Plan for peak loads**

### Network Security

- **Regular policy audits**
- **Monitor denied connections**
- **Keep policies updated**

## Contributing

Feel free to contribute by opening issues or pull requests.

## License

This project is licensed under the MIT License.
