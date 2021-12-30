import sys
import tkinter
import sqlite
from sqlite import get_entries
import sqlite3

conn = sqlite3.connect("todo.db")

root = tkinter.Tk()

root.title("gtdapp")

entries = get_entries(conn)

root.mainloop()
