from flask import jsonify, reqest
from webargs import fields
from webargs.flaskparser import use_args

from app import app

auth_args = {
    'id':fields.Str(),
    'pass':fields.Str()
}


@mod_api.route('/user')
@use_args(auth_args)
def user(args):
    print args['id']
    print args['pass']
    return jsonify(args)


