from django.forms import ModelForm
from django import  forms

class ContactForm(forms.Form):
    toEmail = forms.EmailField(required=True)
    senderName = forms.CharField(required=True)
    phoneNum = forms.IntegerField()
    body = forms.CharField(required=True)

    def __str__(self):
        return self.toEmail