from telegram.ext import Updater, CommandHandler

from subscription import subscription
from config import TOKEN


def hello(update, context):
    update.message.reply_text(
        'Hello {}'.format(update.message.from_user.first_name))


updater = Updater(TOKEN, use_context=True)

updater.start_polling()
updater.idle()
