# Bare bone API using golang

# Running the API
1. Go to the project directory: /play-api.
2. Open terminal on that directory and type: ./run.sh. You should see Listening to http://localhost:5000/.

# Testing the API
Type the following URL in your browser.
1. localhost:5000/ - this should show the following message: {"msg":"Congrats! It works!"}
2. localhost:5000/test?email=youremail@example.com - this should show the following message: {"msg":"youremail@example.com"}
