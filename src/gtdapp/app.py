from flask import Flask, render_template
import sqlite3
import sqlite
from sqlite import get_entries

app = Flask(__name__)

@app.route("/")
def hello_world():
    conn = sqlite3.connect("todo.db")
    entries = get_entries(conn)
    html = render_template('index.html', title="gtdapp", entries=entries)
    return html

if __name__ == "__main__":
    app.run()
