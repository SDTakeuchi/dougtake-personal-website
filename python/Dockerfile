FROM python:3.11.1-slim-buster

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apt-get update \
&&  pip install --upgrade pip \
&&  pip install --upgrade setuptools

WORKDIR /app/

ADD requirements.txt .

RUN pip install -r requirements.txt

COPY ./config ./config

EXPOSE 8000

# CMD ["python", "./config/manage.py", "runserver", "0.0.0.0:8000"]