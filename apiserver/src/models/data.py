from dataclasses import dataclass

@dataclass
class Ticket:
    prefix: str = ""
    ticket_id: int = 0
    first_name: str = ""
    last_name: str = ""
    phone_number: str = ""
    preference: str = ""
