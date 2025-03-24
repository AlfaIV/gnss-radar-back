import os
from typing import List
import numpy as np
import scipy as sp
from scipy.signal import get_window, find_peaks
from scipy.ndimage import gaussian_filter
from app.schemas import SignalDataRequest, SignalProcessingSpectrumResponce, SpectrumPeaksParameters

class SignalProcessing:
    def __init__(self):
        pass
        # TODO: нужно будет сделать, если будем считывать загруенные с телескопа данные

    def custom_peak_widths(signal: List[int], peaks: List[int]):
        widths = []
        for peak in peaks:
            if peak != 0:
                left_index_array = np.argwhere(signal[:peak] < (signal[peak] - 3))
                left_index = None if left_index_array.shape == (0, 1) else left_index_array[-1][0]
            else:
                left_index = None

            if peak != len(signal)-1:
                right_index_array = np.argwhere(signal[peak:] < (signal[peak] - 3))
                right_index = None if right_index_array.shape == (0, 1) else right_index_array[0][0] + peak
            else:
                right_index = None
            widths.append([peak, left_index, right_index])
        return np.array(widths)
    
    def get_spectrum_params(self, data: SignalDataRequest) -> None:
        
        I_signal = np.array([1, 2, 3, 4, 5])
        Q_signal = np.array([6, 7, 8, 9, 10])
        # I_signal = data.data[0::2]
        # Q_signal = data.data[1::2]
        signal = np.empty(len(I_signal), np.complex64)
        for i in range(len(signal)):
            signal[i] = complex(I_signal[i], Q_signal[i])
        
        window = get_window('hamming', len(signal))
        signal = signal * window
        workers = os.cpu_count()
        fft_result = np.fft.fftshift(sp.fft.fft(signal, workers=workers//2))
        frequencies = np.fft.fftfreq(len(signal))
        magnitude = np.abs(fft_result)
        magnitude_normalized = magnitude / np.max(magnitude)
        magnitude_db = 20*np.log10(magnitude_normalized)

        gauss_filt = gaussian_filter(magnitude_db, sigma=150)
        frequencies = frequencies[:len(gauss_filt)]

        peaks, _ = find_peaks(gauss_filt, prominence=3)
        custom_peak_widths = self.custom_peak_widths(gauss_filt, peaks)

        # return SignalProcessingSpectrumResponce(
        #     peaks = [SpectrumPeaksParameters(
        #         peak_number=int(peak[0]),
        #         # peak_height=int(signal[peak[0]]),
        #         # peak_width=int(frequencies(peak[2]) - frequencies(peak[1])) if peak[2] is not None and peak[1] is not None else None
        #     ) for peak in custom_peak_widths],
        # )

        return None