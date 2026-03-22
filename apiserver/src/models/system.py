from dataclasses import dataclass

@dataclass
class ApiKey:
    api_key: str = ""
    computer_name: str = ""

@dataclass
class ApiReq:
    api_pw: str = ""
    computer_name: str = ""

@dataclass
class Prefix:
    prefix: str = ""
    color: str = ""
    weight: int = 0
