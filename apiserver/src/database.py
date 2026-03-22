from .config import Config
import psycopg2

def new_session():
    """Creates and returns a new conn and cursor object for the currently configured db connection."""
    dbc = Config().read_config()["db"]
    conn = psycopg2.connect(
        f"host={dbc["host"]} port={dbc["port"]} user={dbc["user"]} password={dbc["password"]} dbname={dbc["database"]} sslmode=allow"
    )
    cur = conn.cursor()
    return conn, cur

def init_db():
    conn, cur = new_session()
    cur.execute("CREATE TABLE IF NOT EXISTS api_keys (api_key VARCHAR(255) PRIMARY KEY, computer_name VARCHAR(255))")
    cur.execute("CREATE TABLE IF NOT EXISTS prefixes (prefix VARCHAR(255) PRIMARY KEY, color VARCHAR(255), weight INT)")
    conn.commit()
    conn.close()
    print("Database initiated successfully.")
