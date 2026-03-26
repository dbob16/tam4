from .template import RepoTemplate
from ..models import Prefix

class PrefixRepo(RepoTemplate):
    def get_all_prefixes(self) -> list[Prefix]:
        """Returns all prefixes in a list of dataclasses."""
        self.cur.execute("SELECT * FROM prefixes")
        results = self.cur.fetchall()
        return [Prefix(*r) for r in results]
    def get_one_prefix(self, prefix: str) -> Prefix:
        """Returns one prefix as a dataclass."""
        self.cur.execute("SELECT * FROM prefixes WHERE prefix = %s ORDER BY weight, prefix", (prefix,))
        result = self.cur.fetchone()
        if result:
            return Prefix(*result)
        else:
            return Prefix(prefix, "gray", 0)
    def post_one_prefix(self, p: Prefix):
        """Posts one prefix."""
        self.cur.execute("INSERT INTO prefixes VALUES (%s, %s, %s) ON CONFLICT (prefix) DO UPDATE SET color = EXCLUDED.color, weight = EXCLUDED.weight", (p.prefix, p.color, p.weight))
        self.conn.commit()
        return {"detail": "Prefix posted successfully."}
    def delete_one_prefix(self, prefix: str):
        """Deletes one prefix."""
        self.cur.execute("DELETE FROM prefixes WHERE prefix = %s", (prefix,))
        self.conn.commit()
        return {"detail": "Prefix deleted successfully."}
