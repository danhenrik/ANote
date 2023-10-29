import requests
import psycopg2
import json
import time

class pg_conn:
  def __init__(self):
    try:
      conn = psycopg2.connect("host=database user=anote password=anote dbname=anote port=5432")
      self.conn = conn
      self.c = conn.cursor()
    except Exception as e: 
      print(e.__str__())

  def SQLCmd(self, cmd):
    try:
      print("Executing : \"" + cmd + '"')
      self.c.execute(cmd)
      self.conn.commit()
    except Exception as e:
      print("Execution failed!\nError: "+ e.__str__())
      self.conn.commit()

print("Waiting 10 seconds for services to start")
time.sleep(10)

# Create ES Index
try:
  res = requests.put('http://elasticsearch:9200/notes')
  print(res)
  if str(res.status_code).startswith('2'):
    print("ES Index created")
  else:
    print("ES Index creation failed")
    obj = json.loads(res.content.decode('utf-8'),)
    print(obj.get('error').get('root_cause')[0].get('type'))	
except Exception as e: 
  print(e.__str__())

# Configure ES Index
# configure date format
try:
  res = requests.put('http://elasticsearch:9200/notes/_mapping', 
    headers={
      'Content-Type': 'application/json'
    },
    json={
        "properties": {
          "published_date": {
            "type": "date",
            "format": "yyyy-MM-dd"
          },
          "updated_date": {
            "type": "date",
            "format": "yyyy-MM-dd"
          }
        }
      })
  print(res)
  if str(res.status_code).startswith('2'):
    print("ES Index configured")
  else:
    print("ES Index configuration failed")
    obj = json.loads(res.content.decode('utf-8'),)
    print(obj.get('error').get('root_cause')[0].get('type'))
except Exception as e:
  print(e.__str__())

# Create Postgres Tables
pg_conn = pg_conn()

pg_conn.SQLCmd("""CREATE TABLE users (
                id varchar(30) PRIMARY KEY, 
                email varchar(255) NOT NULL UNIQUE, 
                password varchar(72),
                google_id varchar,
                created_at date NOT NULL DEFAULT NOW(),
                avatar varchar(255)
               );""")

pg_conn.SQLCmd("""CREATE TABLE notes (
                id uuid PRIMARY KEY, 
                title varchar NOT NULL, 
                author_id varchar(30) NOT NULL,
                content varchar NOT NULL,
                created_at date NOT NULL DEFAULT NOW(),
                updated_at date NOT NULL DEFAULT NOW(),
                CONSTRAINT user_FK
                 FOREIGN KEY(author_id)
                  REFERENCES users(id) ON DELETE SET NULL
               );""")

pg_conn.SQLCmd("""CREATE TABLE likes (
                id uuid PRIMARY KEY, 
                user_id varchar(30), 
                note_id uuid, 
                created_at timestamp NOT NULL DEFAULT NOW(),
                CONSTRAINT user_FK 
                 FOREIGN KEY(user_id) 
                  REFERENCES users(id) ON DELETE CASCADE, 
                CONSTRAINT notes_FK 
                 FOREIGN KEY(note_id) 
                  REFERENCES notes(id) ON DELETE CASCADE
              );""")

pg_conn.SQLCmd("""ALTER TABLE likes
                  DROP COLUMN id;
              """)

pg_conn.SQLCmd("""ALTER TABLE likes
                  ADD PRIMARY KEY (user_id, note_id);
              """)

pg_conn.SQLCmd("""CREATE TABLE comments (
                id uuid PRIMARY KEY,
                user_id varchar(30),
                note_id uuid,
                content varchar(500) NOT NULL, 
                created_at timestamp NOT NULL DEFAULT NOW(),
                CONSTRAINT user_FK
                 FOREIGN KEY(user_id)
                  REFERENCES users(id) ON DELETE CASCADE,
                CONSTRAINT notes_FK
                 FOREIGN KEY(note_id)
                  REFERENCES notes(id) ON DELETE SET NULL
               );""")

pg_conn.SQLCmd("""CREATE TABLE tags (
                id varchar PRIMARY KEY,
                name varchar(50) NOT NULL
               );""")

pg_conn.SQLCmd("""CREATE TABLE note_tags (
                id uuid PRIMARY KEY,
                note_id uuid,
                tag_id varchar,
                CONSTRAINT notes_FK
                 FOREIGN KEY(note_id)
                  REFERENCES notes(id) ON DELETE CASCADE,
                CONSTRAINT tags_FK
                 FOREIGN KEY(tag_id)
                  REFERENCES tags(id) ON DELETE CASCADE
                );""")

pg_conn.SQLCmd("""CREATE TABLE communities (
                id uuid PRIMARY KEY,
                name varchar(50) NOT NULL,
                background varchar(255)
                );""")

pg_conn.SQLCmd("""CREATE TABLE community_members (
                id uuid PRIMARY KEY,
                user_id varchar(30),
                community_id uuid,
                CONSTRAINT user_FK
                 FOREIGN KEY(user_id)
                  REFERENCES users(id) ON DELETE CASCADE,
                CONSTRAINT community_FK
                 FOREIGN KEY(community_id)
                  REFERENCES communities(id) ON DELETE CASCADE
                );""")

pg_conn.SQLCmd("""CREATE TABLE community_notes (
                id uuid PRIMARY KEY,
                note_id uuid,
                community_id uuid,
                CONSTRAINT notes_FK
                 FOREIGN KEY(note_id)
                  REFERENCES notes(id) ON DELETE CASCADE,
                CONSTRAINT community_FK
                 FOREIGN KEY(community_id)
                  REFERENCES communities(id) ON DELETE CASCADE
                );""")

pg_conn.SQLCmd("""CREATE TABLE tokens (
                id SERIAL PRIMARY KEY,
                token varchar(255) NOT NULL,
                user_id varchar(30) NOT NULL,
                CONSTRAINT user_FK
                 FOREIGN KEY(user_id)
                  REFERENCES users(id) ON DELETE CASCADE
                );""")

# Create Postgres replication slot
pg_conn.SQLCmd("""SELECT * 
                FROM pg_create_logical_replication_slot(
                 'es_replication_slot', 
                 'wal2json',
                 false,
                 true
               );""")
    
# Create Postgres triggers
pg_conn.SQLCmd("""CREATE FUNCTION notify() 
               RETURNS TRIGGER LANGUAGE PLPGSQL AS 
                $$ 
                BEGIN 
                 NOTIFY es_replicate; 
                 RETURN NEW; 
                END; 
                $$""")

pg_conn.SQLCmd("""CREATE TRIGGER notify_notes 
                AFTER INSERT OR UPDATE OR DELETE 
                ON notes FOR EACH ROW 
                EXECUTE PROCEDURE notify();""")

pg_conn.SQLCmd("""CREATE TRIGGER notify_likes 
                AFTER INSERT OR DELETE 
                ON likes FOR EACH ROW 
                EXECUTE PROCEDURE notify();""")
               
pg_conn.SQLCmd("""CREATE TRIGGER notify_comments 
                AFTER INSERT OR DELETE 
                ON comments FOR EACH ROW 
                EXECUTE PROCEDURE notify();""")
               
pg_conn.SQLCmd("""CREATE TRIGGER notify_note_tags 
                AFTER INSERT OR DELETE 
                ON note_tags FOR EACH ROW 
                EXECUTE PROCEDURE notify();""")

pg_conn.SQLCmd("""CREATE TRIGGER notify_community_notes
                AFTER INSERT OR DELETE 
                ON community_notes FOR EACH ROW 
                EXECUTE PROCEDURE notify();""")
