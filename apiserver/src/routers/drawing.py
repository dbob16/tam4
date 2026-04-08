from fastapi import APIRouter

from ..models import Basket
from ..repos import DrawingRepo, ApiKeyRepo

drawing_router = APIRouter(prefix="/api/drawing")

@drawing_router.get("/{prefix}/{start_id}/{end_id}")
def get_drawing_range(api_key: str = "", prefix: str = "", start_id: int = 0, end_id: int = 0):
    ApiKeyRepo().verify_api_key(api_key)
    return DrawingRepo().get_range_winners(prefix, start_id, end_id)

@drawing_router.post("")
def post_drawing_range(api_key: str = "", bs: list[Basket] = []):
    ApiKeyRepo().verify_api_key(api_key)
    return DrawingRepo().post_range_winners(bs)
