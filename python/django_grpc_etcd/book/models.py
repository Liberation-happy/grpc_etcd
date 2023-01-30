from django.db import models


# Create your models here.

class Book(models.Model):
    title = models.CharField(max_length=128, verbose_name="名称")
    comment = models.CharField(max_length=256, verbose_name="描述")

    class Meta:
        verbose_name = "书籍"
        verbose_name_plural = "书籍"
