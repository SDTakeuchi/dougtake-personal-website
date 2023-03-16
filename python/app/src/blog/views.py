from django.shortcuts import render
from .helpers import post, comment

def post_index(request):
    arg = {
        'post_id': request.GET.get('post_id'),
        'tag_id': request.GET.get('tag_id'),
        'search_char': request.GET.get('search_char'),
        'offset': request.GET.get('offset'),
        'limit': request.GET.get('limit'),
    }

    res = post.find_posts(**arg)
    if res.get('data') is None:
        ctx = {
            'message': res.get('message'),
        }
        return render(request, 'error.html', ctx)

    ctx = {
        'res': res,
    }
    return render(request, 'react/main/index.html', ctx)

def post_detail(request):
    if post_id := request.GET.get('post_id') is None:
        arg = {'post_id': post_id}

    res = post.find_posts(**arg)
    if res.get('data') is None:
        ctx = {
            'message': res.get('message'),
        }
        return render(request, 'error.html', ctx)

    ctx = {
        'res': res,
    }
    return render(request, 'react/main/index.html', ctx)

# def create_post(request):
#     ctx = {}
#     return render(request, 'post_form.html', ctx)

# def update_post(request):
#     ctx = {}
#     return render(request, 'post_form.html', ctx)

# comment

def create_comment(request):
    args = {
        'post_id': request.POST.get('post_id'),
        'body': request.POST.get('body'),
    }

    res = comment.create_comment(**args)
    if res.get('data') is None:
        ctx = {
            'message': res.get('message'),
        }
        return render(request, 'error.html', ctx)

    ctx = {
        'res': res,
    }
    # if post creation is done asyncronesly, html does not have to be returned.
    return render(request, 'react/main/index.html', ctx)

# user

# def signup(request):
#     ctx = {}
#     return render(request, 'signup.html', ctx)

# def login(request):
#     ctx = {}
#     return render(request, 'login.html', ctx)