# python
FROM python:3.11.1-slim-buster

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN apt-get update \
&&  pip install --upgrade pip \
&&  pip install --upgrade setuptools

WORKDIR /app/

COPY ./app .

RUN pip install -r ./requirements.txt

# node
ENV NODE_VERSION=18.15.0
ENV NVM_DIR=/root/.nvm
ENV PATH="/root/.nvm/versions/node/v${NODE_VERSION}/bin/:${PATH}"

RUN apt install -y curl \
&&  curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash \
&&  . "$NVM_DIR/nvm.sh" && nvm install ${NODE_VERSION} \
&&  . "$NVM_DIR/nvm.sh" && nvm use v${NODE_VERSION} \
&&  . "$NVM_DIR/nvm.sh" && nvm alias default v${NODE_VERSION} \
&&  node --version \
&&  npm --version \
&&  npm i @mui/material @mui/icons-material @emotion/styled @emotion/react markdown-to-jsx

EXPOSE 8000

# CMD ["python", "./config/manage.py", "runserver", "0.0.0.0:8000"]