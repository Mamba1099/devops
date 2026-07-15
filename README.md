# 🚀 DevOps Template - Go Application with Kubernetes

A production-ready DevOps template with a Go backend, Docker, Kubernetes, CI/CD, and monitoring.

![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)
![Docker](https://img.shields.io/badge/Docker-24.0-blue.svg)
![Kubernetes](https://img.shields.io/badge/Kubernetes-1.27-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

---

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [Local Development](#local-development)
- [Docker](#docker)
- [Kubernetes](#kubernetes)
- [CI/CD Pipeline](#cicd-pipeline)
- [Monitoring](#monitoring)
- [API Endpoints](#api-endpoints)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)

---

## 📖 Overview

This project demonstrates a complete DevOps pipeline for a Go application with:

- ✅ **Containerization** with Docker
- ✅ **Orchestration** with Kubernetes (k3d)
- ✅ **CI/CD** with GitHub Actions
- ✅ **Configuration Management** with Kustomize
- ✅ **Monitoring** with Prometheus & Grafana
- ✅ **Database** with PostgreSQL
- ✅ **Auto-scaling** with HPA
- ✅ **Health Checks** with Liveness & Readiness Probes

---

## 🛠️ Technologies Used

| **Category** | **Technology** | **Version** |
|--------------|----------------|-------------|
| **Language** | Go | 1.23 |
| **Containerization** | Docker | 24.0+ |
| **Orchestration** | Kubernetes (k3d) | 1.27+ |
| **CI/CD** | GitHub Actions | Latest |
| **Configuration** | Kustomize | 4.5+ |
| **Database** | PostgreSQL | 15 |
| **Monitoring** | Prometheus + Grafana | Latest |
| **HTTP Router** | Gorilla Mux | 1.8+ |

---

## 📋 Prerequisites

| **Tool** | **Version** | **Installation** |
|----------|-------------|------------------|
| **Go** | 1.23+ | [Download](https://golang.org/dl/) |
| **Docker** | 24.0+ | [Download](https://docs.docker.com/get-docker/) |
| **kubectl** | 1.27+ | [Download](https://kubernetes.io/docs/tasks/tools/) |
| **k3d** | 5.0+ | `curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh \| bash` |
| **Git** | 2.30+ | [Download](https://git-scm.com/downloads) |
| **PostgreSQL** | 15 | `docker run -d --name postgres -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydb -p 5436:5432 postgres:15-alpine` |

---

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone https://github.com/mambak10099/devops.git
cd devops
```
### 2. Start PostgreSQL
```bash
docker run -d \
  --name postgres \
  -e POSTGRES_USER=user \
  -e POSTGRES_PASSWORD=password \
  -e POSTGRES_DB=yourdb \
  -p 5436:5432 \
  postgres:15-alpine
```
### 3. Run with Docker Compose
```bash
docker-compose up -d --build
```
### 4. Test the App
```bash
curl http://localhost:3000/health
```
### 5. Deploy Kubernetes
```bash
# Create k3d cluster
k3d cluster create myapp-cluster --servers 1 --agents 2

# Deploy
kubectl apply -k k8s/overlays/dev

# Port forward
kubectl port-forward -n dev service/myapp-service 8080:80

# Test
curl http://localhost:8080/health
```
## 📁 Project Structure
```bash
devops/
├── app/
│   ├── cmd/
│   │   └── main.go                 # Entry point
│   ├── internal/
│   │   ├── config/                 # Configuration
│   │   ├── database/               # Database connection
│   │   ├── handlers/               # HTTP handlers
│   │   ├── middleware/             # Middleware
│   │   ├── models/                 # Data models
│   │   └── routes/                 # Route definitions
│   ├── go.mod                      # Go module
│   └── go.sum                      # Dependencies
├── k8s/
│   ├── base/                       # Base K8s manifests
│   │   ├── deployment.yaml
│   │   ├── service.yaml
│   │   ├── configmap.yaml
│   │   ├── ingress.yaml
│   │   └── kustomization.yaml
│   └── overlays/                   # Environment overlays
│       └── dev/
│           ├── kustomization.yaml
│           ├── replica-count.yaml
│           └── env-dev.yaml
├── monitoring/
│   └── prometheus/
│       └── prometheus.yml
├── .github/workflows/
│   ├── ci.yml                      # CI Pipeline
│   └── cd.yml                      # CD Pipeline
├── Dockerfile
├── docker-compose.yml
├── .env
├── .gitignore
└── README.md
```
## 💻 Local Development
### Install Dependencies
```bash
cd app
go mod download
```
### Run Locally
```bash
export DATABASE_URL="postgresql://username:your-password@localhost:5436/yourdb?sslmode=disable"
go run ./cmd/
```
## 🐳 Docker
### Build Image
```bash
docker build -t your-dockerhub-username/devops-app:latest .
```
### Run Container
```bash
docker run -d -p 3000:3000 --name myapp your-dockerhub-username/devops-app:latest
```
### Run with Docker Compose
```bash
docker-compose up -d --build
docker-compose ps
docker-compose logs -f
docker-compose down
```
## ☸️ Kubernetes
### Create CLuster
```bash
k3d cluster create myapp-cluster --servers 1 --agents 2
kubectl get nodes
```
### Deploy Application
```bash
# Create namespace
kubectl create namespace dev

# Deploy
kubectl apply -k k8s/overlays/dev

# Check status
kubectl get pods -n dev
kubectl get deployments -n dev
kubectl get services -n dev
```
### Scale Application
```bash
# Manual scaling
kubectl scale deployment myapp -n dev --replicas=5

# Auto-scaling (HPA)
kubectl get hpa -n dev
```
### Access Application
```bash
kubectl port-forward -n dev service/myapp-service 8080:80
curl http://localhost:8080/health
```
## 🔄 CI/CD Pipeline
### GitHub Actions Workflows
#### CI Pipeline (.github/workflows/ci.yml)
Trigger: On push to main/develop or pull requests

Steps:

-- Checkout code

-- Setup Go

-- Install dependencies

-- Run tests

-- Build Docker image

-- Security scan

#### CD Pipeline (.github/workflows/cd.yml)
Trigger: On push to main

Steps:

Build Docker image

Push to Docker Hub

Deploy to Kubernetes

## 📊 Monitoring
### Prometheus and Grafana are run separately via Docker Compose:
```bash
docker-compose up -d prometheus grafana
```
####  Access Prometheus
```bash
http://localhost:9090
```
#### Access Grafana
It required authentication use the default passwords

Username: admin

Password: admin
```bash
http://localhost:3001 (admin/admin)
```


