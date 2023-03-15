from django.urls import path
from . import views

urlpatterns = [
    path('posts', views.post_index, name='posts_index'),
    path('posts/<str:pk>', views.post_detail, name='post_detail'),

    path('comment', views.create_comment, name='create_comment')
]
