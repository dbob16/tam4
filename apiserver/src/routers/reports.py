from fastapi import APIRouter

from ..models import Winner, WinnerByName
from ..repos import ReportsRepo, ApiKeyRepo

reports_router = APIRouter(prefix="/api/reports")

@reports_router.get("/bybasket/{prefix}")
def get_basket_report(api_key: str = "", prefix: str = "") -> list[Winner]:
    ApiKeyRepo().verify_api_key(api_key)
    return ReportsRepo().winners_by_basket(prefix)

@reports_router.get("/byname/{prefix}")
def get_name_report(api_key: str = "", prefix: str = "") -> list[WinnerByName]:
    ApiKeyRepo().verify_api_key(api_key)
    return ReportsRepo().winners_by_name(prefix)
