from typing import Optional

from sqlmodel import Field, SQLModel, create_engine, Session


class Hero(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    name: str
    secret_name: str
    age: Optional[int] = None



# Create rows
h1 = Hero(name="WangYao", secret_name="ZhangMan")
h2 = Hero(name="Alice", secret_name="Bob")
h3 = Hero(name="Carter", secret_name="Liv", age=99)


engine = create_engine("sqlite:///database.db")

SQLModel.metadata.create_all(engine)

with Session(engine) as session:
    session.add(h1)
    session.add(h2)
    session.add(h3)
    session.commit()

