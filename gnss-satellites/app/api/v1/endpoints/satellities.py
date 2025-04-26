from dependency_injector.wiring import Provide, inject
from fastapi import APIRouter, Depends

from app.core.container import Container
from app.schemas.sattelites_position import (
    SatellitesTimeRequest,
    RadarPositionGeograthRequest,
    SatellitesPositionResponce,
    SatellitesTimeResponce,
    
)
from app.services import SatellitesPositions
from app.core.logger import logger

router = APIRouter(
    prefix="/satellites",
    tags=["satellites"],
)


@router.post("/now", response_model=SatellitesPositionResponce)
@inject
def post_sattelites(
    radar: RadarPositionGeograthRequest,
    service: SatellitesPositions = Depends(Provide[Container.satellite_services]),
):
    logger.info("Запрос к ручке /now")
    return service.get_sattelites_positions(radar)

@router.post("/time", response_model=SatellitesTimeResponce)
@inject
def post_time_sattelites(
    radar: SatellitesTimeRequest,
    service: SatellitesPositions = Depends(Provide[Container.satellite_services]),
):
    logger.info("Запрос к ручке /time")
    return service.get_sattelites_times_vison(radar)