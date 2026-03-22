from ..database import new_session

class RepoTemplate:
    def __init__(self):
        self.conn, self.cur = new_session()
