package model

const (
	ImportsTemplate = `
import numpy as np
import datetime
import asyncio
import aiohttp
import time
from datetime import datetime, timezone, timedelta
import logging
from rtlsdr import RtlSdr
`
	ConstantsTemplate = `
# Константы
BASE_URL = '{{.BaseURL}}'
TOKEN = '{{.Token}}'
SAMPLE_NUM = {{.SampleNum}}

# Настройка логирования
logging.basicConfig(level=logging.DEBUG)

# Адреса обработчиков
HANDLERS = {
    'power': f"{BASE_URL}/power",
    'spectrum': f"{BASE_URL}/spectrum",
    'upload': f"{BASE_URL}/upload",
    'pair_measurement': f"{BASE_URL}/pair_measurement"
}
`
	RequestHandlerTemplate = `
# Класс для обработки запросов
class RequestHandler:
    def __init__(self, token):
        self.token = token

    async def send_request(self, url, data):
        async with aiohttp.ClientSession() as session:
            async with session.post(url, json=data) as response:
                if response.status == 200:
                    logging.info(f'Успешно отправлен запрос на {url}')
                    logging.info('Ответ сервера:', await response.text())
                else:
                    logging.error(f'Ошибка при отправке запроса на {url}: {response.status}')
                    logging.error('Ответ сервера:', await response.text())

    def create_base_data(self, start_time, end_time, group, target, signal):
        return {
            'token': self.token,
            'description': {
                'startTime': start_time,
                'endTime': end_time,
                'group': group,
                'target': target,
                'signal': signal
            }
        }

    async def send_power(self, start_time, time_step, power):
        end_time = datetime.now(timezone.utc).isoformat()
        data = self.create_base_data(start_time, end_time, "GPS", "G12", "L1")
        data['data'] = {
            'power': power.tolist(),
            'startTime': start_time,
            'timeStep': time_step,
        }
        await self.send_request(HANDLERS['power'], data)

    async def send_spectrum(self, start_time, freqs, spectrum):
        end_time = datetime.now(timezone.utc).isoformat()
        data = self.create_base_data(start_time, end_time, "GPS", "G12", "L1")
        data['data'] = {
            'spectrum': np.log10(np.abs(spectrum)).tolist(),
            'startFreq': freqs.tolist()[0],
            'freqStep': freqs.tolist()[1] - freqs.tolist()[0],
            'startTime': start_time,
        }
        await self.send_request(HANDLERS['spectrum'], data)
`
	SignalProcessorTemplate = `
# Класс для обработки сигналов
class SignalProcessor:
    def __init__(self, sample_rate):
        self.sample_rate = sample_rate

    def calculate_spectrum(self, signal_samples):
        spectrum = np.fft.fftshift(np.fft.fft(signal_samples))
        freqs = np.fft.fftshift(np.fft.fftfreq(len(signal_samples), 1/self.sample_rate))
        return freqs, spectrum

    def calculate_mean_power(self, signal_samples):
        instantaneous_power = np.abs(signal_samples) ** 2
        return instantaneous_power
`
	SdrHandlerTemplate = `
# Класс для работы с SDR
class SDRHandler:
    def __init__(self, sample_rate=2.048e6, center_freq=1227.60e6, gain='auto'):
        self.sample_rate = sample_rate
        self.center_freq = center_freq
        self.gain = gain
        self.sdr = None

    def connect(self):
        try:
            self.sdr = RtlSdr()
            self.sdr.sample_rate = self.sample_rate
            self.sdr.center_freq = self.center_freq
            self.sdr.gain = self.gain
            logging.info("Успешное подключение к RTL-SDR.")
            logging.info("Успешная настройка параметров RTL-SDR.")
        except Exception as e:
            logging.error("Ошибка подключения к RTL-SDR:", e)
            raise e

    def disconnect(self):
        if self.sdr:
            self.sdr.close()
            logging.info("Успешное отключение от SDR")

    def read_samples(self, count):
        if self.sdr:
            samples = self.sdr.read_samples(count)
            logging.info(f'Получено {len(samples)} выборок.')
            return samples
        else:
            raise Exception("SDR не подключен")
`
	DataProcessorTemplate = `
# Класс для обработки данных
class DataProcessor:
    def __init__(self, sdr_handler, request_handler, signal_processor):
        self.sdr_handler = sdr_handler
        self.request_handler = request_handler
        self.signal_processor = signal_processor

    async def process_data(self):
        try:
            self.sdr_handler.connect()
            while True:
                start_time = datetime.now(timezone.utc).isoformat()
                time_step = (datetime.now(timezone.utc) + timedelta(seconds=1/self.sdr_handler.sample_rate)).isoformat()
                samples = self.sdr_handler.read_samples(SAMPLE_NUM)

                power = self.signal_processor.calculate_mean_power(samples)
                logging.info(f'Рассчитанная мощность: {power}')
                await self.request_handler.send_power(start_time, time_step, power)

                freqs, spectrum = self.signal_processor.calculate_spectrum(samples)
                await self.request_handler.send_spectrum(start_time, freqs, spectrum)

                await asyncio.sleep(5)
        finally:
            self.sdr_handler.disconnect()
`
	MainFunctionTemplate = `
# Основная функция
async def main():
    sdr_handler = SDRHandler()
    request_handler = RequestHandler(TOKEN)
    signal_processor = SignalProcessor(sdr_handler.sample_rate)
    data_processor = DataProcessor(sdr_handler, request_handler, signal_processor)

    await data_processor.process_data()

if __name__ == "__main__":
    asyncio.run(main())
`
)

type PythonGenConfig struct {
	BaseURL   string
	Token     string
	SampleNum int
}
