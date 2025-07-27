#!/bin/bash

echo "🚀 Deploying WORQ Todo App to Kubernetes..."

# Apply manifests in order
echo "📁 Creating namespace..."
kubectl apply -f namespace/

echo "💾 Setting up storage..."
kubectl apply -f storage/

echo "🗄️ Deploying MongoDB..."
kubectl apply -f mongodb/

echo "⚙️ Deploying Backend..."
kubectl apply -f backend/

echo "🎨 Deploying Frontend..."
kubectl apply -f frontend/

echo "🌐 Setting up Ingress..."
kubectl apply -f ingress/

echo "✅ Deployment complete!"
echo ""
echo "📊 Check status:"
echo "kubectl get all -n worq-todo"
echo ""
echo "🌍 Access app:"
echo "Add '127.0.0.1 worq-todo.local' to /etc/hosts"
echo "Then visit: http://worq-todo.local" 