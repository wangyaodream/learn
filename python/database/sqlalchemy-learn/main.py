from sqlalchemy import (
        create_engine,
        ForeignKey,
        Column,
        String,
        Integer,
        CHAR
)
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.exc import IntegrityError


Base = declarative_base()


class Person(Base):
    __tablename__ = "people"

    ssn = Column("ssn", Integer, primary_key=True)
    name = Column("name", String)
    gender = Column("gender", CHAR)
    age = Column("age", Integer)


    def __init__(self, ssn, name, gender, age) -> None:
        self.ssn = ssn
        self.name = name
        self.gender = gender
        self.age = age

    def __repr__(self):
        return f"({self.ssn}) {self.name} ({self.gender}, {self.age})"



def insert_ignore(session, entry):
    try:
        session.add(entry)
        session.commit()
    except IntegrityError:
        session.rollback()


engine = create_engine("sqlite:///mydb.db", echo=True)
Base.metadata.create_all(bind=engine)

Session = sessionmaker(bind=engine)
session = Session()


def main():
    p1 = Person(12312, "wangyao", "m", 30)
    p2 = Person(12313, "Anna", "f", 30)
    p3 = Person(12314, "Bob", "m", 30)
    p4 = Person(12317, "Cherry", "f", 18)

    p_data = [p1, p2, p3, p4]
    for i in p_data:
        insert_ignore(session, i)

    result = session.query(Person).all()

    print(result)


if __name__ == "__main__":
    main()






