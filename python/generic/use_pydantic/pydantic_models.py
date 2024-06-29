import random
from typing import Self
from datetime import date, timedelta
from uuid import UUID, uuid4
from enum import Enum
from pydantic import BaseModel, EmailStr, Field, field_validator, model_validator


def rand_factory():
    return str(random.randint(1000, 9999))


class Department(Enum):
    HR = "HR"
    SALES = "SALES"
    IT = "IT"
    ENGINEERING = "ENGINEERING"


class Employee(BaseModel):
    employee_id: UUID = Field(default_factory=uuid4, frozen=True)
    name: str = Field(min_length=1, frozen=True)
    email: EmailStr = Field(pattern=r".+@example\.com$")
    date_of_birth: date = Field(alias="birth_date", repr=False, frozen=True)
    salary: float = Field(alias="compensation", gt=0, repr=False)
    department: Department
    elected_benefits: bool

    @field_validator("date_of_birth")
    @classmethod
    def check_valid_age(cls, date_of_birth: date) -> date:
        today = date.today()
        eighteen_years_ago = date(today.year - 18, today.month, today.day)

        if date_of_birth > eighteen_years_ago:
            raise ValueError("Employee must be at least 18 years old.")

        return date_of_birth

    @model_validator(mode="after")
    def check_it_benefits(self) -> Self:
        department = self.department
        elected_benefits = self.elected_benefits

        if department == Department.IT and elected_benefits:
            raise ValueError("IT employees are contractors and don't qualify for benefits")
        return self

class Student(BaseModel):
    student_id: str = Field(default_factory=rand_factory, frozen=True)
    name: str
    age: int = Field(gt=0)
    score: float = Field(repr=False)



def main() -> None:
    young_employee_date = {
            "name": "Jake Bar",
            "email": "jabr@example.com",
            "birth_date": date.today() - timedelta(days=365*17),
            "compensation": 90_000,
            "department": "SALES",
            "elected_benefits": True,
    }

    e = Employee.model_validate(young_employee_date)
    print(e)



if __name__ == "__main__":
    main()


