from django.db import models

# Create your models here.
class Album(models.Model):
    title = models.CharField(max_length=80)
    artist = models.CharField(max_length=80)
    image = models.ImageField()
    release_date = models.DateField()
    landr_url = models.URLField()
    SINGLE = 'Single'
    ALBUM = 'Album'
    EP = 'EP'
    ALBUM_TYPE_CHOICES = (
        (SINGLE, 'Single'),
        (ALBUM, 'Album'),
        (EP, 'EP')
    )
    album_type = models.CharField(max_length=10, choices=ALBUM_TYPE_CHOICES)

    def __str__(self):
        return self.title