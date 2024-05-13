# Builder stage
FROM node:22-alpine as builder
WORKDIR /app

# Install app dependencies by copying package.json and package-lock.json
COPY package*.json ./
RUN npm install

# Bundle app source by copying all other source files
COPY . .

# Build the React application
RUN npm run build

# Final stage
FROM node:22-alpine
WORKDIR /app

# Install 'serve' to serve your app
RUN npm install -g serve

# Copy the build directory from the builder stage
COPY --from=builder /app/build ./build

# Serve the static files
CMD ["serve", "-s", "build", "-l", "3000"]
