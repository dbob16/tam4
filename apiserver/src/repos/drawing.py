from .template import RepoTemplate
from ..models import Basket, Winner

class DrawingRepo(RepoTemplate):
    def get_range_winners(self, prefix: str, start_id: int, end_id: int):
        self.cur.execute("SELECT * FROM winners WHERE prefix = %s AND basket_id BETWEEN %s AND %s", (prefix, start_id, end_id))
        results = self.cur.fetchall()
        return [Winner(*r) for r in results]

    def post_range_winners(self, bs: list[Basket]):
        stmt = "INSERT INTO baskets (prefix, basket_id, winning_ticket) VALUES (%s, %s, %s) ON CONFLICT (prefix, basket_id) DO UPDATE SET winning_ticket = EXCLUDED.winning_ticket"
        for b in bs:
            self.cur.execute(stmt, (b.prefix, b.basket_id, b.winning_ticket))
        self.conn.commit()
        return bs
