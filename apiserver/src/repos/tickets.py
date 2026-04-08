from .template import RepoTemplate
from ..models import Ticket

class TicketRepo(RepoTemplate):
    def get_single_ticket(self, prefix: str = "", ticket_id: int = 0):
        self.cur.execute("SELECT * FROM tickets WHERE prefix = %s AND ticket_id = %s", (prefix, ticket_id))
        result = self.cur.fetchone()
        if result:
            return Ticket(*result)
        else:
            return Ticket(prefix, ticket_id)
    def get_ticket_range(self, prefix: str = "", start_id: int = 0, end_id: int = 0):
        self.cur.execute("SELECT * FROM tickets WHERE prefix = %s AND ticket_id BETWEEN %s and %s", (prefix, start_id, end_id))
        results = self.cur.fetchall()
        return [Ticket(*r) for r in results]
    def post_ticket_range(self, ts: list[Ticket]):
        stmt = "INSERT INTO tickets VALUES (%s, %s, %s, %s, %s, %s) ON CONFLICT (prefix, ticket_id) DO UPDATE SET first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, phone_number = EXCLUDED.phone_number, preference = EXCLUDED.preference"
        for t in ts:
            self.cur.execute(stmt, (t.prefix, t.ticket_id, t.first_name, t.last_name, t.phone_number, t.preference))
        self.conn.commit()
        return ts
