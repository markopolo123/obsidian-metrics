global:
  telegram_api_url: https://api.telegram.org
route:
  group_by: [alertname]
  receiver: "telegram" # default receiver
  repeat_interval: 24h
  routes:
    - receiver: "telegram"
      repeat_interval: 12h
      matchers:
        - severity="insane"

receivers:
  - name: "null"
  - name: "telegram"
    telegram_configs:
      - send_resolved: true
        api_url: https://api.telegram.org #global.telegram_api_url
        bot_token: op://homelab/7wgsgg3dhn3wudfbzo6wffkibe/credential
        chat_id: op://homelab/7wgsgg3dhn3wudfbzo6wffkibe/chat_id
        message: {{ "{{ template \"telegram.default.message\" .}}" }}
        parse_mode: ""
