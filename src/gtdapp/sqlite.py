import sqlite3

conn = sqlite3.connect("todo.db")
print("Opened database successfully")

conn.execute("""CREATE TABLE ITEM
         (ID INT PRIMARY KEY     NOT NULL,
         TITLE           TEXT    NOT NULL,
         DESCRIPTION     CHAR(200));""")
print("Table created successfully")

conn.close()
