from typing import List, Optional
from pydantic import BaseModel, Field

class SpectrumPeaksParameters(BaseModel):
    peak_number: int = Field(ge=0)
    # peak_height: int = Field(le=0)
    # peak_width: Optional[int] = Field(ge=0)

class SignalProcessingSpectrumResponce(BaseModel):
    peaks: List[SpectrumPeaksParameters]

class SignalDataRequest():
    data: List[float]