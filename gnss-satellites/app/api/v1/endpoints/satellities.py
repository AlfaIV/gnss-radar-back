from dependency_injector.wiring import Provide, inject
from fastapi import APIRouter, Depends

from app.core.container import Container
from app.schemas.sattelites_position import (
    RadarPositionRequest,
    SatellitesPositionResponce,
)
from app.services import SatellitesPositions

router = APIRouter(
    prefix="/satellites",
    tags=["satellites"],
)


@router.post("/now", response_model=SatellitesPositionResponce)
@inject
def post_sattelites(
    radar: RadarPositionRequest,
    service: SatellitesPositions = Depends(Provide[Container.satellite_services]),
):
    return service.get_sattelites_positions(radar)
