from datetime import date
from ariadne import convert_kwargs_to_snake_case

from api import db
from api.models import Post


@convert_kwargs_to_snake_case
def create_post_resolver(obj, info, title, description):
    try:
        today = date.today()
        post = Post(
            title=title,
            description=description,
            created_at=today.strftime("%b-%d-%Y")
        )
        db.session.add(post)
        db.session.commit()
        payload = {
            "success": True,
            "post": post.to_dict()
        }
    except ValueError:
        payload = {
            "success": False,
            "errors": ["Incorrect information about post"]
        }
    return payload


@convert_kwargs_to_snake_case
def update_post_reslover(obj, info, id, title=None, description=None):
    try:
        post = Post.query.get(id)
        if post:
            post.title = title if title is not None else post.title
            post.description = description if description is not None else post.description
        db.session.add(post)
        db.session.commit()
        payload = {
            "success": True,
            "post": post.to_dict()
        }
    except AttributeError:
        payload = {
            "success": False,
            "errors": [f"item matching id {id} not found"]
        }

    return payload