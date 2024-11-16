



![352691183-45f4ef12-c5b5-4247-9d43-356b5dfb671b](https://github.com/user-attachments/assets/2164eb5d-4bb0-499a-85a3-b4521c536be4)


# Web-app-DevOps
This project demonstrates the deployment of a web application using DevOps practices and tools to ensure efficient, scalable, and automated application delivery. 
Below are the key steps and technologies implemented in the project:

# Dockerization:

1. Created a Dockerfile to containerize the web application.
2. Built a Docker image for the application and pushed it to a container registry (e.g., Docker Hub, Amazon ECR).

# Kubernetes Setup:

1. Created a dedicated repository (k8s/manifests) to store Kubernetes manifests, ensuring version control for deployment configurations.
2. Set up a Kubernetes cluster using Amazon EKS (Elastic Kubernetes Service).
3. Configured ingress resources to manage external access to the application within the cluster.

# Helm Charts Configuration:

1. Designed Helm charts to package Kubernetes manifests, enabling reusable, scalable, and configurable deployments.

# GitOps with Argo CD:

1. Configured Argo CD for GitOps-based continuous deployment.
2. Used Argo CD to monitor and deploy the Helm charts from the Git repository to the EKS cluster, ensuring declarative and automated deployments.
CI/CD Pipeline:

Built a CI/CD pipeline using GitHub Actions and GitHub Workflows.
Automated testing, building, and pushing the Docker image.
Integrated the CI/CD pipeline to trigger deployments through Argo CD upon code changes.
