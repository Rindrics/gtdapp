from flask import Flask
import sqlite3

app = Flask(__name__)

@app.route("/")
def hello_world():
    html = render_template('index.html', title="gtdapp")
    return html

if __name__ == "__main__":
    app.run()
