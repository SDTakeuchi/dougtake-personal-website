from django.shortcuts import render
from helpers import post

def post_index(request):
    arg = {
        'tag_id': request.GET.get('tag_id'),
        'search_char': request.GET.get('search_char'),
        'offset': request.GET.get('offset'),
        'limit': request.GET.get('limit'),
    }

    posts = post.find_posts(**arg)
    ctx = {
        'posts': posts,
    }
    return render(request, 'index.html', ctx)

def post_detail(request):
    if post_id := request.GET.get('post_id') is None:
        arg = {'post_id': post_id}
    found_post = post.find_posts(request)
    ctx = {
        'post': found_post,
    }
    return render(request, 'detail.html', ctx)

# def create_post(request):
#     ctx = {}
#     return render(request, 'post_form.html', ctx)

# def update_post(request):
#     ctx = {}
#     return render(request, 'post_form.html', ctx)

# user

# def signup(request):
#     ctx = {}
#     return render(request, 'signup.html', ctx)

# def login(request):
#     ctx = {}
#     return render(request, 'login.html', ctx)