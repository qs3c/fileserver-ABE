from charm.toolbox.pairinggroup import PairingGroup, GT
from ABE.ac17 import AC17CPABE
from setup import cpabe,pk


# 随机选择对称加密密钥key
key = pairing_group.random(GT)


# 获取访问结构设定
policy_str = '((ONE and THREE) and (TWO OR FOUR))'

# 产生ABE加密密钥
ctxt = cpabe.encrypt(pk, key, policy_str)

# 发送密钥key给Go做对称加密


# 输出加密密钥为一个加密密钥文件