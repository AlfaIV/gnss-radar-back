from dependency_injector import containers, providers

from app.services.sattelites_positions import SatellitesPositions
from app.services.signal_processing import SignalProcessing


class Container(containers.DeclarativeContainer):
    wiring_config = containers.WiringConfiguration(
        modules=[
            "app.api.v1.endpoints.satellities",
        ]
    )

    satellite_services = providers.Singleton(SatellitesPositions)
    signal_services = providers.Singleton(SignalProcessing)
