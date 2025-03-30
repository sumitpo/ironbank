from sqlalchemy import create_engine
from sqlalchemy.orm import sessionmaker

dsn = "mysql+mysqlconnector://cfo:cfopass@localhost:3306/bank"

def init():
    engine = create_engine(dsn)
    dbSession = sessionmaker(bind=engine)
    
    pass