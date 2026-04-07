from .template import RepoTemplate
from ..models import Basket

class BasketRepo(RepoTemplate):
    def get_basket_range(self, prefix: str = "", start_id: int = 0, end_id: int = 0):
        self.cur.execute("SELECT * FROM baskets WHERE prefix = %s AND basket_id BETWEEN %s AND %s", (prefix, start_id, end_id))
        results = self.cur.fetchall()
        return [Basket(*r) for r in results]
    def post_basket_range(self, bs: list[Basket]):
        stmt = "INSERT INTO baskets VALUES (%s, %s, %s, %s, %s) ON CONFLICT (prefix, basket_id) DO UPDATE SET description = EXCLUDED.description, donors = EXCLUDED.donors"
        for b in bs:
            self.cur.execute(stmt, (b.prefix, b.basket_id, b.description, b.donors, b.winning_ticket))
        self.conn.commit()
        return bs
