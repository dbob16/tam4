from fastapi import HTTPException, status
from .template import RepoTemplate
from ..models import ApiKey
import random as rand
import string

random_pick_from = string.ascii_uppercase + string.digits

class ApiKeyRepo(RepoTemplate):
    """The primary repo to handle API Keys."""
    def get_api_keys(self) -> list[ApiKey]:
        """Returns all API keys from server."""
        self.cur.execute("SELECT * FROM api_keys")
        results = self.cur.fetchall()
        return [ApiKey(*r) for r in results]
    def check_api_key(self, api_key: str) -> bool:
        """Checks API Key and returns true or false indicating status."""
        self.cur.execute("SELECT * FROM api_keys WHERE api_key = %s", (api_key,))
        results = self.cur.fetchall()
        if len(results) > 0:
            return True
        else:
            return False
    def verify_api_key(self, api_key: str):
        """Checks API Key and errors out if it doesn't exist in the DB."""
        if not self.check_api_key(api_key):
            raise HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Invalid API Key.")
        else:
            return True
    def create_api_key(self, computer_name: str):
        """Creates an api key on request."""
        while True:
            new_key = "".join(rand.choice(random_pick_from) for r in range(16))
            self.cur.execute("SELECT * FROM api_keys WHERE api_key = %s", (new_key,))
            results = self.cur.fetchall()
            if len(results) == 0:
                break
        self.cur.execute("INSERT INTO api_keys VALUES (%s, %s)", (new_key, computer_name))
        self.conn.commit()
        return {"api_key": new_key}
    def delete_api_key(self, api_key: str):
        """Deletes an API Key."""
        self.cur.execute("DELETE FROM api_keys WHERE api_key = %s", (api_key,))
        self.conn.commit()
        return "API Key deleted successfully."
