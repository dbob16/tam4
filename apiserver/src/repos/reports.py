from .template import RepoTemplate
from ..models import Winner, WinnerByName

class ReportsRepo(RepoTemplate):
    def winners_by_basket(self, prefix: str = ""):
        self.cur.execute("SELECT * FROM winners WHERE prefix = %s", (prefix,))
        results = self.cur.fetchall()
        return [Winner(*r) for r in results]
    def winners_by_name(self, prefix: str = ""):
        self.cur.execute("SELECT * FROM winners_by_name WHERE prefix = %s", (prefix,))
        results = self.cur.fetchall()
        return [WinnerByName(*r) for r in results]
