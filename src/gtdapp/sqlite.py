import sqlite3

conn = sqlite3.connect("todo.db")
print("Opened database successfully")

conn.execute("""CREATE TABLE IF NOT EXISTS ITEM
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
