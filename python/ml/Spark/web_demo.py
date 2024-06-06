import os

from sparkai.llm.llm import ChatSparkLLM, ChunkPrintHandler
from sparkai.core.messages import ChatMessage


# url
SPARKAI_URL = 'wss://spark-api.xf-yun.com/v1.1/chat'

SPARKAI_APP_ID = os.getenv("SPARK_APPID")
SPARKAI_API_SECRET = os.getenv("SPARK_SEC")
SPARKAI_API_KEY = os.getenv("SPARK_APIKEY")

SPARKAI_DOMAIN = 'general'


if __name__ == "__main__":
    spark = ChatSparkLLM(
            spark_api_url=SPARKAI_URL,
            spark_app_id=SPARKAI_APP_ID,
            spark_api_key=SPARKAI_API_KEY,
            spark_api_secret=SPARKAI_API_SECRET,
            spark_llm_domain=SPARKAI_DOMAIN,
            streaming=False,
            )

    messages = [ChatMessage(role='user', content='你是谁？')]

    handler = ChunkPrintHandler()
    a = spark.generate([messages], callback=[handler])
    print(a)
