
from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE
# from . import setup
import sys
sys.path.append("/root/ABE-master/test")
import setup
import pika

# 创建和队列的连接
connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
channel = connection.channel()
policy_str = ''
# 定义callback处理函数
def callback(ch, method, properties, body):
    global policy_str
    # policy_str = str(body)
    policy_str = body.decode("utf-8")
    channel.stop_consuming()

# choose a random message
sym_key = setup.pairing_group.random(GT)
print("产生了随机群元素作为key",sys.getsizeof(sym_key))

# 发送 sym_key 
# channel.basic_publish(exchange='python.go.trans',routing_key='pytogo',body=sym_key.encode("utf-8"))
channel.basic_publish(exchange='python.go.trans',routing_key='pytogo',body=str(sym_key).encode("utf-8"))

# 获取访问结构
channel.basic_consume(queue='python.go.trans.gotopy', on_message_callback=callback, auto_ack=True)
channel.start_consuming()

# 处理收到的字符串，其实时访问结构+#+目标目录
res = policy_str.split("#")
policy_str = res[0]
dest_dir = res[1]
print("get policy_str:",policy_str)
print("get policy_str:",dest_dir)


# generate ABE(key)
abe_key = setup.cpabe.encrypt(setup.pk, sym_key, policy_str)
print("产生了加密对称密钥ABE_key",type(str(abe_key)))

#abe_key写入文件
# print(abe_key)
with open(dest_dir+'/abe_key', 'wb') as f:
    f.write(str(abe_key).encode("utf-8"))
