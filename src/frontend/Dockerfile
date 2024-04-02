FROM node:21-alpine3.18

WORKDIR /frontend

# Copy the frontend package and package-lock.json file
COPY  .  /frontend

# Install node packages
RUN npm install --force

RUN npm run build 

# Expose $PORT on container.
# We use a varibale here as the port is something that can differ on the environment.
EXPOSE 3000



# Start the app
CMD [ "npm", "run", "preview", "--host" ]