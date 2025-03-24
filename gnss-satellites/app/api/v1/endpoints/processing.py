from typing import Optional
from dependency_injector.wiring import Provide, inject
from fastapi import APIRouter, Depends

from app.core.container import Container
from app.schemas import (
    SignalDataRequest,
    SignalProcessingSpectrumResponce,
)
from app.services import SignalProcessing

router = APIRouter(
    prefix="/processsing",
    tags=["processing"],
)

# @router.post("/spectrum", response_model=SignalProcessingSpectrumResponce)
# @inject
# def post_spectrum(
#     signal_data: SignalDataRequest,
#     service: SignalProcessing = Depends(Provide[Container.signal_services]),
# ):
#     return service.get_spectrum_params(signal_data)

@router.post("/spectrum", response_model=SignalProcessingSpectrumResponce)
@inject
def post_spectrum(
    signal_data: SignalDataRequest,
    service: SignalProcessing = Depends(Provide[Container.signal_services]),
):
    return service.get_spectrum_params(signal_data)