import os


from openai import OpenAI, completions
from dotenv import load_dotenv


def main():
    load_dotenv()
    client = OpenAI(
            api_key=os.getenv('ZHIPU_API_KEY'),
            base_url="https://open.bigmodel.cn/api/paas/v4/"
            )

    completion = client.chat.completions.create(
            model="glm-4-flash",
            messages=[
                {"role": "system", "content": "你是一个经验丰富的python工程师"},
                {"role": "user", "content": "我想知道python中异步编程的相关知识"}
            ],
            top_p=0.7,
            temperature=0.9
    )

    print(completion.choices[0].message.content)


if __name__ == "__main__":
    main()

