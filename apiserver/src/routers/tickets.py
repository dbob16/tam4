from fastapi import APIRouter

from ..models import Ticket
from ..repos import TicketRepo, ApiKeyRepo

ticket_router = APIRouter(prefix="/api/tickets")

@ticket_router.get("/{prefix}/{start_id}/{end_id}")
def get_ticket_range(api_key: str = "", prefix: str = "", start_id: int = 0, end_id: int = 0):
    ApiKeyRepo().verify_api_key(api_key)
    return TicketRepo().get_ticket_range(prefix, start_id, end_id)

@ticket_router.post("")
def post_ticket_range(api_key: str = "", ts: list[Ticket] = []):
    ApiKeyRepo().verify_api_key(api_key)
    return TicketRepo().post_ticket_range(ts)
