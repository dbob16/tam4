from fastapi import APIRouter

from ..models import Basket
from ..repos import BasketRepo, ApiKeyRepo

basket_router = APIRouter(prefix="/api/baskets")

@basket_router.get("/{prefix}/{start_id}/{end_id}")
def get_basket_range(api_key: str = "", prefix: str = "", start_id: int = 0, end_id: int = 0):
    ApiKeyRepo().verify_api_key(api_key)
    return BasketRepo().get_basket_range(prefix, start_id, end_id)

@basket_router.post("")
def post_basket_range(api_key: str = "", bs: list[Basket] = []):
    ApiKeyRepo().verify_api_key(api_key)
    return BasketRepo().post_basket_range(bs)
