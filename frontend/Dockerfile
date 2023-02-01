FROM node:latest

# Create app directory
WORKDIR /app

# Install app dependencies
COPY package*.json ./

RUN npm install

# Bundle app source
COPY . .

EXPOSE 8080

# Execute the app
CMD ["npm" ,"run", "serve"]