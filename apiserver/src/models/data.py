from dataclasses import dataclass

@dataclass
class Ticket:
    prefix: str = ""
    ticket_id: int = 0
    first_name: str = ""
    last_name: str = ""
    phone_number: str = ""
    preference: str = ""

@dataclass
class Basket:
    prefix: str = ""
    basket_id: int = 0
    description: str = ""
    donors: str = ""
    winning_ticket: int = 0
