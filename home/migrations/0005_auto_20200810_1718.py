# Generated by Django 3.1 on 2020-08-10 08:18

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('home', '0004_auto_20200810_1653'),
    ]

    operations = [
        migrations.AlterField(
            model_name='album',
            name='cover_image',
            field=models.ImageField(default='LiS3_1.jpg', upload_to='img/upload/'),
        ),
    ]
