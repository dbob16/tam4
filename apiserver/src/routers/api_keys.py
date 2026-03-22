from fastapi import APIRouter, HTTPException, status
from ..config import Config
from ..models import ApiReq
from ..repos import ApiKeyRepo

api_key_router = APIRouter(prefix="/api/api_keys")

invalid_pw_exception = HTTPException(status_code=status.HTTP_401_UNAUTHORIZED, detail="Invalid API PW")

@api_key_router.get("")
def get_api_keys(api_pw: str = ""):
    sys_api_pw = Config().read_config()["api_pw"]
    if api_pw == sys_api_pw:
        return ApiKeyRepo().get_api_keys()
    else:
        raise invalid_pw_exception

@api_key_router.post("")
def post_api_key(api_req: ApiReq):
    sys_api_pw = Config().read_config()["api_pw"]
    if api_req.api_pw == sys_api_pw:
        return ApiKeyRepo().create_api_key(api_req.computer_name)
    else:
        raise invalid_pw_exception

@api_key_router.delete("")
def delete_api_key(api_pw: str, api_key: str):
    sys_api_pw = Config().read_config()["api_pw"]
    if api_pw == sys_api_pw:
        return {"detail": ApiKeyRepo().delete_api_key(api_key)}
    else:
        raise invalid_pw_exception
