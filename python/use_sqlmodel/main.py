from typing import Optional

import uvicorn
from fastapi import Depends, FastAPI, HTTPException
from pydantic import BaseModel
from sqlmodel import Relationship, SQLModel, create_engine, Session, Field, select


class Team(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    name: str

    heroes: list["Hero"] = Relationship(back_populates="team")


class Hero(SQLModel, table=True):
    id: Optional[int] = Field(primary_key=True, index=True)
    name: str
    secret_name: str
    age: Optional[int] = None
    team_id: Optional[int] = Field(default=None, foreign_key="team.id")

    team: Optional[Team] = Relationship(back_populates="heroes")
    missions: list["Mission"] = Relationship(
        back_populates="heroes",
        link_model="HeroMissionLink",
    )

class Mission(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True)
    description: str

    heroes: list[Hero] = Relationship(
            back_populates="missions", link_model="HeroMissionLink"
    )


class HeroMissionLink(SQLModel, table=True):
    hero_id: Optional[int] = Field(
            default=None, foreign_key="hero.id", primary_key=True
            )
    mission_id: Optional[int] = Field(
            default=None, foreign_key="mission.id", primary_key=True
            )


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

# Create Mission
@app.post("/missions/", response_model=Mission)
def create_mission(mission: Mission, session: Session = Depends(get_session)):
    session.add(mission)
    session.commit()
    session.refresh()

    return mission

@app.put("/missions/{mission_id}/heroes/{hero_id}", response_model=Mission)
def assign_hero_to_mission(mission_id: int, hero_id: int, session: Session = Depends(get_session)):
    hero = session.get(Hero, hero_id)
    mission = session.get(Mission, mission_id)
    if not hero or not mission:
        raise HTTPException(status_code=404, detail="Hero or Mission not found")
    hero_mission_link = HeroMissionLink(hero_id=hero_id, mission_id=mission_id)
    session.add(hero_mission_link)
    session.commit()
    return mission

@app.get("teams/{team_id}", response_model=Team)
def read_team(team_id: int, session: Session = Depends(get_session)):
    team = session.get(Team, team_id)
    if not team:
        raise HTTPException(status_code=404, detail="Team not found")
    return team

@app.get("/test")
def test():
    return {"msg": "test!"}

