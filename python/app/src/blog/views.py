from django.shortcuts import render
from helpers import post

def post_index(request):
    posts = post.find_posts(request)
    ctx = {}
    return render(request, 'index.html', ctx)

def post_detail(request):
    posts = post.find_posts(request)
    ctx = {}
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