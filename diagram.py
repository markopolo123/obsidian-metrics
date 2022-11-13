from diagrams import Diagram

from diagrams.onprem.monitoring import Prometheus, Grafana
from diagrams.saas.chat import Telegram
from diagrams.custom import Custom
from diagrams.programming.language import Go

with Diagram("Obsidian Note Metrics", show=False):
    obsidian = Custom("Obsidian Vault", "./resources/obsidian.png")
    prometheus = Prometheus("Time Metrics")
    alertmanager = Prometheus("Alert Manager")
    grafana = Grafana("Visualisation")
    telegram = Telegram("Notification")
    exporter = Go("Exporter")

    obsidian >> exporter
    exporter >> prometheus
    prometheus >> grafana
    prometheus >> alertmanager
    alertmanager >> telegram
