from typing import Optional

import uvicorn
from fastapi import Depends, FastAPI, HTTPException
from pydantic import BaseModel
from sqlmodel import SQLModel, create_engine, Session, Field, select


class Hero(SQLModel, table=True):
    id: Optional[int] = Field(primary_key=True, index=True)
    name: str
    secret_name: str
    age: Optional[int] = None


app = FastAPI()

# Database Setup
DATABASE_URL = "sqlite:///./database.db"
engine = create_engine(DATABASE_URL, connect_args={"check_same_thread": False})
SQLModel.metadata.create_all(engine)


def get_session():
    with Session(engine) as session:
        yield session

# Create a hero
@app.post("/heroes/", response_model=Hero)
def create_hero(hero: Hero, session: Session = Depends(get_session)):
    session.add(hero)
    session.commit()
    session.refresh(hero)
    return hero

# Read all hero
@app.get("/heroes/", response_model=list[Hero])
def read_heroes(skip: int = 0, limit: int = 10, session: Session = Depends(get_session)):
    heroes = session.exec(select(Hero).offset(skip).limit(limit)).all()
    return heroes


# Read a hero by id
@app.get("/heroes/{hero_id}", response_model=Hero)
def read_hero(hero_id: int, session: Session = Depends(get_session)):
    hero = session.get(Hero, hero_id)
    if not hero:
        raise HTTPException(status_code=404, detail="Hero not found")
    return hero


@app.put("/heroes/{hero_id}", response_model=Hero)
def update_hero(hero_id: int, hero_data: Hero, session: Session = Depends(get_session)):
    hero = session.get(Hero, hero_id)
    if not hero:
        raise HTTPException(status_code=404, detail="Hero not found")
    # Update the hero's attributes
    for field, value in hero_data.model_dump().items():
        setattr(hero, field, value)

    session.commit()
    session.refresh(hero)
    return hero


@app.delete("/heroes/{hero_id}", response_model=Hero)
def delete_hero(hero_id: int, session: Session = Depends(get_session)):
    hero = session.get(Hero, hero_id)
    if not hero:
        raise HTTPException(status_code=404, detail="Hero not found")

    session.delete(hero)
    session.commit()
    return hero

@app.get("/test")
def test():
    return {"msg": "test!"}
