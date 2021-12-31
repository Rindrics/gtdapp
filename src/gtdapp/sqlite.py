import sqlite3
from os.path import exists



conn = sqlite3.connect("todo.db")
print("Opened database successfully")

if not exists("todo.db"):
    conn.execute("""CREATE TABLE ITEM
    (ID INT PRIMARY KEY     NOT NULL,
    TITLE           TEXT    NOT NULL,
    DESCRIPTION     CHAR(200));""")
    print("Table created successfully")

    conn.execute("INSERT INTO ITEM (ID,TITLE,DESCRIPTION) \
      VALUES (1, 'hoge', 'hogehoge')");
    conn.execute("INSERT INTO ITEM (ID,TITLE,DESCRIPTION) \
      VALUES (2, 'fuga', 'fugafuga')");
    conn.execute("INSERT INTO ITEM (ID,TITLE,DESCRIPTION) \
      VALUES (3, 'piyo', 'piyopiyo')");
    conn.commit()
    print("Records created successfully")

conn.close()

def get_entries(conn):
    cursor = conn.execute("SELECT id, title, description from ITEM")
    desc = cursor.description
    column_names = [col[0] for col in desc]
    data = [dict(zip(column_names, row)) for row in cursor.fetchall()]
    return data
