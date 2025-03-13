from fastapi import APIRouter

from app.api.v1.endpoints.satellities import router as satellities_router

routers = APIRouter()
router_list = [satellities_router]

for router in router_list:
    router.tags = routers.tags.append("v1")
    routers.include_router(router)
