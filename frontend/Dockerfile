# Stage 1: Build React app
FROM node:18-alpine as build

WORKDIR /app

# Copy package files
COPY package*.json ./

# 🔧 Fix: Use npm install instead of npm ci --only=production
# We need devDependencies for building (vite, typescript, etc.)
RUN npm install

# Copy source code
COPY . .

# Build the app
RUN npm run build

# Stage 2: Serve with Nginx
FROM nginx:alpine

# Copy built app
COPY --from=build /app/dist /usr/share/nginx/html

# Copy nginx config
COPY nginx.conf /etc/nginx/nginx.conf

# Expose port
EXPOSE 80

# Start nginx
CMD ["nginx", "-g", "daemon off;"]