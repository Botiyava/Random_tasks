import os
import pandas as pd

def free_sms(sms_number):
    if sms_number - 10 <= 0:
        return 0
    else:
        return sms_number - 10

def rate(chosen_user_data):
    if len(chosen_user_data) == 0:
        print("We don't have this number in our list")
        return
    price = [2,0,1] #1 variant
    result_price = 0
    for i in range(len(dfdf)):
        result_price +=  chosen_user_data['call_duration'][i] * price[0] + (price[2] * free_sms(chosen_user_data['sms_number'][i]) )
    return result_price

number = 915783624 #our number or any other if you want
#number = int(input("Phone number: "))

#working directory with CDR
os.chdir("/home/botiyava/materials")
df = pd.read_csv('data.csv')
dfdf = df[df["msisdn_origin"] == number].reset_index()
chosen_user_data = dfdf.drop(columns=["timestamp", "msisdn_origin", "msisdn_dest"])


#billing for chosen user
print(rate(chosen_user_data), "₽ to pay")


