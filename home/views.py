from django.shortcuts import render
from .models import Album

# Create your views here.
def homeView(request):
    albums = Album.objects.all()
    return render(request, 'home.html')

def musicView(request):
    albums = Album.objects.all()
    return render(request, 'music.html', {'albums':albums})
    
def tourView(request):
    return render(request, 'tour.html')
    
def contactView(request):
    return render(request, 'contact.html')

    