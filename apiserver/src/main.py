from fastapi import FastAPI
from fastapi.responses import RedirectResponse
import sys
from .database import init_db
from .repos import ApiKeyRepo
from .router_appender import import_routers

app_title = "TAM4 API Server"
app_desc = "The API Server for Ticket Auction Manager 4."

init_db()

args = sys.argv

if "dev" in args:
    app = FastAPI(title=app_title, description=app_desc)
else:
    app = FastAPI(title=app_title, description=app_desc, docs_url=None, redoc_url=None)

@app.get("/")
def red_to_api():
    return RedirectResponse(url="/api")

@app.get("/api")
def main_route(api_key: str = ""):
    auth_status = ApiKeyRepo().check_api_key(api_key)
    return {"whoami": "TAM4 Server", "status": "healthy", "authenticated": auth_status}

import_routers(app)
