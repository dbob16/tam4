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
    cur.execute("""CREATE TABLE IF NOT EXISTS tickets (
        prefix VARCHAR(255),
        ticket_id INT,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        phone_number VARCHAR(255),
        preference VARCHAR(20),
        PRIMARY KEY (prefix, ticket_id))""")
    cur.execute("""CREATE TABLE IF NOT EXISTS baskets (
        prefix VARCHAR(255),
        basket_id INT,
        description VARCHAR(255),
        donors VARCHAR(255),
        winning_ticket INT,
        PRIMARY KEY (prefix, basket_id))""")
    cur.execute("""CREATE OR REPLACE VIEW winners AS
        SELECT b.prefix, b.basket_id, b.description, b.winning_ticket, t.first_name, t.last_name, t.phone_number, t.preference
        FROM baskets b LEFT JOIN tickets t ON b.prefix = t.prefix AND b.winning_ticket = t.ticket_id
        ORDER BY b.prefix, b.basket_id""")
    conn.commit()
    conn.close()
    print("Database initiated successfully.")
