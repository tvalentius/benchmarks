# FLASK_APP=web-server.py flask run --port=1337

from flask import Flask
app = Flask(__name__)

@app.route("/")
def hello():
    return "Halo Bali!"
