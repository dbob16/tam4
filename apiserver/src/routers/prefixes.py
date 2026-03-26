from fastapi import APIRouter
from ..repos import PrefixRepo, ApiKeyRepo
from ..models import Prefix

prefix_router = APIRouter(prefix="/api/prefixes")

@prefix_router.get("")
def get_all_prefixes(api_key: str = ""):
    """API endpoint that gets all prefixes."""
    ApiKeyRepo().verify_api_key(api_key)
    return PrefixRepo().get_all_prefixes()

@prefix_router.get("/{prefix}")
def get_one_prefix(api_key: str = "", prefix: str = ""):
    """API endpoint that gets one prefix."""
    ApiKeyRepo().verify_api_key(api_key)
    return PrefixRepo().get_one_prefix(prefix)

@prefix_router.post("")
def post_one_prefix(p: Prefix, api_key: str = ""):
    """API endpoint that posts one prefix."""
    ApiKeyRepo().verify_api_key(api_key)
    return PrefixRepo().post_one_prefix(p)

@prefix_router.delete("")
def delete_one_prefix(prefix: str, api_key: str = ""):
    """API endpoint that deletes one prefix."""
    ApiKeyRepo().verify_api_key(api_key)
    return PrefixRepo().delete_one_prefix(prefix)
