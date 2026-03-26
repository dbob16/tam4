from fastapi import FastAPI
from . import routers

def import_routers(app: FastAPI):
    app.include_router(routers.api_key_router)
    app.include_router(routers.prefix_router)
