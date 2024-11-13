FROM python:3.12-slim

WORKDIR /app

COPY . /app

RUN apt-get update && apt-get install libgl1 ffmpeg libsm6 libxext6  -y

RUN pip install --no-cache-dir -r requirements.txt

# Run app.py when the container launches
CMD ["python", "main.py"]