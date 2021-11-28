import requests, random, json, colorama

def login(Username,Password):
    letters = 'qwertyuiopasdfghjklzxcvbnm123456790qwertyiobuzxcvbasdfr142'
    token = ''.join(random.choice(letters) for x in range(23)) 
    print(token)
    session = requests.Session()
    url = "https://twitter.com/sessions"
    session.headers = {
    "Host": "twitter.com",
    "Content-Type": "application/x-www-form-urlencoded",
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"
    }
    cookies = {
    "_mb_tk":token
    }
    data = {
    "authenticity_token":token,
    "session[username_or_email]":Username,
    "session[password]":Password
    }
    response = session.post(url, data=data , cookies=cookies)
    y = session.cookies.get_dict()
    try:
        c = str(y["ct0"])
        v = str(y["auth_token"])	
        print(c)
        print(v)
        return y
    except:
        return False

login("exgoju", "Hokejka1")