from django.shortcuts import render

def post_index(request):
    ctx = {}
    return render(request, 'index.html', ctx)

def post_detail(request):
    ctx = {}
    return render(request, 'detail.html', ctx)

def create_post(request):
    ctx = {}
    return render(request, 'form.html', ctx)

def update_post(request):
    ctx = {}
    return render(request, 'form.html', ctx)

def signup(request):
    ctx = {}
    return render(request, 'signup.html', ctx)

def login(request):
    ctx = {}
    return render(request, 'login.html', ctx)