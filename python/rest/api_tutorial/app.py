import typing

from piccolo.utils.pydantic import create_pydantic_model
from piccolo.engine import engine_finder

from blacksheep import Application, FromJSON, json, Response, Content

from sql_app.tables import Expense


ExpenseModelIn: typing.Any = create_pydantic_model(
        table=Expense, 
        model_name=" ExpenseModelIn")

ExpenseModelOut: typing.Any = create_pydantic_model(
        table=Expense,
        include_default_columns=True,
        model_name=" ExpenseModelOut")

ExpenseModelPartial: typing.Any = create_pydantic_model(
        table=Expense,
        model_name="ExpenseModelPartial",
        all_optional=True)

app = Application()


@app.router.get("/expenses")
async def expenses():
    try:
        expense = await Expense.select()
        return expense
    except Exception:
        return Response(404, content=Content(b"text/plain", b"Not Found"))


