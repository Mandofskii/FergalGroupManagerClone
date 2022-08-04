from pyrogram import Client
from redis import StrictRedis

app = Client("fergalUserBot", 000, "000", phone_number="+000")
redisdb = StrictRedis("127.0.0.1", 6379, 9)
@app.on_message()
async def main(cli: Client, msg):
    if msg.text.startswith("/getid"):
        base = msg.text.split('\n')
        user = (await cli.get_users(base[0].split(' ')[1]))
        redisdb.set('group:' + base[1] + ':hash:' + base[2], user.first_name + '|' + str(user.id))
app.run()
