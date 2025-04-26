from fastapi import FastAPI

from app.api.v1.routes import routers as v1_routers
from app.core.config import configs
from app.core.container import Container
# from app.core.logger import setup_logging

class AppCreator:
    def __init__(self):
        self.app = FastAPI(
            title=configs.PROJECT_NAME,
            openapi_url=f"{configs.API}/openapi.json",
            version="0.0.1",
        )

        self.container = Container()

        # setup_logging()

        @self.app.get("/")
        async def root():
            return f"{configs.PROJECT_NAME} service is working"

        self.app.include_router(v1_routers, prefix=configs.API_V1_STR)


app_creator = AppCreator()
app = app_creator.app
container = app_creator.container
