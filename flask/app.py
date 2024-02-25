from flask import Flask
import random
app = Flask(__name__)
@app.route('/')
def hello_world():
    return random.randint(0, 100)