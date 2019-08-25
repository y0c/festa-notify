# Festa Notify

[![Go Report Card](https://goreportcard.com/badge/github.com/y0c/festa-notify)](https://goreportcard.com/report/github.com/y0c/festa-notify)
[![CircleCI](https://circleci.com/gh/y0c/festa-notify.svg?style=svg)](https://circleci.com/gh/y0c/festa-notify)

## 👀 Overview 

[Festa](https://festa.io/) 플랫폼에 특정 이벤트에 대해서 알림을 받기 위한 서비스입니다. <br/>
키워드를 등록해놓고 매칭되는 이벤트 제목이나, 주최자 에 해당하는 이벤트가 올라오면 메일로 알려줍니다. 

![image](https://user-images.githubusercontent.com/2585676/63648974-9f0ae000-c772-11e9-8785-672b5efa224b.png)



## 🛠 Built with

- Go 1.12
- Serverless (for lambda)
- CloudWatch Event
- AWS SES
- Google Cloud Firestore
- CircleCI

## 📦 Command

- `make build` - build go application
- `make deploy` - deploy app to lambda
