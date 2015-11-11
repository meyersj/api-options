## running
You need to have pythons `virtualenv` installed and on your path

```
# create python virtual environment and fetch dependencies
./setup_env.sh

# run script with the virtual environments python
env/bin/python api.py
```

You should now have the api server running at 127.0.0.1:500

Endpoints supported:

+ Index: 127.0.0.1:5000/
+ User Authenicate: 127.0.0.1:5000/user?id=some_user_id&pass=some_password



