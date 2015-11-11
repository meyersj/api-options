from flask import Flask, jsonify
from webargs import fields
from webargs.flaskparser import use_args

api = Flask(__name__)

auth_args = {
    'id':fields.Str(),
    'pass':fields.Str()
}

def build_response(success, msg):
    return {'Msg':msg, 'Success':success}

def authenticate(user_id, password):
    if user_id == "admin" and password == "admin":
        return build_response(True, "User authenticated")
    return build_response(False, "Try user:admin password:admin")

@api.route('/')
def index():
    return "Index"

@api.route('/user')
@use_args(auth_args)
def user(args):
    # build default response message
    res = build_response(False, "Missing id and/or pass parameters")
    try:
        user_id = args['id']
        password = args['pass']
        # user_id and password parameters exist so authenticate them
        res = authenticate(user_id, password)
    except:
        pass
    # return json response
    return jsonify(res)

if __name__ == '__main__':
    print "Starting server at 127.0.0.1:5000"
    api.run()
