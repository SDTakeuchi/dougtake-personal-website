from django.db import models

# Create your models here.
class Album(models.Model):
    title = models.TextField()
    cover_image = models.ImageField(upload_to='img/upload/', default='LiS3_1.jpg')
    release_date = models.DateField()
    landr_url = models.URLField()

    def __str__(self):
        return self.title

# def snippet(self):
#     return self.body[:50] + "..."
#     then it be like,
#     => I think that... 