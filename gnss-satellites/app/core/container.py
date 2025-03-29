from dependency_injector import containers, providers

from app.core.s3_database import S3Database

from app.repositories import S3Repository

from app.services import SatellitesPositions


class Container(containers.DeclarativeContainer):
    wiring_config = containers.WiringConfiguration(
        modules=[
            "app.api.v1.endpoints.satellities",
        ]
    )

    # Инициализация S3
    s3_db = providers.Singleton(S3Database)
    
    # Репозиторий получает метод session как фабрику
    s3_repository = providers.Factory(
        S3Repository,
        get_session=s3_db.provided.session  # Используем provided
    )
    satellite_services = providers.Singleton(SatellitesPositions, s3_repository=s3_repository)
