# clonTwitter
 This is a Twitter clon, made following the tutorial of "Learn Go language from Zero" (Aprende lenguaje GO (GOLANG) desde 0).
 The project helped me to understand the basics of the Go language, manipulation of non-relational databases using mongoDB, testing http requests using Postman and deployement of apps using the cloud platform Heroku.
## Usage
### MongoDB
To access to the Mongo database as guest, you need to paste the connection string alocated in the file "mongoUri.txt" into a New connection of mongo compass.
### Postman
At first, you need to fork the collection of available requests from postman, clicking the button below.

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/18697537-c74b8786-72c3-4129-abc8-68389d84bfb0?action=collection%2Ffork&collection-url=entityId%3D18697537-c74b8786-72c3-4129-abc8-68389d84bfb0%26entityType%3Dcollection%26workspaceId%3D2245a3c9-3b7d-4f2c-884c-d24fdcc6649d#?env%5BLocal%5D=W3sia2V5IjoiUnV0YSIsInZhbHVlIjoibG9jYWxob3N0OjgwODAiLCJlbmFibGVkIjp0cnVlLCJzZXNzaW9uVmFsdWUiOiJodHRwOi8vbG9jYWxob3N0OjgwODAiLCJzZXNzaW9uSW5kZXgiOjB9XQ==)

After that, set an new environment to place the path (Ruta) as a variable for connecting to heroku. It should look like [this](https://www.postman.com/technical-cosmonaut-67158360/workspace/my-workspace/environment/18697537-08c9534b-252c-48ce-8003-4ebce6c859bd)

The clon also uses a token that expires at certain time! It's putted as a variable inside the Heroku environment. You migth saw it when you entered to the link above. So if the app fails, you probably have to Log In ("Login" request) again and overwrite the old token with the new one 
 
Done! Now you can do http requests and visualize them in the mongo database.






