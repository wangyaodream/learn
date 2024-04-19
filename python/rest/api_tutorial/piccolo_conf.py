from piccolo.conf.apps import AppRegistry
from piccolo.engine.postgres import PostgresEngine


DB = PostgresEngine(config={
    "database": "api_project",
    "user": "dbuser",
    "password": "Dream462213925",
    "host": "localhost",
    "port": 5432,
})


# A list of paths to piccolo apps
# e.g. ['blog.piccolo_app']
APP_REGISTRY = AppRegistry(apps=['sql_app.piccolo_app'])
