from django.shortcuts import render, redirect
from django.core.mail import EmailMessage
from django.contrib import messages
from django.template.loader import render_to_string
from dougtake.settings import EMAIL_HOST_USER, MAIN_EMAIL
from .models import *
from .forms import *
from . import forms

# Create your views here.
def businessTopView (request):
    contact_form = forms.ContactForm()

    if request.method == 'POST':
        form = forms.ContactForm(request.POST)
        context={'form':form}
        subject = render_to_string('mail_templates/send_message/subject.txt')
        message = render_to_string('mail_templates/send_message/message.txt', context)
        msg = EmailMessage(subject, message, EMAIL_HOST_USER, [MAIN_EMAIL, EMAIL_HOST_USER])
        msg.send()

        messages.success(request, "Message has been sent to Douglas.")
        return redirect('businessTop')

    return render(request, 'business/business.html')

def musicTopView (request):
    return render(request, 'music/index.html')